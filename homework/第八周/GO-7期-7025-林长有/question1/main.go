package main

import (
	"fmt"
	"crypto"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"crypto/sha1"
	"os"
)

// 1 实现数字签名
func ReadFile(keyFile string) ([]byte, error) {
	if f, err := os.Open(keyFile); err != nil {
		return nil, err
	} else {
		content := make([]byte, 4096)
		if n, err := f.Read(content); err != nil {
			return nil, err
		} else {
			return content[:n], nil
		}
	}
}

func DigitalSignature(trade string) []byte {
	sha1 := sha1.New()
	sha1.Write([]byte(trade))
	digest := sha1.Sum([]byte("")) // 双引号里面不能有空格

	privateKey, _ := ReadFile("./data/rsa_private_key.pem")
	block, _ := pem.Decode(privateKey)
	priv, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
	// 用私钥生成签名
	signature, _ := rsa.SignPKCS1v15(nil, priv, crypto.Hash(0), digest)
	return signature
}

/// 验证数字签名
func VerifySignature(trade string, signature []byte) bool {
	sha1 := sha1.New()
	sha1.Write([]byte(trade))
	digest := sha1.Sum([]byte(""))

	publicKey, _ := ReadFile("./data/rsa_public_key.pem")
	block, _ := pem.Decode(publicKey)
	pubInterface, _ := x509.ParsePKIXPublicKey(block.Bytes)
	pub := pubInterface.(*rsa.PublicKey)

	// 用公钥验证签名
	return rsa.VerifyPKCS1v15(pub, crypto.Hash(0), digest, signature) == nil
}

func main() {
	trade := "这是明文"
	signature := DigitalSignature(trade)
	fmt.Println("验证数字签名：", VerifySignature(trade, signature))
}

/*
生成1024位的RSA私钥：
openssl genrsa -out data/rsa_private_key.pem 1024
根据私钥生成公钥：
openssl rsa -in data/rsa_private_key.pem -pubout -out data/rsa_public_key.pem

pem 是一种标准的格式，它通常包含页眉和页脚

*/