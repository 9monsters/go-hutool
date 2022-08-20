package jwt

type Payload struct {
	*Claims
}

func (p Payload) addPayloads() Payload {
	return p
}
