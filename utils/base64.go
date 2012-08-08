package main

import (
	"bytes"
	"encoding/base64"
	"flag"
	"fmt"
	"os"
)

func encoding(s string) string {
	var buf bytes.Buffer
	var encoder = base64.NewEncoder(base64.StdEncoding, &buf)
	encoder.Write([]byte(s))
	encoder.Close()
	return buf.String()
}

func decoding(s string) string {
	var buf = bytes.NewBufferString(s)
	decoder := base64.NewDecoder(base64.StdEncoding, buf)
	var res bytes.Buffer
	res.ReadFrom(decoder)
	return res.String()
}

var isDeconder = flag.Bool("d", false, "decoding the string with base64.")
var help = flag.Bool("help", false, "show this info.")

func main() {
	var s string
	flag.Parse()
	if flag.NArg() != 1 {
		flag.PrintDefaults()
		os.Exit(1)
	}
	if *isDeconder {
		s = decoding(flag.Arg(0))
	} else {
		s = encoding(flag.Arg(0))
	}
	//fmt.Println(flag.Arg(0), "=>", s)
	fmt.Println(s)
}
