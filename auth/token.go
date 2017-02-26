package auth

type TokenAttribute string

const (
	TokenSurnameAttribute   TokenAttribute = "surname"
	TokenGivenNameAttribute TokenAttribute = "given_name"
	TokenNameAttribute      TokenAttribute = "name"
	TokenEmailAttribute     TokenAttribute = "email"
	TokenUsernameAttribute  TokenAttribute = "username"
)

type Token struct {
	ID         string
	Groups     []string
	Attributes map[TokenAttribute]*string
}

func (t Token) Name() string {
	return *t.Attributes[TokenNameAttribute]
}

func (t Token) Email() string {
	return *t.Attributes[TokenEmailAttribute]
}

func (t Token) Username() string {
	return *t.Attributes[TokenUsernameAttribute]
}