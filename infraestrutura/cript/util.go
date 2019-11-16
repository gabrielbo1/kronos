package cript

import (
	"math/rand"
	"time"
)

const charset = "0123465789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

const charsetNumber = "0123465789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func stringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func newStringRand(length int) string {
	return stringWithCharset(length, charset)
}

func newStringNumberRand(length int) string {
	return stringWithCharset(length, charsetNumber)
}
