package config

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	. "github.com/onsi/ginkgo"
	"io/ioutil"
	"os"
)

const keyPath = "./keys/id_rsa"

const (
	GlobalValue = "globalVal"
	Password = encryptedProperty("j2/6gYN+yDBk3x1JEQof50lghzkUPs0rV8PWINh6QEuhCdR03Y9LwEnr6bL/v6uGZ+uHJSkMU2isj6o/tM5Rx1T/X3xYWchMIZbEG6fevSdM6umvq6poG/0dCI9uv56y7TWU7ulVcHblezVNli6wHadN0yQahhlfRtsmfn6vFwLHlphTIW1fFY8El0B3X386yyHfpcFFW77c0ccwSS37a5UPwYEyVuE4GmGQMkZYI7A9ma/cLkF2K7gsOdMBbn1gzDBjSs3W6uT2+n6XUm24cNfblL8/33POPxOZshEPPoT5yTpZC2IZqIwPYi1JOODzm6aPPRhFEqDAqQkBsCkiJ9GkCZreT6ps0++vF9dcenmS5ZpadQ/0Lwz+P7fz9CdXmk0iDCm/LSwgZ1IEpI+r0OB8/2jYCJlQ8/aMjVVQGaNAD5bhiQDoc+c8Kwa4Lv398mXdymw3nynA1KBALvd/XOnoAH6NfPiyyyvLLApohYElzPB1iMvSbD7fZKqj+lOeMEEWmFt7W+J0j9lsZmPSgKYBMngufBxSN2srDhhR1FJNmX5K5S2apE9PjgJzDFEWclb/ZM3VhrByDdX2IXJcbsah1OETFBLyuMRKH3WyhrIP1CMcIehCzcxMLkUwu6uKbgHbXdh9GXg60g/tF7cOEJObqYHd9dYc1EUyWrYM3L4=")
)

const (
	PROD environment = iota
	STAGING
	DEV
	LOCAL
)

var privateKey *rsa.PrivateKey

type environment int
type encryptedProperty string

func (e environment) String() string {
	return [...]string{"PROD", "STAGING", "DEV", "LOCAL"}[e]
}

func (p encryptedProperty) Decrypt() string {
	secret, err := base64.StdEncoding.DecodeString(string(p))
	plaintext, err := rsa.DecryptOAEP(sha512.New(), rand.Reader, privateKey, secret, nil)
	if err != nil {
		return fmt.Sprintf("FAILED(%s)", err.Error())
	}
	fmt.Printf("Decrypting: %s...: %s****** \n", p[:80], plaintext[:3])
	return string(plaintext)
}

func Encrypt() {
	publicKey := privateKey.Public()
	rsaPublicKey := publicKey.(*rsa.PublicKey)
	ciphertext, _ := rsa.EncryptOAEP(sha512.New(), rand.Reader, rsaPublicKey, []byte("qwerty1"), nil)
	uEnc := base64.StdEncoding.EncodeToString(ciphertext)
	fmt.Printf("sample password encryption: %v \n", uEnc)
}

func LoadKey() {
	privateKeyFile, err:=ioutil.ReadFile(keyPath)
	if err != nil {
		fmt.Printf("key not found: %v \n", keyPath)
		os.Exit(1)
	}
	block, _ := pem.Decode(privateKeyFile)
	if block.Type != "RSA PRIVATE KEY" {
		fmt.Printf("invalid key: %v \n", block.Type)
		os.Exit(1)
	}

	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		fmt.Printf("invalid key: %v \n", err)
		os.Exit(1)
	}

	privateKey = key
}


// SkipIn takes 1 or more environment parameters if the current environment matches the supplied environments the test will not be executed.
// This method should be placed at the beginning of any test that needs to be excluded from running in specific environments
func SkipIn(environments ...environment){
	for _, e := range environments {
		if e == Env {
			Skip(fmt.Sprintf("Skipped: %v, Env: %v",CurrentGinkgoTestDescription().TestText, e.String()))
		}
	}
}

// OnlyIn takes 1 or more environment parameters if the current environment does not any of the supplied environments the test will not be executed.
// This method should be placed at the beginning of any test that can only be run in specific environments
func OnlyIn(environments ...environment){
	for _, e := range environments {
		if e == Env {
			return
		}
	}
	Skip(fmt.Sprintf("Skipped: %v, Env: %v",CurrentGinkgoTestDescription().TestText, Env))
}



