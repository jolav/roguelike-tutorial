//go:build js && wasm

/* */

package main

import (
	cryptoRand "crypto/rand"
	"encoding/base32"
	"io"
	"math/rand"
	"net/http"
	"strings"
)

func randomInt(min, max int, x rand.Rand) int {
	return x.Intn(max-min+1) + min
}

func generateNick(filePath string, x rand.Rand) string {
	resp, err := http.Get(filePath)
	if err != nil {
		return "PLAYER1"
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "PLAYER1"
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "PLAYER1"
	}
	return strings.Split(string(data), "\n")[randomInt(0, 88798, x)]
}

func generateToken(n int) string {
	randomBytes := make([]byte, n)
	_, err := cryptoRand.Read(randomBytes)
	if err != nil {
		return "TOKEN"
	}
	token := base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(randomBytes)
	return token
}
