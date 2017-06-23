// Copyright 2017 ecofast. All rights reserved.
// Use of this source code is governed by a BSD-style license.

// Package borcrypto was translated from the public code of Delphi from Borland,
// which is really quite simple but very high to try hack.
// They are suitable for passwords and similar situations.
package borcrypto

// You can modify mulKey and addKey freely within 65535(MaxUInt16)
const (
	mulKey = 52845
	addKey = 11719
)

// Avoid use key less than 256 for safety
func Encrypt(plain []byte, key uint16) []byte {
	ret := make([]byte, len(plain))
	for i, c := range plain {
		b := c ^ byte(key>>8)
		ret[i] = b
		key = (uint16(b)+key)*mulKey + addKey
	}
	return ret[:]
}

func Decrypt(cipher []byte, key uint16) []byte {
	ret := make([]byte, len(cipher))
	for i, c := range cipher {
		b := c ^ byte(key>>8)
		ret[i] = b
		key = (uint16(c)+key)*mulKey + addKey
	}
	return ret[:]
}

func EncryptStr(plainText string, key uint16) string {
	bs := Encrypt([]byte(plainText), key)
	return string(bs)
}

func DecryptStr(cipherText string, key uint16) string {
	bs := Decrypt([]byte(cipherText), key)
	return string(bs)
}
