// Copyright muflax <mail@muflax.com>, 2015
// License: GNU GPLv3 (or later) <http://www.gnu.org/copyleft/gpl.html>

package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func charPassword(length int) {
	pool := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	password := make([]byte, length)

	for i := 0; i < length; i++ {
		password[i] = pool[rand.Intn(len(pool))]
	}

	fmt.Println(string(password))
}

func wordPassword(bits int) {
}

func main() {
	var (
		short = flag.Bool("s", false, "short password?")
		long  = flag.Bool("l", false, "long password?")
	)
	flag.Parse()

	rand.Seed(time.Now().UnixNano())

	if *short {
		charPassword(12)
	} else if *long {
		charPassword(32)
	} else {
		wordPassword(70)
	}
}
