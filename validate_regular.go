package validate

import (
	"regexp"
)

// Basic regular expressions for validating strings.
// (there are from package "asaskevich/govalidator")
const (
	Email        = "^(((([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+(\\.([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+)*)|((\\x22)((((\\x20|\\x09)*(\\x0d\\x0a))?(\\x20|\\x09)+)?(([\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x7f]|\\x21|[\\x23-\\x5b]|[\\x5d-\\x7e]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(\\([\\x01-\\x09\\x0b\\x0c\\x0d-\\x7f]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}]))))*(((\\x20|\\x09)*(\\x0d\\x0a))?(\\x20|\\x09)+)?(\\x22)))@((([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])([a-zA-Z]|\\d|-|\\.|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.)+(([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])([a-zA-Z]|\\d|-|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.?$"
	UUID3        = "^[0-9a-f]{8}-[0-9a-f]{4}-3[0-9a-f]{3}-[0-9a-f]{4}-[0-9a-f]{12}$"
	UUID4        = "^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$"
	UUID5        = "^[0-9a-f]{8}-[0-9a-f]{4}-5[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$"
	UUID         = "^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$"
	Int          = "^(?:[-+]?(?:0|[1-9][0-9]*))$"
	Float        = "^(?:[-+]?(?:[0-9]+))?(?:\\.[0-9]*)?(?:[eE][\\+\\-]?(?:[0-9]+))?$"
	RGBColor     = "^rgb\\(\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*\\)$"
	FullWidth    = "[^\u0020-\u007E\uFF61-\uFF9F\uFFA0-\uFFDC\uFFE8-\uFFEE0-9a-zA-Z]"
	HalfWidth    = "[\u0020-\u007E\uFF61-\uFF9F\uFFA0-\uFFDC\uFFE8-\uFFEE0-9a-zA-Z]"
	Base64       = "^(?:[A-Za-z0-9+\\/]{4})*(?:[A-Za-z0-9+\\/]{2}==|[A-Za-z0-9+\\/]{3}=|[A-Za-z0-9+\\/]{4})$"
	Latitude     = "^[-+]?([1-8]?\\d(\\.\\d+)?|90(\\.0+)?)$"
	Longitude    = "^[-+]?(180(\\.0+)?|((1[0-7]\\d)|([1-9]?\\d))(\\.\\d+)?)$"
	DNSName      = `^([a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62}){1}(\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*[\._]?$`
	FullURL      = `^(?:ftp|tcp|udp|wss?|https?):\/\/[\w\.\/#=?&]+$`
	URLSchema    = `((ftp|tcp|udp|wss?|https?):\/\/)`
	URLUsername  = `(\S+(:\S*)?@)`
	URLPath      = `((\/|\?|#)[^\s]*)`
	URLPort      = `(:(\d{1,5}))`
	URLIP        = `([1-9]\d?|1\d\d|2[01]\d|22[0-3])(\.(1?\d{1,2}|2[0-4]\d|25[0-5])){2}(?:\.([0-9]\d?|1\d\d|2[0-4]\d|25[0-4]))`
	URLSubdomain = `((www\.)|([a-zA-Z0-9]+([-_\.]?[a-zA-Z0-9])*[a-zA-Z0-9]\.[a-zA-Z0-9]+))`
	WinPath      = `^[a-zA-Z]:\\(?:[^\\/:*?"<>|\r\n]+\\)*[^\\/:*?"<>|\r\n]*$`
	UnixPath     = `^(/[^/\x00]*)+/?$`
)

// some string regexp. (it is from package "asaskevich/govalidator")
var (
	// rxUser           = regexp.MustCompile("^[a-zA-Z0-9!#$%&'*+/=?^_`{|}~.-]+$")
	// rxHostname       = regexp.MustCompile("^[^\\s]+\\.[^\\s]+$")
	// rxUserDot        = regexp.MustCompile("(^[.]{1})|([.]{1}$)|([.]{2,})")
	rxEmail     = regexp.MustCompile(Email)
	rxISBN10    = regexp.MustCompile("^(?:[0-9]{9}X|[0-9]{10})$")
	rxISBN13    = regexp.MustCompile("^(?:[0-9]{13})$")
	rxUUID3     = regexp.MustCompile(UUID3)
	rxUUID4     = regexp.MustCompile(UUID4)
	rxUUID5     = regexp.MustCompile(UUID5)
	rxUUID      = regexp.MustCompile(UUID)
	rxAlpha     = regexp.MustCompile("^[a-zA-Z]+$")
	rxAlphaNum  = regexp.MustCompile("^[a-zA-Z0-9]+$")
	rxAlphaDash = regexp.MustCompile(`^(?:[\w-]+)$`)
	rxNumber    = regexp.MustCompile("^[0-9]+$")
	rxInt       = regexp.MustCompile(Int)
	rxFloat     = regexp.MustCompile(Float)
	rxCnMobile  = regexp.MustCompile(`^1\d{10}$`)
	rxHexColor  = regexp.MustCompile("^#?([0-9a-fA-F]{3}|[0-9a-fA-F]{6})$")
	rxRGBColor  = regexp.MustCompile(RGBColor)
	rxASCII     = regexp.MustCompile("^[\x00-\x7F]+$")
	// --
	rxHexadecimal    = regexp.MustCompile("^[0-9a-fA-F]+$")
	rxPrintableASCII = regexp.MustCompile("^[\x20-\x7E]+$")
	rxMultiByte      = regexp.MustCompile("[^\x00-\x7F]")
	// rxFullWidth      = regexp.MustCompile(FullWidth)
	// rxHalfWidth      = regexp.MustCompile(HalfWidth)
	rxBase64    = regexp.MustCompile(Base64)
	rxDataURI   = regexp.MustCompile(`^data:.+/(.+);base64,(?:.+)`)
	rxLatitude  = regexp.MustCompile(Latitude)
	rxLongitude = regexp.MustCompile(Longitude)
	rxDNSName   = regexp.MustCompile(DNSName)
	rxFullURL   = regexp.MustCompile(FullURL)
	rxURLSchema = regexp.MustCompile(URLSchema)
	// rxSSN            = regexp.MustCompile(`^\d{3}[- ]?\d{2}[- ]?\d{4}$`)
	rxWinPath  = regexp.MustCompile(WinPath)
	rxUnixPath = regexp.MustCompile(UnixPath)
	// --
	rxHasLowerCase = regexp.MustCompile(".*[[:lower:]]")
	rxHasUpperCase = regexp.MustCompile(".*[[:upper:]]")
)

/*************************************************************
 * global validators
 *************************************************************/

// global validators. contains built-in and user custom

