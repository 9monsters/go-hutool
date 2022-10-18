package net

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"net/http"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/nine-monsters/go-hutool/hutool-core/text"
)

const (
	LOCAL_IP           string = "127.0.0.1"
	IP_SPLIT_MARK      string = text.DASHED
	IP_MASK_SPLIT_MARK string = text.SLASH
	IP_MASK_MAX        int    = 32
)

// IsIpV4 return a string if in IPv4 format.
func IsIpV4(ip string) bool {
	return strings.Count(ip, ":") < 2
}

// IsIpV6 return a string if in IPv6 format.
func IsIpV6(ip string) bool {
	return strings.Count(ip, ":") >= 2
}

// ListAllIp list all ip
func ListAllIp(predicate func(ip net.IP) bool, ifaceNames ...string) ([]net.IP, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, fmt.Errorf("fail to get interfaces, err: %w", err)
	}

	ips := make([]net.IP, 0)
	matcher := NewIfaceNameMatcher(ifaceNames)
	for _, i := range interfaces {
		f := i.Flags
		// skip interface not up, loopback and other mater name
		if i.HardwareAddr == nil ||
			f&net.FlagUp != net.FlagUp ||
			f&net.FlagLoopback == net.FlagLoopback ||
			!matcher.Matches(i.Name) {
			continue
		}
		addrs, err := i.Addrs()

		// skip error
		if err != nil {
			continue
		}

		ips = collectAddresses(predicate, addrs, ips)
	}

	return ips, nil
}

// ListAllIpV4 list all IPv4 addresses.
// ifaceNames are used to specified interface names (filename wild match pattern supported also, like eth*).
func ListAllIpV4(ifaceNames ...string) ([]string, error) {
	ips := make([]string, 0)
	_, err := ListAllIp(func(ip net.IP) (yes bool) {
		s := ip.String()
		if yes = IsIpV4(s); yes {
			ips = append(ips, s)
		}

		return yes
	}, ifaceNames...)

	return ips, err
}

// ListAllIpV6 list all IPv6 addresses.
// ifaceNames are used to specified interface names (filename wild match pattern supported also, like eth*).
func ListAllIpV6(ifaceNames ...string) ([]string, error) {
	ips := make([]string, 0)

	_, err := ListAllIp(func(ip net.IP) (yes bool) {
		s := ip.String()
		if yes = IsIpV6(s); yes {
			ips = append(ips, s)
		}

		return yes
	}, ifaceNames...)

	return ips, err
}

// ListIfaceNames list all net interface names.
func ListIfaceNames() (names []string) {
	list, err := net.Interfaces()
	if err != nil {
		return nil
	}

	for _, i := range list {
		f := i.Flags
		// skip not up nor loop back
		if i.HardwareAddr == nil || f&net.FlagUp == 0 || f&net.FlagLoopback == 1 {
			continue
		}

		names = append(names, i.Name)
	}

	return names
}

func MakeSliceMap(ss []string) map[string]bool {
	m := make(map[string]bool)

	for _, s := range ss {
		if s != "" {
			m[s] = true
		}
	}

	return m
}

type IfaceNameMatcher struct {
	ifacePatterns map[string]bool
}

func NewIfaceNameMatcher(names []string) IfaceNameMatcher {
	return IfaceNameMatcher{ifacePatterns: MakeSliceMap(names)}
}

func (i IfaceNameMatcher) Matches(name string) bool {
	if len(i.ifacePatterns) == 0 {
		return true
	}

	if _, ok := i.ifacePatterns[name]; ok {
		return true
	}

	for k := range i.ifacePatterns {
		if ok, _ := filepath.Match(k, name); ok {
			return true
		}
	}

	for k := range i.ifacePatterns {
		if strings.Contains(k, name) {
			return true
		}
	}

	return false
}

// collectAddresses collect distinct ip list
func collectAddresses(predicate func(net.IP) bool, addrs []net.Addr, ips []net.IP) []net.IP {
	for _, addr := range addrs {
		var ip net.IP
		switch v := addr.(type) {
		case *net.IPAddr:
			ip = v.IP
		case *net.IPNet:
			ip = v.IP
		default:
			continue
		}
		pFlag := predicate != nil && predicate(ip)
		if !ContainsIs(ips, ip) && pFlag {
			ips = append(ips, ip)
		}
	}
	return ips
}

// ContainsIs judge ip is or not in list
func ContainsIs(ips []net.IP, ip net.IP) bool {
	for _, n := range ips {
		if n.Equal(ip) {
			return true
		}
	}
	return false
}

