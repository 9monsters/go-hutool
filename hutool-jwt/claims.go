package jwt

type Claims struct {
	claims map[string]any
}

type claimsInterface interface {
	SetClaim(name string, data any)
	GetClaim(name string) any
	PutAll(data map[string]any)
}

func claimJSON() {

}

func (c Claims) SetClaim(name string, data any) {
	if name == "" {
		panic("name can not null !")
	}
	if data == nil {
		delete(c.claims, name)
		return
	}
	c.claims[name] = data

}

func (c Claims) GetClaim(name string) any {
	if name == "" {
		panic("name can not null !")
	}
	return c.claims[name]

}

func (c Claims) PutAll(data map[string]any) {
	if data != nil && len(data) > 0 {
		for key, value := range data {
			c.SetClaim(key, value)
		}
	}
}
