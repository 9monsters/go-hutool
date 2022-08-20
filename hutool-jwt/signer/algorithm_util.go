package signer

import (
	"strings"

	"github.com/nine-monsters/go-hutool/hutool-core/util"
	"github.com/nine-monsters/go-hutool/hutool-crypto/asymmetric"
	"github.com/nine-monsters/go-hutool/hutool-crypto/digest"
)

var dataMap = make(map[string]string, 10)

func init() {
	dataMap["HS256"] = digest.HmacSHA256.GetValue()
	dataMap["HS384"] = digest.HmacSHA384.GetValue()
	dataMap["HS512"] = digest.HmacSHA512.GetValue()

	dataMap["RS256"] = asymmetric.SHA256withRSA.GetValue()
	dataMap["RS384"] = asymmetric.SHA384withRSA.GetValue()
	dataMap["RS512"] = asymmetric.SHA512withRSA.GetValue()

	dataMap["ES256"] = asymmetric.SHA256withECDSA.GetValue()
	dataMap["ES384"] = asymmetric.SHA384withECDSA.GetValue()
	dataMap["ES512"] = asymmetric.SHA512withECDSA.GetValue()

	dataMap["PS256"] = asymmetric.SHA256withRSA_PSS.GetValue()
	dataMap["PS384"] = asymmetric.SHA384withRSA_PSS.GetValue()
	dataMap["PS512"] = asymmetric.SHA512withRSA_PSS.GetValue()
}

func GetAlgorithm(idOrAlgorithm string) string {
	return util.DefaultIfNil(GetAlgorithmById(idOrAlgorithm), idOrAlgorithm)
}

func GetId(idOrAlgorithm string) string {
	return util.DefaultIfNil(GetIdByAlgorithm(idOrAlgorithm), idOrAlgorithm)
}

func GetAlgorithmById(algorithm string) string {
	return dataMap[strings.ToUpper(algorithm)]
}

func GetIdByAlgorithm(algorithm string) string {
	return dataMap[algorithm]
}
