package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"flag"
)

// 使用md5 产生AES key
func aesKey(s string) []byte {
	hash := md5.New()
	hash.Write([]byte(s))
	return hash.Sum(nil)
}

// 补齐src长度为16的整数倍，不足的在结尾增加' '的slice，加密src并返回
// json package 填充结尾可以多余' ','\t','\n','\r'
func aesEncrypt(block cipher.Block, src []byte) []byte {
	var dst = make([]byte, 16)
	var fill = []byte("                ")
	var src_len = len(src)
	if src_len%16 != 0 {
		src = append(src, fill[src_len%16:]...)
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

var pass = flag.String("p", "", "the password to de decrypt")

var block cipher.Block

func init() {
	flag.Parse()
	if *pass != "" {
		var err error
		block, err = aes.NewCipher(aesKey(*pass))
		if err != nil {
			panic(err)
		}
	}
}
