package middleware

//AUTHOR:HILMAN
import (
	"app/configs"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

//
func MustAuth(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(h1 http.ResponseWriter, h2 *http.Request) {
		log.Println("XXXX")
		authHeader, err := checkAuthBearer(h1, h2)
		if err != nil {
			return
		}
		defer fmt.Fprintf(h1, "Invalid Key")
		myCipher, err := aes.NewCipher(configs.GetConfig().JwtKey)

		if err != nil {
			panic("JWT INVALID:CIPHER " + err.Error())

			return
		}
		myGCm, err := cipher.NewGCM(myCipher)
		if err != nil {
			panic("JWT INVALID:GCM")
			return
		} else {
			log.Println("GCM OK")
		}
		nonce := make([]byte, myGCm.NonceSize())
		if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
			log.Println("PANIC READFUL")
			panic(err.Error())
		}
		log.Println("OK READFUL")
		cipherText := myGCm.Seal(nonce, nonce, []byte("asdadasd"), nil)

		fmt.Println(cipherText)
		log.Println("CIPHERTEXT OK")
	}
}

//FlagNoKey asdasd
const FlagNoKey, FlagCipherErr, FlagGcmErr, FlagRanderr = "nokey", "cipherErr", "gcmerr", "randerr"

func checkAuthBearer(h1 http.ResponseWriter, h2 *http.Request) (*string, error) {
	authHeader := h2.Header["Authorization"]
	if authHeader == nil {
		return nil, errors.New(FlagNoKey)
	}
	if len(authHeader) == 0 {
		return nil, errors.New(FlagNoKey)
	}
	if strings.Index(authHeader[0], "Bearer") < 0 {
		return nil, errors.New(FlagNoKey)
	}
	return &authHeader[0], nil
}

func generateAesKey(key *string, data string) ([]byte, error) {
	myCipher, err := aes.NewCipher([]byte(*key))

	if err != nil {
		return nil, errors.New(FlagCipherErr)
	}
	myGCm, err := cipher.NewGCM(myCipher)
	if err != nil {
		return nil, errors.New(FlagGcmErr)
	}
	nonce := make([]byte, myGCm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, errors.New(FlagRanderr)
	}
	cipherText := myGCm.Seal(nonce, nonce, []byte(data), nil)
	return cipherText, nil
}
