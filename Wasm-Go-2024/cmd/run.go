//go:build js && wasm

/* */

package main

import (
	"math/rand"
	"time"
)

type run struct {
	Nick  string
	Token string
	Seed  int64
	rnd   *rand.Rand
}

func newRun(a app) *run {
	var seed = time.Now().UnixNano()
	var random = rand.New(rand.NewSource(seed))
	r := &run{
		Nick:  generateNick(a.C.NicksFile, *random),
		Token: generateToken(a.C.TokenLen),
		Seed:  seed,
		rnd:   random,
	}
	return r
}
