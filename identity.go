package spi

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
	"math/rand"
	"strings"
	"time"
)

type Token struct {
	data []byte
}

func NewToken(bytes []byte) Token {
	token := new(Token)

	token.data = make([]byte, len(bytes))
	copy(token.data, bytes)

	return token
}

func NewTokenFromHex(str string) Token {
	var token *Token

	str = strings.Replace(str, "0x", "", 1)

	if bytes, ok := hex.DecodeString(str); ok {
		token = NewToken(len(bytes))
	}

	return token
}

func NewTokenFromBase64(str string) Token {
	var token *Token

	if bytes, ok := base64.StdEncoding.DecodeString(str); ok {
		token = NewToken(len(bytes))
	}

	return token
}

func NewTokenFromBase32(str string) Token {
	var token *Token

	if bytes, ok := base32.StdEncoding.DecodeString(str); ok {
		token = NewToken(len(bytes))
	}

	return token
}

func (token Token) ToHex() string {
	return hex.EncodeToString(token.data)
}

func (token Token) ToBase64() string {
	return base64.StdEncoding.Encode(token.data)
}

func (token Token) ToBase32() string {
	return base32.StdEncoding.Encode(token.data)
}

func NewIdentityToken() IdentityToken {
	var buffer []byte

	buffer = strconv.AppendInt(buffer, time.Now().Unix(), 10)
	buffer = strconv.AppendInt(buffer, rand.Int63(), 10)

	return IdentityToken{NewToken(md5.Sum(buffer))}
}

func IdentityTokenFromBytes(bytes []byte) (IdentityToken, bool) {
	token := IdentityToken{}
	ok := false

	if len(bytes) == 16 {
		token = IdentityToken{NewToken(bytes)}
	}

	return token, ok
}

func NewAccessToken() AccessToken {
	var buffer []byte

	buffer = strconv.AppendInt(buffer, time.Now().Unix(), 10)
	buffer = strconv.AppendInt(buffer, rand.Int63(), 10)

	return AccessToken{NewToken(sha256.Sum256(buffer))}
}

func AccessTokenFromBytes(bytes []byte) (AccessToken, bool) {
	token := AccessToken{}
	ok := false

	if len(bytes) == 32 {
		token = IdentityToken{NewToken(bytes)}
	}

	return token, ok
}

func AccessTokenFromBase32(str string) (AccessToken, bool) {
	token := NewTokenFromBase32(str)
	if token.Length() != 32 {
	}
}
