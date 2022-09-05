package signer

import (
	"github.com/nine-monsters/go-hutool/hutool-core/util"
	"github.com/nine-monsters/go-hutool/hutool-crypto/asymmetric"
	"github.com/nine-monsters/go-hutool/hutool-crypto/digest"
	"strings"
)

var dataMap = make(map[string]string, 12)

func init() {
	dataMap["HS256"] = digest.HmacSHA256.GetValue()
	dataMap["HS384"] = digest.HmacSHA384.GetValue()
	dataMap["HS512"] = digest.HmacSHA512.GetValue()
	dataMap[digest.HmacSHA256.GetValue()] = "HS256"
	dataMap[digest.HmacSHA384.GetValue()] = "HS384"
	dataMap[digest.HmacSHA512.GetValue()] = "HS512"

	dataMap["RS256"] = asymmetric.SHA256withRSA.GetValue()
	dataMap["RS384"] = asymmetric.SHA384withRSA.GetValue()
	dataMap["RS512"] = asymmetric.SHA512withRSA.GetValue()
	dataMap[asymmetric.SHA256withRSA.GetValue()] = "RS256"
	dataMap[asymmetric.SHA384withRSA.GetValue()] = "RS384"
	dataMap[asymmetric.SHA512withRSA.GetValue()] = "RS512"

	dataMap["ES256"] = asymmetric.SHA256withECDSA.GetValue()
	dataMap["ES384"] = asymmetric.SHA384withECDSA.GetValue()
	dataMap["ES512"] = asymmetric.SHA512withECDSA.GetValue()
	dataMap[asymmetric.SHA256withECDSA.GetValue()] = "ES256"
	dataMap[asymmetric.SHA384withECDSA.GetValue()] = "ES384"
	dataMap[asymmetric.SHA512withECDSA.GetValue()] = "ES512"

	dataMap["PS256"] = asymmetric.SHA256withRSA_PSS.GetValue()
	dataMap["PS384"] = asymmetric.SHA384withRSA_PSS.GetValue()
	dataMap["PS512"] = asymmetric.SHA512withRSA_PSS.GetValue()
	dataMap[asymmetric.SHA256withRSA_PSS.GetValue()] = "PS256"
	dataMap[asymmetric.SHA384withRSA_PSS.GetValue()] = "PS384"
	dataMap[asymmetric.SHA512withRSA_PSS.GetValue()] = "PS512"
}

func GetAlgorithm(idOrAlgorithm string) string {
	return util.DefaultIfEmpty(GetAlgorithmById(idOrAlgorithm), idOrAlgorithm)
}
func GetId(idOrAlgorithm string) string {
	return util.DefaultIfEmpty(GetIdByAlgorithm(idOrAlgorithm), idOrAlgorithm)
}

func GetAlgorithmById(id string) string {
	return dataMap[strings.ToUpper(id)]
}

func GetIdByAlgorithm(algorithm string) string {
	return dataMap[algorithm]
}
