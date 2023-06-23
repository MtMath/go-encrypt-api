package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func encrypt(plainText []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plainText))
	iv := ciphertext[:aes.BlockSize]

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plainText)

	return ciphertext, nil
}

func decrypt(ciphertext []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < aes.BlockSize {
		return nil, fmt.Errorf("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	stream.XORKeyStream(ciphertext, ciphertext)

	return ciphertext, nil
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/decrypt/{key}/{cipherText}", DecryptHandler).Methods("GET")
	router.HandleFunc("/encrypt/{key}/{plainText}", EncryptHandler).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	fmt.Println("Lintening http://localhost:", port)
	http.ListenAndServe(":"+port, router)
}

func DecryptHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cipherText := vars["cipherText"]
	key := vars["key"]
	ciphertextBytes, _ := hex.DecodeString(cipherText)
	result, err := decrypt(ciphertextBytes, []byte(key))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("Decrypt: " + cipherText + "\nIn:" + string(result))
	fmt.Fprintln(w, string(result))
}

func EncryptHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	plainText := vars["plainText"]
	key := vars["key"]
	cipherText, err := encrypt([]byte(plainText), []byte(key))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("Encrypt: " + plainText + "\nIn:" + hex.EncodeToString(cipherText))
	fmt.Fprintln(w, hex.EncodeToString(cipherText))
}
