package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func getCurrUnixTimeMsec() int64 {
	return time.Now().UnixNano() / 1000000
}

var signingKeyForHMAC = []byte("6IkpXVCJ9.eyJzdWIiOiJ0ZX")
var signingKeyForRSA *rsa.PrivateKey = generateSigningKeyForRSA()
var verifyingKeyForRSA *rsa.PublicKey = getVerifyingKeyForRSA(signingKeyForRSA)

// 틀린 비밀키
var wrongSigningKey = []byte("6IkpXVCJ9.eyJzdWIiOiJ0Zx")

var pastTime int64 = time.Now().Unix() - 1
var currentTime int64 = time.Now().Unix()
var futureTime int64 = time.Now().Add(time.Second * 1).Unix()

func generateSigningKeyForRSA() *rsa.PrivateKey {
	privateKey, _ := rsa.GenerateKey(rand.Reader, 1024)
	return privateKey
}

func getVerifyingKeyForRSA(privateKey *rsa.PrivateKey) *rsa.PublicKey {
	bytes, _ := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	pem := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: bytes,
	})

	fmt.Printf("pem: %#v\n", string(pem))
	/*
			pem = []byte(`
		-----BEGIN RSA PUBLIC KEY-----
		MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCTgCSvm1OdS5SRHdmStt1NdEpw
		XDTiSzS61fp6/GPtILq8y69hMNBf0ZCdlfunv/s0zjVpLpuS9lrm31JPnAh7QHRU
		Sm/GkM6ucolrjZNJFSZ8ukri3WxyQdnEJqNWlDBy367KzLq4ZGp39R6rJFcnNKyt
		km/pJrbzm1wdNk3ajQIDAQAB
		-----END RSA PUBLIC KEY-----`)
	*/

	publicKey, _ := jwt.ParseRSAPublicKeyFromPEM(pem)

	return publicKey
}

func standardClaimsTest() {

	claims := &jwt.StandardClaims{
		Audience:  "Uangel",
		ExpiresAt: futureTime,
		Id:        "hanmouse",
		IssuedAt:  currentTime,
		Issuer:    "test",
		NotBefore: pastTime,
		Subject:   "hanmouse test",
	}

	//token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	fmt.Printf("token: %#v\n", token)

	//signingKey := signingKeyForHMAC
	signingKey := signingKeyForRSA

	signedJWT, err := token.SignedString(signingKey)
	if err == nil {
		fmt.Printf("encoded JWT: %#v\n", signedJWT)
		tokens := strings.Split(signedJWT, ".")
		for i, v := range tokens {
			fmt.Printf("JWT[%#v]: %#v\n", i, v)
		}
	} else {
		fmt.Println(err)
	}

	parsedToken, err := jwt.Parse(signedJWT, func(token *jwt.Token) (interface{}, error) {
		/*
			아래는 Type Assertion으로,
			token.Method가 *jwt.SigningMethodHMAC 타입이 아니면 ok가 false로 설정된다.
		*/
		//if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		verifyingKey := verifyingKeyForRSA
		//verifyingKey := signingKeyForHMAC

		return verifyingKey, nil
		//return wrongSigningKey, nil
	})

	if err != nil {
		fmt.Println(err.Error())
	}

	if parsedToken != nil && parsedToken.Valid {
		fmt.Println("token is valid")
		claims := parsedToken.Claims.(jwt.MapClaims)
		fmt.Printf("claims: %#v\n", claims)

	} else {
		fmt.Println(err)
	}

	/*
		if parsedToken.Valid {
			fmt.Printf("Token is valid\n")
		} else if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				fmt.Printf("Token is expired\n")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				fmt.Printf("Token is not valid yet\n")
			} else {
				fmt.Printf("Unknown error\n")
			}
		} else {
			fmt.Println("Couldn't handle this token: ", err)
		}

		if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
			fmt.Printf("%#v\n", claims)
		} else {
			fmt.Println(err)
		}
	*/
}

// TODO
func customClaimsTest() {

	fmt.Println("[customClaimsTest]")

	type MyCustomClaims struct {
		jwt.StandardClaims
		Scope string `json:"scope"`
	}

	inputClaims := MyCustomClaims{
		jwt.StandardClaims{
			Audience:  "Uangel",
			ExpiresAt: futureTime,
			Id:        "hanmouse_jti_1",
			IssuedAt:  currentTime,
			Issuer:    "test",
			NotBefore: pastTime,
			Subject:   "hanmouse test",
		},
		"smsf-sms",
	}

	encodedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, inputClaims)

	signingKey := signingKeyForHMAC
	signedJWT, _ := encodedToken.SignedString(signingKey)
	fmt.Printf("encodedToken: %#v\n", signedJWT)

	decodedToken, err := jwt.ParseWithClaims(signedJWT, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})

	if claims, ok := decodedToken.Claims.(*MyCustomClaims); ok && decodedToken.Valid {
		fmt.Printf("claims: %#v\n", claims)
		/*
			logger.Debug("verifyAccessToken: iss[ %#v ]", claims["iss"])
			logger.Debug("verifyAccessToken: sub[ %#v ]", claims["sub"])
			logger.Debug("verifyAccessToken: aud[ %#v ]", claims["aud"])
			logger.Debug("verifyAccessToken: scope[ %#v ]", claims["scope"])
			t := time.Unix(int64(claims["iat"].(float64)), 0)
			logger.Debug("verifyAccessToken: iat[ %#v ]", t.String())
			t = time.Unix(int64(claims["exp"].(float64)), 0)
			logger.Debug("verifyAccessToken: exp[ %#v ]", t.String())
		*/
	} else {
		fmt.Println(err)
	}
}

func main() {
	standardClaimsTest()
	//customClaimsTest()
}