// Outbound  gets preferred outbound ip of this machine.
func Outbound() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return ""
	}

	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("close connection error, %s", err)
		}
	}(conn)

	s := conn.LocalAddr().String()
	return s[:strings.LastIndex(s, ":")]
}

// MainIP tries to get the main IP address and the IP addresses.
func MainIP(ifaceName ...string) (string, []string) {
	return MainIPVerbose(false, ifaceName...)
}

// MainIPVerbose tries to get the main IP address and the IP addresses.
func MainIPVerbose(verbose bool, ifaceName ...string) (string, []string) {
	ips, _ := ListAllIpV4(ifaceName...)
	if len(ips) == 1 {
		return ips[0], ips
	}

	if s := findMainIPByIfconfig(verbose, ifaceName); s != "" {
		return s, ips
	}

	if out := Outbound(); out != "" && contains(ips, out) {
		return out, ips
	}

	if len(ips) > 0 {
		return ips[0], ips
	}

	return "", nil
}

func findMainIPByIfconfig(verbose bool, ifaceName []string) string {
	names := ListIfaceNames()
	if verbose {
		log.Printf("iface names: %s", names)
	}

	var matchedNames []string
	matcher := NewIfaceNameMatcher(ifaceName)
	for _, n := range names {
		if matcher.Matches(n) {
			matchedNames = append(matchedNames, n)
		}
	}

	if verbose && len(matchedNames) < len(names) {
		log.Printf("matchedNames: %s", matchedNames)
	}

	if len(matchedNames) == 0 {
		return ""
	}

	name := matchedNames[0]
	for _, n := range matchedNames {
		// for en0 on mac or eth0 on linux
		if strings.HasPrefix(n, "e") && strings.HasSuffix(n, "0") {
			name = n
			break
		}
	}

	/*
		# ifconfig en0
		en0: flags=8863<UP,BROADCAST,SMART,RUNNING,SIMPLEX,MULTICAST> mtu 1500
		options=50b<RXCSUM,TXCSUM,VLAN_HWTAGGING,AV,CHANNEL_IO>
		ether 14:98:77:60:f7:b6
		inet6 fe80::c3e:cd25:e594:8269%en0 prefixlen 64 secured scopeid 0x6
		inet 192.168.16.153 netmask 0xfffffe00 broadcast 192.168.17.255
		nd6 options=201<PERFORMNUD,DAD>
		media: autoselect (1000baseT <full-duplex>)
		status: active
	*/
	re := regexp.MustCompile(`inet\s+([\w.]+?)\s+`)
	if verbose {
		log.Printf("exec comd: ifconfig %s", name)
	}
	c := exec.Command("ifconfig", name)
	if co, err := c.Output(); err == nil {
		if verbose {
			log.Printf("output: %s", co)
		}
		sub := re.FindStringSubmatch(string(co))
		if len(sub) > 1 {
			if verbose {
				log.Printf("found: %s", sub[1])
			}
			return sub[1]
		}
	} else if verbose {
		log.Printf("error: %v", err)
	}

	return ""
}

func contains(strArr []string, target string) bool {
	sort.Strings(strArr)
	index := sort.SearchStrings(strArr, target)

	if index < len(strArr) && strArr[index] == target {
		return true
	}
	return false
}

func LongToIPv4(longIP int64) string {
	var buffer bytes.Buffer
	buffer.WriteString(strconv.FormatInt(longIP>>24&0xFF, 10))
	buffer.WriteString(text.DOT)
	buffer.WriteString(strconv.FormatInt(longIP>>16&0xFF, 10))
	buffer.WriteString(text.DOT)
	buffer.WriteString(strconv.FormatInt(longIP>>8&0xFF, 10))
	buffer.WriteString(text.DOT)
	buffer.WriteString(strconv.FormatInt(longIP&0xFF, 10))
	return buffer.String()
}

// Define http headers.
const (
	XForwardedFor = "X-Forwarded-For"
	XRealIP       = "X-Real-IP"
	XClientIP     = "x-client-ip"
)

// GetLocalIP returns the non loopback local IP of the host.
func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "127.0.0.1"
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return "127.0.0.1"
}

// RemoteIP returns the remote ip of the request.
func RemoteIP(req *http.Request) string {
	remoteAddr := req.RemoteAddr
	if ip := req.Header.Get(XClientIP); ip != "" {
		remoteAddr = ip
	} else if ip := req.Header.Get(XRealIP); ip != "" {
		remoteAddr = ip
	} else if ip = req.Header.Get(XForwardedFor); ip != "" {
		remoteAddr = ip
	} else {
		remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
	}

	if remoteAddr == "::1" {
		remoteAddr = "127.0.0.1"
	}

	return remoteAddr
}
