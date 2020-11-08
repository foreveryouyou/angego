package libs

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

type MyAes struct {
	Key       []byte
	block     cipher.Block
	blockSize int
}

func NewMyAes(key string) (p *MyAes, err error) {
	p = new(MyAes)
	p.Key = []byte(key)
	// 16,24,32
	lKey := len(p.Key)
	if lKey < 16 {
		err = errors.New("key至少16位")
		return
	}
	if lKey >= 32 {
		p.Key = p.Key[:32]
	} else if lKey >= 24 {
		p.Key = p.Key[:24]
	} else {
		p.Key = p.Key[:16]
	}
	block, err := aes.NewCipher(p.Key)
	if err != nil {
		return
	}
	p.block = block
	p.blockSize = block.BlockSize()
	return
}

func (p *MyAes) Encrypt(src []byte) (crypted []byte) {
	src = p.pKCS7Padding(src, p.blockSize)
	blockMode := cipher.NewCBCEncrypter(p.block, p.Key[:p.blockSize])
	crypted = make([]byte, len(src))
	blockMode.CryptBlocks(crypted, src)
	return
}

func (p *MyAes) Decrypt(crypted []byte) (src []byte, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.New(e.(string))
			return
		}
	}()
	blockMode := cipher.NewCBCDecrypter(p.block, p.Key[:p.blockSize])
	src = make([]byte, len(crypted))
	blockMode.CryptBlocks(src, crypted)
	src, err = p.pKCS7UnPadding(src)
	return
}

func (p *MyAes) EncryptStr(src string) (crypted string) {
	crypted = base64.RawURLEncoding.EncodeToString(p.Encrypt([]byte(src)))
	return
}

func (p *MyAes) DecryptStr(crypted string) (src string, err error) {
	_crypted, err := base64.RawURLEncoding.DecodeString(crypted)
	if err != nil {
		return
	}
	_src, err := p.Decrypt(_crypted)
	if err != nil {
		return
	}
	src = string(_src)
	return
}

func (p *MyAes) pKCS7Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}

func (p *MyAes) pKCS7UnPadding(origData []byte) (dst []byte, err error) {
	length := len(origData)
	if length <= 0 {
		err = errors.New("下标越界1")
		return
	}
	unpadding := int(origData[length-1])
	if length <= unpadding {
		err = errors.New("下标越界")
		return
	}
	dst = origData[:(length - unpadding)]
	return
}
