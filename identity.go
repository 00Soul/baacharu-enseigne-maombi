package oxpit

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type Token struct {
	data []byte
}

func NewToken(bytes []byte) Token {
	var token Token

	token.data = make([]byte, len(bytes))
	copy(token.data, bytes)

	return token
}

func NewTokenFromHex(str string) (Token, error) {
	var token Token

	str = strings.Replace(str, "0x", "", 1)

	bytes, err := hex.DecodeString(str)
	if err == nil {
		token = NewToken(bytes)
	}

	return token, err
}

func NewTokenFromBase64(str string) (Token, error) {
	var token Token

	bytes, err := base64.StdEncoding.DecodeString(str)
	if err == nil {
		token = NewToken(bytes)
	}

	return token, err
}

func NewTokenFromBase32(str string) (Token, error) {
	var token Token

	bytes, err := base32.StdEncoding.DecodeString(str)
	if err == nil {
		token = NewToken(bytes)
	}

	return token, err
}

func (token Token) Length() int {
	return len(token.data)
}

func (token Token) ToBytes() []byte {
	bytes := make([]byte, len(token.data))

	copy(bytes, token.data)

	return bytes
}

func (token Token) ToHex() string {
	return hex.EncodeToString(token.data)
}

func (token Token) ToBase64() string {
	return base64.StdEncoding.EncodeToString(token.data)
}

func (token Token) ToBase32() string {
	return base32.StdEncoding.EncodeToString(token.data)
}

func NewIdentityToken() IdentityToken {
	var buffer []byte

	buffer = strconv.AppendInt(buffer, time.Now().Unix(), 10)
	buffer = strconv.AppendInt(buffer, rand.Int63(), 10)

	hash := md5.Sum(buffer)

	return IdentityToken{NewToken(hash[0:16])}
}

const newIdentityErrorString = "Conversion of the string into base32 does not result in a 16 byte token"

const newAccessErrorString = "Conversion of the string into base32 does not result in a 32 byte token"

func NewIdentityTokenFromBase32(str string) (IdentityToken, error) {
	var identityToken IdentityToken

	token, err := NewTokenFromBase32(str)
	if err == nil {
		if token.Length() == 16 {
			identityToken = IdentityToken{token}
		} else {
			err = errors.New(newIdentityErrorString)
		}
	}

	return identityToken, err
}

func NewIdentityTokenFromBytes(bytes []byte) (IdentityToken, error) {
	token := IdentityToken{}
	var err error = nil

	if len(bytes) == 16 {
		token = IdentityToken{NewToken(bytes)}
	} else {
		err = errors.New(newIdentityErrorString)
	}

	return token, err
}

func NewAccessToken() AccessToken {
	var buffer []byte

	buffer = strconv.AppendInt(buffer, time.Now().Unix(), 10)
	buffer = strconv.AppendInt(buffer, rand.Int63(), 10)

	hash := sha256.Sum256(buffer)

	return AccessToken{NewToken(hash[0:32])}
}

func NewAccessTokenWithIdentity(token IdentityToken) AccessToken {
	buffer := token.ToBytes()
	buffer = strconv.AppendInt(buffer, time.Now().Unix(), 10)
	buffer = strconv.AppendInt(buffer, rand.Int63(), 10)

	hash := sha256.Sum256(buffer)

	return AccessToken{NewToken(hash[0:32])}
}

func NewAccessTokenFromBytes(bytes []byte) (AccessToken, error) {
	var token AccessToken
	var err error = nil

	if len(bytes) == 32 {
		token = AccessToken{NewToken(bytes)}
	} else {
		err = errors.New(newAccessErrorString)
	}

	return token, err
}

func NewAccessTokenFromBase32(str string) (AccessToken, error) {
	var accessToken AccessToken

	token, err := NewTokenFromBase32(str)
	if err == nil {
		if token.Length() == 32 {
			accessToken = AccessToken{token}
		} else {
			err = errors.New(newAccessErrorString)
		}
	}

	return accessToken, err
}
