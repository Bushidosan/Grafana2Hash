package main

import (
	b64 "encoding/base64"
	"encoding/hex"
	"fmt"
	"os"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println("Example: go run Grafana2Hash.go <password> <hash>")
		os.Exit(0)
	}
	password := os.Args[1]
	salt := os.Args[2]

	decoded_hash, _ := hex.DecodeString(password)
	hash64 := b64.StdEncoding.EncodeToString([]byte(decoded_hash))
	salt64 := b64.StdEncoding.EncodeToString([]byte(salt))
	fmt.Println("sha256:10000:" + salt64 + ":" + hash64 + "\n")

}
