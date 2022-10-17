package lang

const (
	// GeneralReg 英文字母 、数字和下划线
	GeneralReg string = "^\\w+$"

	// NumbersReg 数字
	NumbersReg string = "\\d+"

	// WordReg 字母
	WordReg string = "[a-zA-Z]+"

	// ChineseReg 单个中文汉字
	ChineseReg string = "[\u2E80-\u2EFF\u2F00-\u2FDF\u31C0-\u31EF\u3400-\u4DBF\u4E00-\u9FFF\uF900-\uFAFF]"

	// ChinesesReg 中文汉字
	ChinesesReg string = ChineseReg + "+"

	// GroupVarReg 分组
	GroupVarReg string = "\\$(\\d+)"

	// IPV4_REG IP v4
	IPV4_REG string = "^(25[0-5]|2[0-4]\\d|[0-1]?\\d?\\d)\\.(25[0-5]|2[0-4]\\d|[0-1]?\\d?\\d)\\.(25[0-5]|2[0-4]\\d|[0-1]?\\d?\\d)\\.(25[0-5]|2[0-4]\\d|[0-1]?\\d?\\d)$"

	// IPV6_REG IP v6
	IPV6_REG string = "(([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]+|::(ffff(:0{1,4})?:)?((25[0-5]|(2[0-4]|1?[0-9])?[0-9])\\.){3}(25[0-5]|(2[0-4]|1?[0-9])?[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1?[0-9])?[0-9])\\.){3}(25[0-5]|(2[0-4]|1?[0-9])?[0-9]))"

	// MONEY_REG 	 * 货币
	MONEY_REG string = "^(\\d+(?:\\.\\d+)?)$"

	//EMAIL_REG 邮件，符合RFC 5322规范，正则来自：http://emailregex.com/
	// What is the maximum length of a valid email address? https://stackoverflow.com/questions/386294/what-is-the-maximum-length-of-a-valid-email-address/44317754
	// 注意email 要宽松一点。比如 jetz.chong@hutool.cn、jetz-chong@ hutool.cn、jetz_chong@hutool.cn、dazhi.duan@hutool.cn 宽松一点把，都算是正常的邮箱
	//
	EMAIL_REG string = "(?:[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*|\"(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x21\\x23-\\x5b\\x5d-\\x7f]|\\\\[\\x01-\\x09\\x0b\\x0c\\x0e-\\x7f])*\")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\\[(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?|[a-z0-9-]*[a-z0-9]:(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x21-\\x5a\\x53-\\x7f]|\\\\[\\x01-\\x09\\x0b\\x0c\\x0e-\\x7f])+)])"

	// MOBILE_REG 移动电话
	MOBILE_REG string = "(?:0|86|\\+86)?1[3-9]\\d{9}"

	// MOBILE_HK_REG 中国香港移动电话
	// eg: 中国香港： +852 5100 4810， 三位区域码+10位数字, 中国香港手机号码8位数
	// eg: 中国大陆： +86  180 4953 1399，2位区域码标示+13位数字
	// 中国大陆 +86 Mainland China
	// 中国香港 +852 Hong Kong
	// 中国澳门 +853 Macao
	// 中国台湾 +886 Taiwan
	MOBILE_HK_REG string = "(?:0|852|\\+852)?\\d{8}"

	// MOBILE_TW_REG 中国台湾移动电话
	// eg: 中国台湾： +886 09 60 000000， 三位区域码+号码以数字09开头 + 8位数字, 中国台湾手机号码10位数
	// 中国台湾 +886 Taiwan 国际域名缩写：TW
	MOBILE_TW_REG string = "(?:0|886|\\+886)?(?:|-)09\\d{8}"

	// MOBILE_MO_REG 中国澳门移动电话
	// eg: 中国台湾： +853 68 00000， 三位区域码 +号码以数字6开头 + 7位数字, 中国台湾手机号码8位数
	// 中国澳门 +853 Macao 国际域名缩写：MO
	//
	MOBILE_MO_REG string = "(?:0|853|\\+853)?(?:|-)6\\d{7}"

	// TEL_REG 座机号码
	//  pr#387@Gitee
	TEL_REG string = "(010|02\\d|0[3-9]\\d{2})-?(\\d{6,8})"

	// TEL_400_800_REG 座机号码+400+800电话
	// @see <a href="https://baike.baidu.com/item/800">800</a>
	TEL_400_800_REG string = "0\\d{2,3}[\\- ]?[1-9]\\d{6,7}|[48]00[\\- ]?[1-9]\\d{6}"

	// CITIZEN_ID_REG 18位身份证号码
	CITIZEN_ID_REG string = "[1-9]\\d{5}[1-2]\\d{3}((0\\d)|(1[0-2]))(([012]\\d)|3[0-1])\\d{3}(\\d|X|x)"

	// ZIP_CODE_REG 邮编，兼容港澳台
	ZIP_CODE_REG string = "^(0[1-7]|1[0-356]|2[0-7]|3[0-6]|4[0-7]|5[0-7]|6[0-7]|7[0-5]|8[0-9]|9[0-8])\\d{4}|99907[78]$"

	// BIRTHDAY_REG 生日
	BIRTHDAY_REG string = "^(\\d{2,4})([/\\-.年]?)(\\d{1,2})([/\\-.月]?)(\\d{1,2})日?$"

	// URI_REG URI
	// 定义见：https://www.ietf.org/rfc/rfc3986.html#appendix-B
	URI_REG string = "^(([^:/?#]+):)?(//([^/?#]*))?([^?#]*)(\\?([^#]*))?(#(.*))?"

	// URL_REG URL
	URL_REG string = "[a-zA-Z]+://[\\w-+&@#/%?=~_|!:,.;]*[\\w-+&@#/%=~_|]"

	// URL_HTTP_REG Http URL（来自：http://urlregex.com/）
	// 此正则同时支持FTP、File等协议的URL
	URL_HTTP_REG string = "(https?|ftp|file)://[\\w-+&@#/%?=~_|!:,.;]*[\\w-+&@#/%=~_|]"

	// GENERAL_WITH_CHINESE_REG 中文字、英文字母、数字和下划线
	GENERAL_WITH_CHINESE_REG string = "^[\u4E00-\u9FFF\\w]+$"

	// UUID_REG UUID
	UUID_REG string = "^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$"

	// UUID_SIMPLE_REG 不带横线的UUID
	UUID_SIMPLE_REG string = "^[0-9a-fA-F]{32}$"

	// MAC_ADDRESS_REG MAC地址正则
	MAC_ADDRESS_REG string = "((?:[a-fA-F0-9]{1,2}[:-]){5}[a-fA-F0-9]{1,2})|0x(\\d{12}).+ETHER"

	// HexReg 16进制字符串
	HexReg string = "^[a-fA-F0-9]+$"

	// TIME_REG 时间正则
	TIME_REG string = "\\d{1,2}:\\d{1,2}(:\\d{1,2})?"

	// PLATE_NUMBER_REG 中国车牌号码（兼容新能源车牌）
	PLATE_NUMBER_REG string = "^(([京津沪渝冀豫云辽黑湘皖鲁新苏浙赣鄂桂甘晋蒙陕吉闽贵粤青藏川宁琼使领][A-Z](([0-9]{5}[ABCDEFGHJK])|([ABCDEFGHJK]([A-HJ-NP-Z0-9])[0-9]{4})))|" +
		"([京津沪渝冀豫云辽黑湘皖鲁新苏浙赣鄂桂甘晋蒙陕吉闽贵粤青藏川宁琼使领]\\d{3}\\d{1,3}[领])|" +
		"([京津沪渝冀豫云辽黑湘皖鲁新苏浙赣鄂桂甘晋蒙陕吉闽贵粤青藏川宁琼使领][A-Z][A-HJ-NP-Z0-9]{4}[A-HJ-NP-Z0-9挂学警港澳使领]))$"

	// CREDIT_CODE_REG 统一社会信用代码
	// 第一部分：登记管理部门代码1位 (数字或大写英文字母)
	// 第二部分：机构类别代码1位 (数字或大写英文字母)
	// 第三部分：登记管理机关行政区划码6位 (数字)
	// 第四部分：主体标识码（组织机构代码）9位 (数字或大写英文字母)
	// 第五部分：校验码1位 (数字或大写英文字母)
	CREDIT_CODE_REG string = "^[0-9A-HJ-NPQRTUWXY]{2}\\d{6}[0-9A-HJ-NPQRTUWXY]{10}$"

	// CAR_VIN_REG 车架号
	// 别名：车辆识别代号 车辆识别码
	// eg:LDC613P23A1305189
	// eg:LSJA24U62JG269225
	// 十七位码、车架号
	// 车辆的唯一标示
	CAR_VIN_REG string = "^[A-HJ-NPR-Z0-9]{8}[0-9X][A-HJ-NPR-Z0-9]{2}\\d{6}$"

	// CAR_DRIVING_LICENCE_REG 驾驶证  别名：驾驶证档案编号、行驶证编号
	// eg:430101758218
	// 12位数字字符串
	// 仅限：中国驾驶证档案编号
	CAR_DRIVING_LICENCE_REG string = "^[0-9]{12}$"

	// CHINESE_NAME_REG 中文姓名
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
	CHINESE_NAME_REG string = "^[\u2E80-\u9FFF·]{2,60}$"
)
