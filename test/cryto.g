package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"mana/info"
	"net/http"
)

// 使用md5 产生AES key
func aesKey(s string) []byte {
	hash := md5.New()
	hash.Write([]byte(s))
	return hash.Sum(nil)
}

// 补齐src长度为16的整数倍，不足的在结尾增加空的slice，加密src并返回
func aesEncrypt(block cipher.Block, src []byte) []byte {
	var dst = make([]byte, 16)
    var fill = []byte("                ")
	var src_len = len(src)
	if src_len%16 != 0 {
		src = append(src, fill[16-src_len%16:]...)
	}
	var enc []byte
	for i := 0; i < src_len; i += 16 {
		block.Encrypt(dst, src[i:i+16])
		enc = append(enc, dst...)
	}
	return enc
}

// 从src解密并返回
func aesDecrypt(block cipher.Block, src []byte) []byte {
	var dst = make([]byte, 16)
	var src_len = len(src)
	var dec []byte
	for i := 0; i < src_len; i += 16 {
		block.Decrypt(dst, src[i:i+16])
		dec = append(dec, dst...)
	}
	return dec
}

var iscrypto = flag.String("p", "", "the password to de decrypt")

var block cipher.Block

func init() {
	flag.Parse()
	if *iscrypto != "" {
		var err error
		block, err = aes.NewCipher(aesKey(*iscrypto))
		if err != nil {
			panic(err)
		}
	}
}

func main() {
    var fill = "                "
    fmt.Printf("%#x\n", fill)
    fmt.Printf("%d\n", len(fill))
	resp, err := http.Get("http://127.0.0.1:12345/system")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var b = body
	if block != nil {
		b = aesDecrypt(block, b)
	}
	var sys info.System
	err = json.Unmarshal(b, &sys)
	if err != nil {
		fmt.Printf("%s\n", b)
		panic(err)
	}
	fmt.Printf("%s\n", sys)
    fmt.Printf("%#v\n",sys)
}
