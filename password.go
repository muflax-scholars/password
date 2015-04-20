// Copyright muflax <mail@muflax.com>, 2015
// License: GNU GPLv3 (or later) <http://www.gnu.org/copyleft/gpl.html>

package main

import (
	"flag"
	"fmt"
	"math"
	"math/rand"
	"strings"
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

func wordPassword(bits float64, pool []string) {
	length := int(math.Ceil(bits / math.Log2(float64(len(pool)))))

	password := make([]string, length)
	for i := 0; i < length; i++ {
		password[i] = pool[rand.Intn(len(pool))]
	}

	fmt.Println(strings.Join(password, " "))
}

func main() {
	var (
		short = flag.Bool("s", false, "short password?")
		long  = flag.Bool("l", false, "long password?")
		debug = flag.Bool("d", false, "debug info?")
	)
	flag.Parse()

	rand.Seed(time.Now().UnixNano())

	if *short {
		charPassword(12)
	} else if *long {
		charPassword(32)
	} else {
		pool := englishWords

		if *debug {
			poolBits := math.Log2(float64(len(pool)))
			avgLen := 0.0

			for _, c := range pool {
				avgLen += float64(len(c))
			}
			avgLen /= float64(len(pool))
			fmt.Printf("%6.2f avg length / word\n", avgLen)
			fmt.Printf("%6.2f bits / word\n", poolBits)
			fmt.Printf("%6.2f bits / char\n", poolBits/avgLen)
		}

		wordPassword(70, pool)
	}
}
