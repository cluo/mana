package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"os"
	"strings"
)

func Md5String(s string) string {
	var hh = md5.New()
	hh.Write([]byte(s))

	result := fmt.Sprintf("%x", hh.Sum(nil))
	return result
}

func Md5File(s string) string {
	const SIZE = 1024
	var buf [SIZE]byte
	var hh = md5.New()
	file, err := os.Open(s)
	if file == nil {
		fmt.Fprintln(os.Stderr, "open file: ", err)
		os.Exit(1)
	}
	for {
		fn, _ := file.Read(buf[:])
		if fn <= 0 {
			break
		}
		hh.Write(buf[0:fn])
	}
	file.Close()
	result := fmt.Sprintf("%x", hh.Sum(nil))
	return result
}

func short(s string) string {
	var r string = s[8:24]
	return r
}

func upper(s string) string {
	return strings.ToUpper(s)
}

var isUpper = flag.Bool("u", false, "print result in uppercase.")
var isShort = flag.Bool("s", false, "print result of 16-bit.")
var isFile = flag.Bool("f", false, "check the file.")

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		flag.PrintDefaults()
		os.Exit(1)
	}
	var s string
	if *isFile {
		s = Md5File(flag.Arg(0))
	} else {
		s = Md5String(flag.Arg(0))
	}
	if *isUpper {
		var Us = upper(s)
		if *isShort {
			fmt.Println(short(Us))
		} else {
			fmt.Println(Us)
		}
	} else {
		if *isShort {
			fmt.Println(short(s))
		} else {
			fmt.Println(s)
		}
	}
}
