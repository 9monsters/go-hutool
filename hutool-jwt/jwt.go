package jwt

import (
	"encoding/json"
	"fmt"
	"github.com/nine-monsters/go-hutool/hutool-jwt/signer"

	"github.com/nine-monsters/go-hutool/hutool-core/codec"
	//"github.com/nine-monsters/go-hutool/hutool-jwt/signer"
)

type Jwt struct {
	Raw       string        // The raw token.  Populated when you Parse a token
	Signer    signer.Signer // The signing method used or to be used
	Claims    Claims        // The second segment of the token
	Signature string        // The third segment of the token.  Populated when you Parse a token
	Valid     bool          // Is the token valid?  Populated when you Parse/Verify a token

	Header  Header  // The first segment of the token
	Payload Payload // Is the token valid?  Populated when you Parse/Verify a token
}

type jwter interface {
	setKey(key []byte) *Jwt
	setSigner(signer signer.Signer) *Jwt
	sign() string
	signWithSigner(signer signer.Signer) string
}

func Create() *Jwt {
	jwt := &Jwt{}
	jwt.Header = *createHeader()
	jwt.Payload = *createPayload()
	return jwt
}

func (j *Jwt) setKey(key []byte) *Jwt {
	j.setSigner(signer.HS256(key))
	return j
}
func (j *Jwt) setSigner(signer signer.Signer) *Jwt {
	j.Signer = signer
	return j
}

func (j *Jwt) sign() string {
	return j.signWithSigner(j.Signer)
}

func (j *Jwt) signWithSigner(sg signer.Signer) string {
	if sg == nil {
		panic("No Signer provided!")
	}
	alg := j.Header.getClaim(Algorithm)

	s, ok := alg.(string)
	// alg is empty
	if (ok && s == "") || (!ok && alg == nil) {
		j.Header.setClaim(Algorithm, signer.GetId(sg.GetAlgorithm()))
	}
	hJson, _ := json.Marshal(j.Header)
	hSafeStr := codec.Base64.EncodeUrlSafeStr(hJson)

	pJson, _ := json.Marshal(j.Payload)
	pSafeStr := codec.Base64.EncodeUrlSafeStr(pJson)

	sign := sg.Sign(hSafeStr, pSafeStr)
	return fmt.Sprintf("%s.%s.%s", hSafeStr, pSafeStr, sign)
}
