package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"math/big"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

type JWK struct {
	Kty string `json:"kty"`
	K   string `json:"k"`
	N   string `json:"n"`
}

func main() {
	jwkJSON := `{"kty":"RSA","k":"AQAB","n":"u1SU1LfVLPHCozMxH2Mo4lgOEePzNm0tRgeLezV6ffAt0gunVTLw7onLRnrq0_IzW7yWR7QkrmBL7jTKEn5u-qKhbwKfBstIs-bMY2Zkp18gnTxKLxoS2tFczGkPLPgizskuemMghRniWaoLcyehkd3qqGElvW_VDL5AaWTg0nLVkjRo9z-40RQzuVaE8AkAFmxZzow3x-VJYKdjykkJ0iT9wCS0DRTXu269V264Vf_3jvredZiKRkgwlL9xNAwxXFg0x_XFw005UWVRIkdgcKWTjpBP2dPwVZ4WWC-9aGVd-Gyn1o0CLelf4rEjGoXbAAEgAqeGUxrcIlbjXfbcmw"}`

	pemCert, err := JWKToPEM(jwkJSON)
	if err != nil {
		fmt.Println("Erro ao converter JWK para PEM:", err)
		return
	}

	fmt.Println("Certificado PEM:")
	fmt.Println(pemCert)

	token := "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTUxNjIzOTAyMn0.NHVaYe26MbtOYhSKkoKYdFVomg4i8ZJd8_-RU8VNbftc4TSMb4bXP3l3YlNWACwyXPGffz5aXHc6lty1Y2t4SWRqGteragsVdZufDn5BlnJl9pdR_kdVFUsra2rWKEofkZeIC4yWytE58sMIihvo9H1ScmmVwBcQP6XETqYd0aSHp1gOa9RdUPDvoXQ5oqygTqVtxaDr6wUFKrKItgBMzWIdNZ6y7O9E0DhEPTbE9rfBo6KTFsHAZnMg4k68CDp2woYIaXbmYTWcvbzIuHO7_37GT79XdIwkm95QJ7hYC9RiwrV7mesbY4PAahERJawntho0my942XheVLmGwLMBkQ"

	valid, err := ValidateJWT(token, pemCert)
	if err != nil {
		fmt.Println("Erro ao validar token JWT:", err)
		return
	}

	if valid {
		fmt.Println("Token JWT válido!")
	} else {
		fmt.Println("Token JWT inválido!")
	}
}

func JWKToPEM(jwkJSON string) (string, error) {
	var jwk JWK
	err := json.Unmarshal([]byte(jwkJSON), &jwk)
	if err != nil {
		return "", err
	}

	nBytes, err := base64.RawURLEncoding.DecodeString(jwk.N)
	if err != nil {
		return "", err
	}

	key := &rsa.PublicKey{
		N: new(big.Int).SetBytes(nBytes),
		E: 65537,
	}

	pubKeyBytes, err := x509.MarshalPKIXPublicKey(key)
	if err != nil {
		return "", err
	}

	pemCert := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubKeyBytes,
	})

	return string(pemCert), nil
}

func ValidateJWT(token string, pemCert string) (bool, error) {
	key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(pemCert))
	if err != nil {
		return false, err
	}

	claims := jwt.MapClaims{}

	_, err = jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
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
