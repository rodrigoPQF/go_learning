package jwk

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"math/big"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

type JWK struct {
	Kty string `json:"kty"`
	K   string `json:"k"`
	N   string `json:"n"`
}

type VerifyJWKtoJWT struct {
	jwk      JWK
	jwtToken string
	pemCert  string
}

func NewVerifyJWKtoJWT(verify VerifyJWKtoJWT) *VerifyJWKtoJWT {
	return &VerifyJWKtoJWT{
		jwk: JWK{
			Kty: verify.jwk.Kty,
			K:   verify.jwk.K,
			N:   verify.jwk.N,
		},
		jwtToken: verify.jwtToken,
	}
}

func (vjj *VerifyJWKtoJWT) ToPEM() (VerifyJWKtoJWT, error) {

	nBytes, err := base64.RawURLEncoding.DecodeString(vjj.jwk.N)
	if err != nil {
		return *vjj, err
	}

	key := &rsa.PublicKey{
		N: new(big.Int).SetBytes(nBytes),
		E: 65537,
	}

	pubKeyBytes, err := x509.MarshalPKIXPublicKey(key)
	if err != nil {
		return *vjj, err
	}

	pemCert := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubKeyBytes,
	})

	vjj.pemCert = string(pemCert)

	return *vjj, nil
}

func (vjj *VerifyJWKtoJWT) ValidateJWT() (bool, error) {
	key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(vjj.pemCert))
	if err != nil {
		return false, err
	}

	claims := jwt.MapClaims{}

	_, err = jwt.ParseWithClaims(vjj.jwtToken, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		if strings.Contains(err.Error(), "Token is expired") {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
