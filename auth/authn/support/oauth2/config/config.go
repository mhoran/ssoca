package config

import (
	"crypto/rsa"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dpb587/ssoca/auth"
)

var JWTSigningMethod = jwt.SigningMethodRS256

const CookieStateName = "ssoca_oauth_state"
const CookieClientPortName = "ssoca_oauth_clientport"
const CookieClientVersionName = "ssoca_client_version"

type UserProfileLoader func(*http.Client) (auth.Token, error)

type JWT struct {
	PrivateKey   rsa.PrivateKey
	Validity     time.Duration
	ValidityPast time.Duration
}

type URLs struct {
	Origin      string
	AuthSuccess string
	AuthFailure string
}
