package main

import (
	"bufio"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func generateIV(bytes int) []byte {
	b := make([]byte, bytes)
	rand.Read(b)
	return b
}

func encrypt(block cipher.Block, value []byte, iv []byte) []byte {
	stream := cipher.NewCTR(block, iv)
	ciphertext := make([]byte, len(value))
	stream.XORKeyStream(ciphertext, value)
	return ciphertext
}

func decrypt(block cipher.Block, ciphertext []byte, iv []byte) []byte {
	stream := cipher.NewCTR(block, iv)
	plain := make([]byte, len(ciphertext))
	// XORKeyStream is used to decrypt too!
	stream.XORKeyStream(plain, ciphertext)
	return plain
}

func main1() {
	block, err := aes.NewCipher([]byte("1234567890123456"))
	if err != nil {
		panic(err)
	}

	iv := generateIV(block.BlockSize())

	if bs, err := ioutil.ReadFile("./qiyu.jpeg"); err != nil {
		panic(err)
	} else {
		fmt.Println("====== file len====: ", len(bs))
	}

	// read a file
	fi, err := os.Open("./qiyu.jpeg")
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	bufioReader := bufio.NewReader(fi)
	var encryptedBytes []byte

	fmt.Println("============================ Begin to encrypt")
	for {
		tmpBytes := make([]byte, block.BlockSize())
		if _, err := bufioReader.Read(tmpBytes); err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		} else {
			ciphertext := encrypt(block, tmpBytes, iv)
			encryptedBytes = append(encryptedBytes, ciphertext...)
		}
	}
	fmt.Println("======", len(encryptedBytes))

	newFile, err := os.Create("./qiyu1.jpeg")

	fmt.Println("========== test encrypted file")
	if _, err := newFile.Write(encryptedBytes); err != nil {
		panic(err)
	}

	fmt.Println("============================== Begin to decrypt")

	bufReader := bytes.NewReader(encryptedBytes)
	newFile, err = os.Create("./qiyu2.jpeg")
	// for {
	// 	tmpBytes := make([]byte, block.BlockSize())
	// 	if _, err := bufReader.Read(tmpBytes); err != nil {
	// 		if err == io.EOF {
	// 			break
	// 		} else {
	// 			panic(err)
	// 		}
	// 	} else {
	// 		plaintext := decrypt(block, tmpBytes, iv)
	// 		if _, err := newFile.Write(plaintext); err != nil {
	// 			panic(err)
	// 		}
	// 	}
	// }

	aescrt(bufReader, newFile, block, iv)
}

var key = []uint8{0x30, 0x82, 0x04, 0xa2, 0x02, 0x01, 0x00, 0x02, 0x82, 0x01, 0x01, 0x00, 0xbe, 0xea, 0xce, 0x8b}
var iv = []uint8{0xc1, 0x2a, 0x9a, 0xcc, 0x16, 0x78, 0xa5, 0x7f, 0x1c, 0x39, 0x22, 0x8c, 0x17, 0x0b, 0x4f, 0xc8}

func main() {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	test(block, iv)
	test1(block, iv)
	// createIV()
}

func createIV() {
	const key = "48656c6c6f20476f"
	iv := generateIV(16)
	var (
		block       cipher.Block
		encryptedIV []byte
		err         error
	)

	if block, err = aes.NewCipher([]byte(key)); err != nil {
		panic(err)
	}

	encryptedIV = aesctrHelp(iv, block, []byte(iv))
	fmt.Println("=======", base64.URLEncoding.EncodeToString(encryptedIV))
}

func aesctrHelp(value []byte, block cipher.Block, iv []byte) []byte {
	stream := cipher.NewCTR(block, iv)
	ciphertext := make([]byte, len(value))
	stream.XORKeyStream(ciphertext, value)
	return ciphertext
}

func test(block cipher.Block, iv []byte) {
	fmt.Println("======== test")
	fi, err := os.Open("./qiyu.jpeg")
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	newFile, err := os.Create("./qiyu1.jpeg")
	defer newFile.Close()

	aescrt(fi, newFile, block, iv)

	// newFile1, err := os.Create("./qiyu2.jpeg")

	// aescrt(newFile, newFile1, block, iv)

}

func test1(block cipher.Block, iv []byte) {
	fmt.Println("======== test1")
	fi, err := os.Open("./qiyu1.jpeg")
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	newFile, err := os.Create("./qiyu2.jpeg")
	defer newFile.Close()

	aescrt(fi, newFile, block, iv)
}

func aescrt(rc io.Reader, wc io.Writer, block cipher.Block, iv []byte) {

	for {
		tmpBytes := make([]byte, block.BlockSize())
		fmt.Println("===============")
		if _, err := rc.Read(tmpBytes); err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		} else {
			plaintext := decrypt(block, tmpBytes, iv)
			if _, err := wc.Write(plaintext); err != nil {
				panic(err)
			}
		}
	}
}
