package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

// var privateKey = []byte(`
// -----BEGIN RSA PRIVATE KEY-----
// MIIEowIBAAKCAQEAwOJK1RJBUwRu/5aCyktTaietXFMOAAkElhSq1M6BocVWs7yD
// y592CX30Bl0Ul4faWM9EZSlhak8Ay1CdMNis+lBZanKmAO2bPmSIIYBDQE2BzLIo
// MoJWi/Cd5PevioKSRPytqVB/S4+xz1IOD8Y407SZM3LfZ5XMfqC+VHpcnAycQ8iT
// FK0s3yjImathFNF3U7fiEzU4G7PJRn8e9ntubDd1pXYABqrVF/REcd/3Rs/qrlhG
// v3b7tAXZb2lkSLdCq3Md+BMksxUCoH3rZijSphbZSCdIrzofg+IG0y5WtdsBz6uw
// Ol2QX/EUoEdO+xhLgaOFykUoWz037ZzkLEhKkQIDAQABAoIBAB+1lAPPSnnxYqYW
// Ak5rb70l5LQm20haMyzRHPx7Loh/vq8xsKELCAardDCPoNEAfn7XJDFVSjSF5GWI
// TS84j8de6jQ7wNqqNTleoZqQUX4Cv/H83+rdzoiW9/4qUet9Z7p7p7kMCMFNUDf7
// D2C8f58eM4lnux52W/X9SwzsSMlGaGHcAKPz4vXUFWyt3naVtANhdkHjgKxA0Ev4
// W7yRgpbOKruPKzBNTRXAqq+yHZj/pONtXl8do+plwhHU8CW0BPyvkU4DH52lxWza
// mM71ow8UJC30FXF/NZ+wthFnRZO3/dhaeuNYgX7yAs3DhNn7Q8nzU4ujd8ug2OGf
// iJ9C8YECgYEA32KthV7VTQRq3VuXMoVrYjjGf4+z6BVNpTsJAa4kF+vtTXTLgb4i
// jpUrq6zPWZkQ/nR7+CuEQRUKbky4SSHTnrQ4yIWZTCPDAveXbLwzvNA1xD4w4nOc
// JgG/WYiDtAf05TwC8p/BslX20Ox8ZAXUq6pkAeb1t8M2s7uDpZNtBMkCgYEA3QuU
// vrrqYuD9bQGl20qJI6svH875uDLYFcUEu/vA/7gDycVRChrmVe4wU5HFErLNFkHi
// uifiHo75mgBzwYKyiLgO5ik8E5BJCgEyA9SfEgRHjozIpnHfGbTtpfh4MQf2hFsy
// DJbeeRFzQs4EW2gS964FK53zsEtnr7bphtvfY4kCgYEAgf6wr95iDnG1pp94O2Q8
// +2nCydTcgwBysObL9Phb9LfM3rhK/XOiNItGYJ8uAxv6MbmjsuXQDvepnEp1K8nN
// lpuWN8rXTOG6yG1A53wWN5iK0WrHk+BnTA7URcwVqJzAvO3RYVPqqlcwTKByOtrR
// yhxcGmdHMusdWDaVA7PpS1ECgYATCGs/XQLPjsXje+/XCPzz+Epvd7fi12XpwfQd
// Z5j/q82PsxC+SQCqR38bwwJwELs9/mBSXRrIPNFbJEzTTbinswl9YfGNUbAoT2AK
// GmWz/HBY4uBoDIgEQ6Lu1o0q05+zV9LgaKExVYJSL0EKydRQRUimr8wK0wNTivFi
// rk322QKBgHD3aEN39rlUesTPX8OAbPD77PcKxoATwpPVrlH8YV7TxRQjs5yxLrxL
// S21UkPRxuDS5CMXZ+7gA3HqEQTXanNKJuQlsCIWsvipLn03DK40nYj54OjEKYo/F
// UgBgrck6Zhxbps5leuf9dhiBrFUPjC/gcfyHd/PYxoypHuQ3JUsJ
// -----END RSA PRIVATE KEY-----
// `)

type config struct {
	PortalPrivateKey string `yaml:"portal_private_key"`
}

var allConfigs struct {
	Release config
	Debug   config
	Test    config
}

func loadConfig(filepath, currentEnv string) (*config, error) {
	buf, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("Failed to open async task config file: %s", err)
	}

	err = yaml.Unmarshal(buf, &allConfigs)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse async task config file: %s", err)
	}
	switch currentEnv {
	case "release":
		return &allConfigs.Release, nil
	case "debug":
		return &allConfigs.Debug, nil
	case "test":
		return &allConfigs.Test, nil
	default:
		return nil, fmt.Errorf("Unsupported environment: %s", currentEnv)
	}
}

const originPassowrd = "LoveMeLoveMyDog"

var priv *rsa.PrivateKey

func DecryptByPrivateKey(ciphertext []byte) ([]byte, error) {
	if len(ciphertext) == 0 {
		return []byte{}, nil
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

func EncryptedByPublicKey(cleartext string) ([]byte, error) {
	if cleartext == "" {
		return []byte{}, nil
	}

	return rsa.EncryptPKCS1v15(rand.Reader, &priv.PublicKey, []byte(cleartext))
}

func main() {
	var (
		err error
		pp  *config
	)

	if pp, err = loadConfig("./private_key.yml", "test"); err != nil {
		log.Println("load config failed", err)
	}
	log.Println("pptest\n", pp.PortalPrivateKey)

	block, _ := pem.Decode([]byte(pp.PortalPrivateKey))
	if block == nil {
		log.Println("decode failed")
		return
	}
	priv, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		log.Println("parse failed")
		return
	}
	log.Println("privtest", priv)
	ePassowrd, err := EncryptedByPublicKey(originPassowrd)
	if err != nil {
		log.Println("encrypted Failed", err)
		return
	}
	tt := string(ePassowrd)
	fmt.Println("eetest", []byte(tt))

	pwd, err := DecryptByPrivateKey(ePassowrd)
	if err != nil {
		log.Println("decrypted Failed", err)
		return
	}

	fmt.Println("pwd", string(pwd))
}
