package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"flag"
	"log"
)

var iscrypto = flag.String("crypto", "", "crypto data")

var block cipher.Block

func aesKey(s string) []byte {
	hash := md5.New()
	hash.Write([]byte(s))
	return hash.Sum(nil)
}

func aesEncrypt(block cipher.Block, src []byte) []byte {
	var dst = make([]byte, 16)
	var src_len = len(src)
	if src_len%16 != 0 {
		src = append(src, make([]byte, 16-src_len%16)...)
	}
	var enc []byte
	for i := 0; i < src_len; i += 16 {
		block.Encrypt(dst, src[i:i+16])
		enc = append(enc, dst...)
	}
	return enc
}

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

func init() {
	flag.Parse()
	if *iscrypto != "" {
		var err error
		block, err = aes.NewCipher(aesKey("123"))
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatalln("crypto=\"\"")
	}
}

func main() {
	var message = `#UPDATE:12-06-01 22:50

127.0.0.1   localhost

#SmartHosts START

#Google Services START
203.208.47.1    0.docs.google.com
203.208.46.170  0-open-opensocial.googleusercontent.com
203.208.46.170  0-focus-opensocial.googleusercontent.com
203.208.47.1    1.docs.google.com
203.208.46.170  1-focus-opensocial.googleusercontent.com
203.208.46.170  1-open-opensocial.googleusercontent.com
203.208.47.1    2.docs.google.com
203.208.46.170  2-focus-opensocial.googleusercontent.com
203.208.46.170  2-open-opensocial.googleusercontent.com
203.208.47.1    3.docs.google.com`
	log.Printf("base 16 message: %d\n%x\n", len(message), message)
	log.Printf("string message: %d\n%s\n", len(message), message)
	msg_enc := aesEncrypt(block, []byte(message))
	log.Printf("base 16 msg_enc: %d\n%x\n", len(msg_enc), msg_enc)
	msg_dec := aesDecrypt(block, msg_enc)
	log.Printf("base 16 msg_dec: %d\n%x\n", len(msg_dec), msg_dec)
	log.Printf("string message: %d\n%s\n", len(msg_dec), msg_dec)
}
