package jwt

import (
	"encoding/json"
	"sync"

	"github.com/nine-monsters/go-hutool/hutool-core/codec"
)

type Claims struct {
	claims map[string]any
}

var once sync.Once

type Claimer interface {
	SetClaim(name string, data any)
	GetClaim(name string) any
	PutAll(data map[string]any)
	Parse(tokenPart string)
}

func (c *Claims) SetClaim(name string, data any) {
	c.initClaims()
	if name == "" {
		panic("name can not null !")
	}
	if data == nil {
		delete(c.claims, name)
		return
	}
	c.claims[name] = data

}

func (c *Claims) GetClaim(name string) any {
	c.initClaims()
	if name == "" {
		panic("name can not null !")
	}
	return c.claims[name]

}

func (c *Claims) PutAll(data map[string]any) {
	c.initClaims()
	if data != nil && len(data) > 0 {
		for key, value := range data {
			c.SetClaim(key, value)
		}
	}
}

func (c *Claims) Parse(tokenPart string) {
	str, _ := codec.Base64.DecodeStr(tokenPart)
	mapData := make(map[string]any, 5)
	err := json.Unmarshal([]byte(str), &mapData)
	if err != nil {
		c.claims = mapData
	}
}

func (c *Claims) initClaims() {
	if c.claims == nil {
		once.Do(func() {
			c.claims = make(map[string]any, 5)
		})
	}
}

func (c *Claims) GetClaims() map[string]any {
	c.initClaims()
	return c.claims
}
