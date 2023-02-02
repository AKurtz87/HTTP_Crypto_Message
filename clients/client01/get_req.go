package main

import (
	"bufio"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const endpoint = "http://127.0.0.1:8081"

func encrypt(plaintext string, key []byte) string {
	// Create the AES cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// Pad the plaintext to a multiple of the block size
	plaintext = pad(plaintext)

	// Encrypt the plaintext
	ciphertext := make([]byte, len(plaintext))
	mode := cipher.NewCBCEncrypter(block, key[:block.BlockSize()])
	mode.CryptBlocks(ciphertext, []byte(plaintext))

	// Encode the ciphertext to a base64 string
	return base64.StdEncoding.EncodeToString(ciphertext)
}

// Pad the plaintext to a multiple of the block size using PKCS7 padding
func pad(plaintext string) string {
	blockSize := aes.BlockSize
	padLength := blockSize - (len(plaintext) % blockSize)
	padText := bytes.Repeat([]byte{byte(padLength)}, padLength)
	return plaintext + string(padText)
}

func makeGetRequest(str string) {
	u, err := url.Parse(endpoint)
	if err != nil {
		log.Fatalf("Failed to parse endpoint: %s\n", err)
	}
	q := u.Query()
	q.Set("arg", str)
	u.RawQuery = q.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		log.Fatalf("Failed to make GET request: %s\n", err)
	}
	defer resp.Body.Close()

	//body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %s\n", err)
	}
	//log.Printf("GET request returned status code %d\n", resp.StatusCode)
	//log.Printf("Response body: %s\n", string(body))
}

func inputText(text string) {
	fmt.Println("The text input is:", text)
}

func main() {
	
	readerKey := bufio.NewReader(os.Stdin)
	fmt.Print("Enter key: ")
	inputKey, _ := readerKey.ReadString('\n')
	fmt.Println(inputKey)
	// input key must be a string of 31 chars, use the dedicated function

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text: ")
		text, _ := reader.ReadString('\n')
		text = username + " >>>   "+ strings.TrimRight(text, "\n")
		ciphertext := encrypt(text, []byte(inputKey))
		

		makeGetRequest(ciphertext)
	}
}



