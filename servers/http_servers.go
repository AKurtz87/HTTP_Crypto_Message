package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
)

var key string

// Unpad the plaintext using PKCS7 padding
func unpad(plaintext string) string {
	length := len(plaintext)
	padLength := int(plaintext[length-1])
	return plaintext[:length-padLength]
}

func decrypt(ciphertext string, key []byte) string {
	// Decode the base64 ciphertext
	ciphertextBytes, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		panic(err)
	}

	// Create the AES cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// Decrypt the ciphertext
	plaintext := make([]byte, len(ciphertextBytes))
	mode := cipher.NewCBCDecrypter(block, key[:block.BlockSize()])
	mode.CryptBlocks(plaintext, ciphertextBytes)

	// Unpad the plaintext
	return unpad(string(plaintext))
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	arg := r.URL.Query().Get("arg")
	fmt.Println("#! ", decrypt(arg, []byte(key)))
}

func main() {
	readerKey := bufio.NewReader(os.Stdin)
	fmt.Print("Enter key: ")
	key, _ = readerKey.ReadString('\n')
	//fmt.Println(key)

	//for i := 1; i<11; i++ {
		http.HandleFunc("/", handleGet)
		//link := "127.0.0.1:808" + strconv.Itoa(i)
		link := "127.0.0.1:8081"
		err := http.ListenAndServe(link, nil)  
		if err != nil {
			fmt.Println("Failed to start server:", err)
		}

	//}

}







