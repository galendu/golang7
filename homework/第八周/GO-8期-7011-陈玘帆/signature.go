package main

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func ReadFile(keyFile string) ([]byte, error) {
	f, err := os.Open(keyFile)
	if err != nil {
		return nil, err
	}
	content := make([]byte, 4096)
	if n, err := f.Read(content); err != nil {
		return nil, err
	} else {
		return content[:n], nil
	}
}

func DigitalSignature(trade string) []byte {
	sha1 := sha1.New()
	sha1.Write([]byte(trade))
	digest := sha1.Sum([]byte(""))

	privateKey, _ := ReadFile("/Users/sy/go_project/ras_private_key.pem")
	block, _ := pem.Decode(privateKey)
	priv, _ := x509.ParsePKCS1PrivateKey(block.Bytes)

	signature, _ := rsa.SignPKCS1v15(nil, priv, crypto.Hash(0), digest)
	return signature
}

func VerifySignature(trade string, signature []byte) bool {
	sha1 := sha1.New()
	sha1.Write([]byte(trade))
	digest := sha1.Sum([]byte(""))

	publicKey, _ := ReadFile("/Users/sy/go_project/rsa_public_key.pem")
	block, _ := pem.Decode(publicKey)
	pubInterface, _ := x509.ParsePKIXPublicKey(block.Bytes)

	pub := pubInterface.(*rsa.PublicKey)

	return rsa.VerifyPKCS1v15(pub, crypto.Hash(0), digest, signature) == nil
}


func main() {
	trade := "测试一下数字签名"
	signature := DigitalSignature(trade)

	fmt.Println(signature)
	fmt.Println(VerifySignature(trade, signature))
}
