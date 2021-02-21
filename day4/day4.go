package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"strings"

	flag "github.com/spf13/pflag"
)

var zeros = flag.IntP("zeros", "z", 5, "Number of leading zeros.  Default 5")

func encrypt(s string, d int) (string, error) {
	m := md5.New()
	_, err := io.WriteString(m, s)
	if err != nil {
		return "", err
	}
	fmt.Fprintf(m, "%d", d)
	return fmt.Sprintf("%x", m.Sum(nil)), nil
}

func main() {
	flag.Parse()
	secretKey := "iwrupvqb"
	if flag.NArg() > 0 {
		secretKey = flag.Args()[0]
	}

	fmt.Printf("Secret key %q\n", secretKey)
	prefix := strings.Repeat("0", *zeros)
	var e string
	var err error
	var i int
	for {
		e, err = encrypt(secretKey, i)
		if err != nil {
			fmt.Printf("%g\n", err)
			os.Exit(1)
		}

		if strings.HasPrefix(e, prefix) {
			break
		}
		i++
	}

	fmt.Printf("%s %d\n", e, i)
}
