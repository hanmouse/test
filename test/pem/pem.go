package main

import (
	"bufio"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
)

// GenSymmetricKey ...
func GenSymmetricKey(bits int) (k []byte, err error) {
	if bits <= 0 || bits%8 != 0 {
		return nil, errors.New("Key size error")
	}

	size := bits / 8
	k = make([]byte, size)
	if _, err = rand.Read(k); err != nil {
		return nil, err
	}

	return k, nil
}

func rsaTest() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	publicKey := &privateKey.PublicKey
	fmt.Println("Private Key: ", privateKey)
	fmt.Println("Public key: ", publicKey)

	const privateKeyFileName = "private_key.pem"

	pemPrivateFile, err := os.Create(privateKeyFileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer pemPrivateFile.Close()

	var pemPrivateBlock = &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}

	err = pem.Encode(pemPrivateFile, pemPrivateBlock)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	privateKeyFile, err := os.Open(privateKeyFileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	pemfileinfo, _ := privateKeyFile.Stat()
	var size int64 = pemfileinfo.Size()
	pembytes := make([]byte, size)
	buffer := bufio.NewReader(privateKeyFile)
	_, err = buffer.Read(pembytes)
	data, _ := pem.Decode([]byte(pembytes))
	privateKeyFile.Close()

	privateKeyImported, err := x509.ParsePKCS1PrivateKey(data.Bytes)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Private Key Imported: ", privateKeyImported)
}

func hmacTest() {
	key, err := GenSymmetricKey(256)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("[before] hmac key: %#v\n", key)

	const privateKeyFileName = "hmac_private_key.pem"

	privateKeyFile, err := os.Create(privateKeyFileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer privateKeyFile.Close()

	var block = &pem.Block{
		Type:  "HMAC PRIVATE KEY",
		Bytes: key,
	}

	err = pem.Encode(privateKeyFile, block)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	privateKeyFile, err = os.Open(privateKeyFileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer privateKeyFile.Close()

	pemfileinfo, _ := privateKeyFile.Stat()
	var size int64 = pemfileinfo.Size()
	pembytes := make([]byte, size)
	buffer := bufio.NewReader(privateKeyFile)
	_, err = buffer.Read(pembytes)
	data, _ := pem.Decode([]byte(pembytes))

	fmt.Printf("[after] hmac key: %#v\n", data.Bytes)
}

func main() {
	//rsaTest()
	hmacTest()
}
