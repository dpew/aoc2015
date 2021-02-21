package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"strings"
)

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
	secretKey := "iwrupvqb"
	if len(os.Args) > 1 {
		secretKey = os.Args[1]
	}

	fmt.Printf("Secret key %q\n", secretKey)
	var e string
	var err error
	var i int
	for {
		e, err = encrypt(secretKey, i)
		if err != nil {
			fmt.Printf("%g\n", err)
			os.Exit(1)
		}

		if strings.HasPrefix(e, "00000") {
			break
		}
		i++
	}

	fmt.Printf("%s %d\n", e, i)
}
