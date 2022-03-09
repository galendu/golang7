package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

//RSA加密，公钥加密
func RSAEncrypt(plainText []byte, fileName string ) []byte {
	//1.打开文件，读出内容
	file,err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	fileInfo ,err := file.Stat()
	if err != nil {
		panic(err)
	}
	buf := make([]byte,fileInfo.Size())
	file.Read(buf)
	file.Close()
	//2.pem解码
	block, _:= pem.Decode(buf)
	pubInterface,err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	pubKey := pubInterface.(*rsa.PublicKey)
	//3.使用公钥加密
	cipherText,err := rsa.EncryptPKCS1v15(rand.Reader,pubKey,plainText)
	if err != nil {
		panic(err)
	}
	return cipherText
}


//RSA解密，私钥解密
func RSADecrypt(cipherText []byte,fileName string) []byte {
	//1.打开文件，读出内容
	file,err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	fileInfo,err := file.Stat()
	if err != nil {
		panic(err)
	}
	buf := make([]byte,fileInfo.Size())
	file.Read(buf)
	file.Close()
	//2.pem解码
	block,_ := pem.Decode(buf)
	privateKey,err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	//3.使用私钥解密
	plainText,err := rsa.DecryptPKCS1v15(rand.Reader,privateKey,cipherText)
	if err !=nil {
		panic(err)
	}
	return plainText
}

//RSA签名 -- 私钥
func SignatureRSA(plainText []byte,fileName string) []byte {
	file,err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	//将私钥文件内容读出
	info,err := file.Stat()
	if err != nil {
		panic(err)
	}
	buf := make([]byte,info.Size())
	file.Read(buf)
	file.Close()
	block,_ := pem.Decode(buf)
	privateKey,err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	myhash := sha512.New()
	myhash.Write(plainText)
	hashText := myhash.Sum(nil)
	//使用rsa中函数对散列值签名
	sigText,err := rsa.SignPKCS1v15(rand.Reader,privateKey,crypto.SHA512,hashText)
	if err != nil {
		panic(err)
	}
	return sigText
}

//rsa签名认证
func VerifyRSA(plainText, sigTxt []byte ,pubFileName string) bool {
	file, err := os.Open(pubFileName)
	if err != nil {
		panic(err)
	}
	info ,err := file.Stat()
	if err != nil {
		panic(err)
	}
	buf := make([]byte, info.Size())
	file.Read(buf)
	file.Close()
	block,_ := pem.Decode(buf)
	pubInterface , err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	//进行类型断言 -> 得到公钥结构体
	publicKey := pubInterface.(*rsa.PublicKey)
	hashText := sha512.Sum512(plainText)
	err = rsa.VerifyPKCS1v15(publicKey,crypto.SHA512,hashText[:],sigTxt)
	if err == nil {
		return true
	}
	fmt.Printf("err______%s",err)
	return false
}

//生成公、私钥
func createKeyPEM (bits int) {
	key, err := rsa.GenerateKey(rand.Reader,bits)
	if err != nil {
		fmt.Println(err)
	}
	//生成私钥
	priviteKey := x509.MarshalPKCS1PrivateKey(key)
	block := &pem.Block{
		Type:"ASOUL PRIVITE KEY",
		Bytes: priviteKey,
	}
	//写入文件
	file, err := os.Create("private.pem")
	if err != nil {
		panic(err)
	}
	file.Close()
	pem.Encode(file,block)

	//产生公钥
	publicKey := &key.PublicKey
	//公钥从私钥中产生
	pkixPublicKey,err := x509.MarshalPKIXPublicKey(publicKey)
	block1 := &pem.Block{
		Type: "ASOUL PUBLIC KEY",
		Bytes: pkixPublicKey,
	}
	file2,err := os.Create("public.pem")
	if err != nil {
		panic(err)
	}
	file2.Close()
	encode := pem.Encode(file2,block1)
	fmt.Println(encode)

}

func main()  {
	//	1.实现数字签名
	//createKeyPEM(1024)
	//src := []byte("数字签名测试12345")
	//cipherText := RSAEncrypt(src,"public.pem")
	//fmt.Println(cipherText)
	//plainText := RSADecrypt(cipherText,"private.pem")
	//fmt.Println(plainText)
	//sigText := SignatureRSA(src,"private.pem")
	//b1 := VerifyRSA(src,sigText,"public.pem")
	//fmt.Printf("校验结果：%t\n",b1)
	//	2.用map和链表实现LRU缓存

	//	3.用map和堆 表实现超时缓存

}
