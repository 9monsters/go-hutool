package validator

const (
	// General 英文字母 、数字和下划线
	General string = "^\\w+$"

	// Numbers 数字
	Numbers string = "\\d+"

	// Word 字母
	Word string = "[a-zA-Z]+"

	// Chinese 单个中文汉字
	Chinese string = "[\u2E80-\u2EFF\u2F00-\u2FDF\u31C0-\u31EF\u3400-\u4DBF\u4E00-\u9FFF\uF900-\uFAFF]"

	// Chineses 中文汉字
	Chineses string = Chinese + "+"

	// GroupVar 分组
	GroupVar string = "\\$(\\d+)"

	// IPV4 IP v4
	IPV4 string = "^(25[0-5]|2[0-4]\\d|[0-1]?\\d?\\d)\\.(25[0-5]|2[0-4]\\d|[0-1]?\\d?\\d)\\.(25[0-5]|2[0-4]\\d|[0-1]?\\d?\\d)\\.(25[0-5]|2[0-4]\\d|[0-1]?\\d?\\d)$"

	// IPV6 IP v6
	IPV6 string = "(([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]+|::(ffff(:0{1,4})?:)?((25[0-5]|(2[0-4]|1?[0-9])?[0-9])\\.){3}(25[0-5]|(2[0-4]|1?[0-9])?[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1?[0-9])?[0-9])\\.){3}(25[0-5]|(2[0-4]|1?[0-9])?[0-9]))"

	// MONEY 	 * 货币
	MONEY string = "^(\\d+(?:\\.\\d+)?)$"

	//EMAIL 邮件，符合RFC 5322规范，正则来自：http://emailregex.com/
	// What is the maximum length of a valid email address? https://stackoverflow.com/questions/386294/what-is-the-maximum-length-of-a-valid-email-address/44317754
	// 注意email 要宽松一点。比如 jetz.chong@hutool.cn、jetz-chong@ hutool.cn、jetz_chong@hutool.cn、dazhi.duan@hutool.cn 宽松一点把，都算是正常的邮箱
	//
	EMAIL string = "(?:[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*|\"(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x21\\x23-\\x5b\\x5d-\\x7f]|\\\\[\\x01-\\x09\\x0b\\x0c\\x0e-\\x7f])*\")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\\[(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?|[a-z0-9-]*[a-z0-9]:(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x21-\\x5a\\x53-\\x7f]|\\\\[\\x01-\\x09\\x0b\\x0c\\x0e-\\x7f])+)])"

	// MOBILE 移动电话
	MOBILE string = "(?:0|86|\\+86)?1[3-9]\\d{9}"

	// MOBILE_HK 中国香港移动电话
	// eg: 中国香港： +852 5100 4810， 三位区域码+10位数字, 中国香港手机号码8位数
	// eg: 中国大陆： +86  180 4953 1399，2位区域码标示+13位数字
	// 中国大陆 +86 Mainland China
	// 中国香港 +852 Hong Kong
	// 中国澳门 +853 Macao
	// 中国台湾 +886 Taiwan
	MOBILE_HK string = "(?:0|852|\\+852)?\\d{8}"

	// MOBILE_TW 中国台湾移动电话
	// eg: 中国台湾： +886 09 60 000000， 三位区域码+号码以数字09开头 + 8位数字, 中国台湾手机号码10位数
	// 中国台湾 +886 Taiwan 国际域名缩写：TW
	MOBILE_TW string = "(?:0|886|\\+886)?(?:|-)09\\d{8}"

	// MOBILE_MO 中国澳门移动电话
	// eg: 中国台湾： +853 68 00000， 三位区域码 +号码以数字6开头 + 7位数字, 中国台湾手机号码8位数
	// 中国澳门 +853 Macao 国际域名缩写：MO
	//
	MOBILE_MO string = "(?:0|853|\\+853)?(?:|-)6\\d{7}"

	// TEL 座机号码
	//  pr#387@Gitee
	TEL string = "(010|02\\d|0[3-9]\\d{2})-?(\\d{6,8})"

	// TEL_400_800 座机号码+400+800电话
	// @see <a href="https://baike.baidu.com/item/800">800</a>
	TEL_400_800 string = "0\\d{2,3}[\\- ]?[1-9]\\d{6,7}|[48]00[\\- ]?[1-9]\\d{6}"

	// CITIZEN_ID 18位身份证号码
	CITIZEN_ID string = "[1-9]\\d{5}[1-2]\\d{3}((0\\d)|(1[0-2]))(([012]\\d)|3[0-1])\\d{3}(\\d|X|x)"

	// ZIP_CODE 邮编，兼容港澳台
	ZIP_CODE string = "^(0[1-7]|1[0-356]|2[0-7]|3[0-6]|4[0-7]|5[0-7]|6[0-7]|7[0-5]|8[0-9]|9[0-8])\\d{4}|99907[78]$"

	// BIRTHDAY 生日
	BIRTHDAY string = "^(\\d{2,4})([/\\-.年]?)(\\d{1,2})([/\\-.月]?)(\\d{1,2})日?$"

	// URI URI
	// 定义见：https://www.ietf.org/rfc/rfc3986.html#appendix-B
	URI string = "^(([^:/?#]+):)?(//([^/?#]*))?([^?#]*)(\\?([^#]*))?(#(.*))?"

	// URL_HTTP Http URL（来自：http://urlregex.com/）
	// 此正则同时支持FTP、File等协议的URL
	URL_HTTP string = "(https?|ftp|file)://[\\w-+&@#/%?=~_|!:,.;]*[\\w-+&@#/%=~_|]"

	// GENERAL_WITH_CHINESE 中文字、英文字母、数字和下划线
	GENERAL_WITH_CHINESE string = "^[\u4E00-\u9FFF\\w]+$"

	// UUID_SIMPLE 不带横线的UUID
	UUID_SIMPLE string = "^[0-9a-fA-F]{32}$"

	// MAC_ADDRESS MAC地址正则
	MAC_ADDRESS string = "((?:[a-fA-F0-9]{1,2}[:-]){5}[a-fA-F0-9]{1,2})|0x(\\d{12}).+ETHER"

	// Hex 16进制字符串
	Hex string = "^[a-fA-F0-9]+$"

	// TIME 时间正则
	TIME string = "\\d{1,2}:\\d{1,2}(:\\d{1,2})?"

	// PLATE_NUMBER 中国车牌号码（兼容新能源车牌）
	PLATE_NUMBER string = "^(([京津沪渝冀豫云辽黑湘皖鲁新苏浙赣鄂桂甘晋蒙陕吉闽贵粤青藏川宁琼使领][A-Z](([0-9]{5}[ABCDEFGHJK])|([ABCDEFGHJK]([A-HJ-NP-Z0-9])[0-9]{4})))|" +
		"([京津沪渝冀豫云辽黑湘皖鲁新苏浙赣鄂桂甘晋蒙陕吉闽贵粤青藏川宁琼使领]\\d{3}\\d{1,3}[领])|" +
		"([京津沪渝冀豫云辽黑湘皖鲁新苏浙赣鄂桂甘晋蒙陕吉闽贵粤青藏川宁琼使领][A-Z][A-HJ-NP-Z0-9]{4}[A-HJ-NP-Z0-9挂学警港澳使领]))$"

	// CREDIT_CODE 统一社会信用代码
	// 第一部分：登记管理部门代码1位 (数字或大写英文字母)
	// 第二部分：机构类别代码1位 (数字或大写英文字母)
	// 第三部分：登记管理机关行政区划码6位 (数字)
	// 第四部分：主体标识码（组织机构代码）9位 (数字或大写英文字母)
	// 第五部分：校验码1位 (数字或大写英文字母)
	CREDIT_CODE string = "^[0-9A-HJ-NPQRTUWXY]{2}\\d{6}[0-9A-HJ-NPQRTUWXY]{10}$"

	// CAR_VIN 车架号
	// 别名：车辆识别代号 车辆识别码
	// eg:LDC613P23A1305189
	// eg:LSJA24U62JG269225
	// 十七位码、车架号
	// 车辆的唯一标示
	CAR_VIN string = "^[A-HJ-NPR-Z0-9]{8}[0-9X][A-HJ-NPR-Z0-9]{2}\\d{6}$"

	// CAR_DRIVING_LICENCE 驾驶证  别名：驾驶证档案编号、行驶证编号
	// eg:430101758218
	// 12位数字字符串
	// 仅限：中国驾驶证档案编号
	CAR_DRIVING_LICENCE string = "^[0-9]{12}$"

	// CHINESE_NAME 中文姓名
	// 维吾尔族姓名里面的点是 · 输入法中文状态下，键盘左上角数字1前面的那个符号；
	// 错误字符：{@code ．.。．.}
	// 正确维吾尔族姓名：
	// 霍加阿卜杜拉·麦提喀斯木
	// 玛合萨提别克·哈斯木别克
	// 阿布都热依木江·艾斯卡尔
	// 阿卜杜尼亚孜·毛力尼亚孜
	// ----------
	// 错误示例：孟  伟                reason: 有空格
	// 错误示例：连逍遥0               reason: 数字
	// 错误示例：依帕古丽-艾则孜        reason: 特殊符号
	// 错误示例：牙力空.买提萨力        reason: 新疆人的点不对
	// 错误示例：王建鹏2002-3-2        reason: 有数字、特殊符号
	// 错误示例：雷金默(雷皓添）        reason: 有括号
	// 错误示例：翟冬:亮               reason: 有特殊符号
	// 错误示例：李                   reason: 少于2位
	// ----------
	// 总结中文姓名：2-60位，只能是中文和维吾尔族的点·
	// 放宽汉字范围：如生僻姓名 刘欣䶮yǎn
	//
	CHINESE_NAME string = "^[\u2E80-\u9FFF·]{2,60}$"

	Email             string = "^(((([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+(\\.([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+)*)|((\\x22)((((\\x20|\\x09)*(\\x0d\\x0a))?(\\x20|\\x09)+)?(([\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x7f]|\\x21|[\\x23-\\x5b]|[\\x5d-\\x7e]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(\\([\\x01-\\x09\\x0b\\x0c\\x0d-\\x7f]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}]))))*(((\\x20|\\x09)*(\\x0d\\x0a))?(\\x20|\\x09)+)?(\\x22)))@((([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])([a-zA-Z]|\\d|-|\\.|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.)+(([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])([a-zA-Z]|\\d|-|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.?$"
	CreditCard        string = "^(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|(222[1-9]|22[3-9][0-9]|2[3-6][0-9]{2}|27[01][0-9]|2720)[0-9]{12}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\\d{3})\\d{11}|6[27][0-9]{14})$"
	ISBN10            string = "^(?:[0-9]{9}X|[0-9]{10})$"
	ISBN13            string = "^(?:[0-9]{13})$"
	UUID3             string = "^[0-9a-f]{8}-[0-9a-f]{4}-3[0-9a-f]{3}-[0-9a-f]{4}-[0-9a-f]{12}$"
	UUID4             string = "^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$"
	UUID5             string = "^[0-9a-f]{8}-[0-9a-f]{4}-5[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$"
	UUID              string = "^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$"
	Alpha             string = "^[a-zA-Z]+$"
	Alphanumeric      string = "^[a-zA-Z0-9]+$"
	Numeric           string = "^[0-9]+$"
	Int               string = "^(?:[-+]?(?:0|[1-9][0-9]*))$"
	Float             string = "^(?:[-+]?(?:[0-9]+))?(?:\\.[0-9]*)?(?:[eE][\\+\\-]?(?:[0-9]+))?$"
	Hexadecimal       string = "^[0-9a-fA-F]+$"
	Hexcolor          string = "^#?([0-9a-fA-F]{3}|[0-9a-fA-F]{6})$"
	RGBcolor          string = "^rgb\\(\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*\\)$"
	ASCII             string = "^[\x00-\x7F]+$"
	Multibyte         string = "[^\x00-\x7F]"
	FullWidth         string = "[^\u0020-\u007E\uFF61-\uFF9F\uFFA0-\uFFDC\uFFE8-\uFFEE0-9a-zA-Z]"
	HalfWidth         string = "[\u0020-\u007E\uFF61-\uFF9F\uFFA0-\uFFDC\uFFE8-\uFFEE0-9a-zA-Z]"
	Base64            string = "^(?:[A-Za-z0-9+\\/]{4})*(?:[A-Za-z0-9+\\/]{2}==|[A-Za-z0-9+\\/]{3}=|[A-Za-z0-9+\\/]{4})$"
	PrintableASCII    string = "^[\x20-\x7E]+$"
	DataURI           string = "^data:.+\\/(.+);base64$"
	MagnetURI         string = "^magnet:\\?xt=urn:[a-zA-Z0-9]+:[a-zA-Z0-9]{32,40}&dn=.+&tr=.+$"
	Latitude          string = "^[-+]?([1-8]?\\d(\\.\\d+)?|90(\\.0+)?)$"
	Longitude         string = "^[-+]?(180(\\.0+)?|((1[0-7]\\d)|([1-9]?\\d))(\\.\\d+)?)$"
	DNSName           string = `^([a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62}){1}(\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*[\._]?$`
	IP                string = `(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))`
	URLSchema         string = `((ftp|tcp|udp|wss?|https?):\/\/)`
	URLUsername       string = `(\S+(:\S*)?@)`
	URLPath           string = `((\/|\?|#)[^\s]*)`
	URLPort           string = `(:(\d{1,5}))`
	URLIP             string = `([1-9]\d?|1\d\d|2[01]\d|22[0-3]|24\d|25[0-5])(\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])){2}(?:\.([0-9]\d?|1\d\d|2[0-4]\d|25[0-5]))`
	URLSubdomain      string = `((www\.)|([a-zA-Z0-9]+([-_\.]?[a-zA-Z0-9])*[a-zA-Z0-9]\.[a-zA-Z0-9]+))`
	URL                      = `^` + URLSchema + `?` + URLUsername + `?` + `((` + URLIP + `|(\[` + IP + `\])|(([a-zA-Z0-9]([a-zA-Z0-9-_]+)?[a-zA-Z0-9]([-\.][a-zA-Z0-9]+)*)|(` + URLSubdomain + `?))?(([a-zA-Z\x{00a1}-\x{ffff}0-9]+-?-?)*[a-zA-Z\x{00a1}-\x{ffff}0-9]+)(?:\.([a-zA-Z\x{00a1}-\x{ffff}]{1,}))?))\.?` + URLPort + `?` + URLPath + `?$`
	SSN               string = `^\d{3}[- ]?\d{2}[- ]?\d{4}$`
	WinPath           string = `^[a-zA-Z]:\\(?:[^\\/:*?"<>|\r\n]+\\)*[^\\/:*?"<>|\r\n]*$`
	UnixPath          string = `^(/[^/\x00]*)+/?$`
	WinARPath         string = `^(?:(?:[a-zA-Z]:|\\\\[a-z0-9_.$●-]+\\[a-z0-9_.$●-]+)\\|\\?[^\\/:*?"<>|\r\n]+\\?)(?:[^\\/:*?"<>|\r\n]+\\)*[^\\/:*?"<>|\r\n]*$`
	UnixARPath        string = `^((\.{0,2}/)?([^/\x00]*))+/?$`
	Semver            string = "^v?(?:0|[1-9]\\d*)\\.(?:0|[1-9]\\d*)\\.(?:0|[1-9]\\d*)(-(0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*)(\\.(0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*))*)?(\\+[0-9a-zA-Z-]+(\\.[0-9a-zA-Z-]+)*)?$"
	tagName           string = "valid"
	hasLowerCase      string = ".*[[:lower:]]"
	hasUpperCase      string = ".*[[:upper:]]"
	hasWhitespace     string = ".*[[:space:]]"
	hasWhitespaceOnly string = "^[[:space:]]+$"
	IMEI              string = "^[0-9a-f]{14}$|^\\d{15}$|^\\d{18}$"
	IMSI              string = "^\\d{14,15}$"
	E164              string = `^\+?[1-9]\d{1,14}$`
)
