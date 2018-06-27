package testpackage

import (
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	resu/*
描述 :  golang  AES/ECB/PKCS5  加密解密
date : 2016-04-08
author: herohenu
*/

	package main

	import (
		"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"strings"
	)

	func main() {
		/*
		*src 要加密的字符串
		*key 用来加密的密钥 密钥长度可以是128bit、192bit、256bit中的任意一个
		*16位key对应128bit
		 */
		src := "0.56"
		key := "0123456789abcdef"

		crypted := AesEncrypt(src, key)
		AesDecrypt(crypted, []byte(key))
		Base64URLDecode("39W7dWTd_SBOCM8UbnG6qA")
	}

	func Base64URLDecode(data string) ([]byte, error) {
	var missing = (4 - len(data)%4) % 4
	data += strings.Repeat("=", missing)
	res, err := base64.URLEncoding.DecodeString(data)
	fmt.Println("  decodebase64urlsafe is :", string(res), err)
	return base64.URLEncoding.DecodeString(data)
	}

	func Base64UrlSafeEncode(source []byte) string {
		// Base64 Url Safe is the same as Base64 but does not contain '/' and '+' (replaced by '_' and '-') and trailing '=' are removed.
		bytearr := base64.StdEncoding.EncodeToString(source)
		safeurl := strings.Replace(string(bytearr), "/", "_", -1)
		safeurl = strings.Replace(safeurl, "+", "-", -1)
		safeurl = strings.Replace(safeurl, "=", "", -1)
		return safeurl
	}

	func AesDecrypt(crypted, key []byte) []byte {
		block, err := aes.NewCipher(key)
		if err != nil {
		fmt.Println("err is:", err)
	}
		blockMode := NewECBDecrypter(block)
		origData := make([]byte, len(crypted))
		blockMode.CryptBlocks(origData, crypted)
		origData = PKCS5UnPadding(origData)
		fmt.Println("source is :", origData, string(origData))
		return origData
	}

	func AesEncrypt(src, key string) []byte {
		block, err := aes.NewCipher([]byte(key))
		if err != nil {
		fmt.Println("key error1", err)
	}
		if src == "" {
		fmt.Println("plain content empty")
	}
		ecb := NewECBEncrypter(block)
		content := []byte(src)
		content = PKCS5Padding(content, block.BlockSize())
		crypted := make([]byte, len(content))
		ecb.CryptBlocks(crypted, content)
		// 普通base64编码加密 区别于urlsafe base64
		fmt.Println("base64 result:", base64.StdEncoding.EncodeToString(crypted))

		fmt.Println("base64UrlSafe result:", Base64UrlSafeEncode(crypted))
		return crypted
	}

	func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
		padding := blockSize - len(ciphertext)%blockSize
		padtext := bytes.Repeat([]byte{byte(padding)}, padding)
		return append(ciphertext, padtext...)
	}

	func PKCS5UnPadding(origData []byte) []byte {
		length := len(origData)
		// 去掉最后一个字节 unpadding 次
		unpadding := int(origData[length-1])
		return origData[:(length - unpadding)]
	}

	type ecb struct {
		b         cipher.Block
		blockSize int
	}

	func newECB(b cipher.Block) *ecb {
		return &ecb{
		b:         b,
		blockSize: b.BlockSize(),
	}
	}

	type ecbEncrypter ecb

	// NewECBEncrypter returns a BlockMode which encrypts in electronic code book
	// mode, using the given Block.
	func NewECBEncrypter(b cipher.Block) cipher.BlockMode {
		return (*ecbEncrypter)(newECB(b))
	}
	func (x *ecbEncrypter) BlockSize() int { return x.blockSize }
	func (x *ecbEncrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
	panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
	panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
	x.b.Encrypt(dst, src[:x.blockSize])
	src = src[x.blockSize:]
	dst = dst[x.blockSize:]
	}
	}

	type ecbDecrypter ecb

	// NewECBDecrypter returns a BlockMode which decrypts in electronic code book
	// mode, using the given Block.
	func NewECBDecrypter(b cipher.Block) cipher.BlockMode {
		return (*ecbDecrypter)(newECB(b))
	}
	func (x *ecbDecrypter) BlockSize() int { return x.blockSize }
	func (x *ecbDecrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
	panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
	panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
	x.b.Decrypt(dst, src[:x.blockSize])
	src = src[x.blockSize:]
	dst = dst[x.blockSize:]
	}
	}lt := Add(5, 6)

	/*1.log和logf
	 *级别低，不影响后续测试
	*/
	if false {
		t.Log(result)
		t.Logf("结果是:%d", result)
	}

	/*2.Error和Errorf
	  *级别低，不影响后续
	*/
	if false {
		t.Error("出错了")
		t.Errorf("%s", "出错了")
	}

	/*3.Fatal和Fatlf
	 * 致命错误会终止当前测试，所以只输出了第一句话
	*/
	if false {
		t.Fatal("致命错误")
		t.Fatalf("%s", "致命错误")
	}

	/*4.Fail和FailNow
	 * Fail判定失败，但当前测试继续进行，FailNow当前测试结束，fail now不被输出
	*/
	if false {
		t.Fail()
		t.Log("fail")

		t.FailNow()
		t.Log("fail now")
	}

	/*5.设置打印查看test的log细节
	  *go test -v funs_test.go funs.go
	  *输出结果:
		=== RUN   TestAdd
		--- PASS: TestAdd (0.00s)
				funs_test.go:14: 11
				funs_test.go:15: 结果是:11
		=== RUN   TestDelete
		--- PASS: TestDelete (0.00s)
				funs_test.go:55: 5
		PASS
		ok      test_testingT/testpackage       0.047s
	*/

	/*6.设置超时时间
	 *go test -timeout 100ms
	 *panic: test timed out after 100ms
	*/
	if false {
		time.Sleep(2 * time.Second)
	}

	/*7.设置short模式
	 * go test -short
	 **PASS
       ok      test_testingT/testpackage       0.148s
	 * go test
	 **PASS
	   ok      test_testingT/testpackage       3.046s
	*/
	if false {
		if testing.Short() {
			time.Sleep(100 * time.Millisecond)
		} else {
			time.Sleep(3 * time.Second)
		}

	}
}

func TestDelete(t *testing.T) {
	result := Delete(8, 3)
	t.Log(result)
}
