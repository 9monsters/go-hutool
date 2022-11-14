package validator

import "regexp"

// Used by IsFilePath func
const (
	// Unknown is unresolved OS type
	Unknown = iota
	// Win is Windows type
	Win
	// Unix is *nix OS types
	Unix
)

var (
	RX_GENERAL           = regexp.MustCompile(General)
	RX_NUMBERS           = regexp.MustCompile(Numbers)
	RX_HEX               = regexp.MustCompile(Hex)
	RX_Email             = regexp.MustCompile(Email)
	RX_CreditCard        = regexp.MustCompile(CreditCard)
	RX_ISBN10            = regexp.MustCompile(ISBN10)
	RX_ISBN13            = regexp.MustCompile(ISBN13)
	RX_UUID3             = regexp.MustCompile(UUID3)
	RX_UUID4             = regexp.MustCompile(UUID4)
	RX_UUID5             = regexp.MustCompile(UUID5)
	RX_UUID              = regexp.MustCompile(UUID)
	RX_Alpha             = regexp.MustCompile(Alpha)
	RX_Alphanumeric      = regexp.MustCompile(Alphanumeric)
	RX_Numeric           = regexp.MustCompile(Numeric)
	RX_Int               = regexp.MustCompile(Int)
	RX_Float             = regexp.MustCompile(Float)
	RX_Hexadecimal       = regexp.MustCompile(Hexadecimal)
	RX_Hexcolor          = regexp.MustCompile(Hexcolor)
	RX_RGBcolor          = regexp.MustCompile(RGBcolor)
	RX_ASCII             = regexp.MustCompile(ASCII)
	RX_PrintableASCII    = regexp.MustCompile(PrintableASCII)
	RX_Multibyte         = regexp.MustCompile(Multibyte)
	RX_FullWidth         = regexp.MustCompile(FullWidth)
	RX_HalfWidth         = regexp.MustCompile(HalfWidth)
	RX_Base64            = regexp.MustCompile(Base64)
	RX_DataURI           = regexp.MustCompile(DataURI)
	RX_MagnetURI         = regexp.MustCompile(MagnetURI)
	RX_Latitude          = regexp.MustCompile(Latitude)
	RX_Longitude         = regexp.MustCompile(Longitude)
	RX_DNSName           = regexp.MustCompile(DNSName)
	RX_URL               = regexp.MustCompile(URL)
	RX_SSN               = regexp.MustCompile(SSN)
	RX_WinPath           = regexp.MustCompile(WinPath)
	RX_UnixPath          = regexp.MustCompile(UnixPath)
	RX_ARWinPath         = regexp.MustCompile(WinARPath)
	RX_ARUnixPath        = regexp.MustCompile(UnixARPath)
	RX_Semver            = regexp.MustCompile(Semver)
	RX_HasLowerCase      = regexp.MustCompile(hasLowerCase)
	RX_HasUpperCase      = regexp.MustCompile(hasUpperCase)
	RX_HasWhitespace     = regexp.MustCompile(hasWhitespace)
	RX_HasWhitespaceOnly = regexp.MustCompile(hasWhitespaceOnly)
	RX_IMEI              = regexp.MustCompile(IMEI)
	RX_IMSI              = regexp.MustCompile(IMSI)
	RX_E164              = regexp.MustCompile(E164)

	RX_userRegexp    = regexp.MustCompile("^[a-zA-Z0-9!#$%&'*+/=?^_`{|}~.-]+$")
	RX_hostRegexp    = regexp.MustCompile("^[^\\s]+\\.[^\\s]+$")
	RX_userDotRegexp = regexp.MustCompile("(^[.]{1})|([.]{1}$)|([.]{2,})")
)
