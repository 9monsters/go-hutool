package http

import (
	"net"
	"net/http"

	hunet "github.com/nine-monsters/go-hutool/hutool-net/net"
)

const (
	XForwardedFor = "X-Forwarded-For"
	XRealIP       = "X-Real-IP"
	XClientIP     = "x-client-ip"
)

// GetLocalIP returns the non-loop back local IP of the host.
func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return hunet.LOCAL_IPV4_IP
	}
	for _, address := range addrs {
		// check the address type and if it is not a loop back, the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return hunet.LOCAL_IPV4_IP
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

	if remoteAddr == hunet.LOCAL_IPV6_IP {
		remoteAddr = hunet.LOCAL_IPV4_IP
	}

	return remoteAddr
}
