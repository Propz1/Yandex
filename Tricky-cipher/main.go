package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"sync"
)

var (
	num_ciph string
	FNP      int = 2
	date     int = 4

	factor1 int
	factor2 int
)

type People struct {
	mu  sync.Mutex
	cip map[int]string
}

func main() {

	var ciphers string

	store := New()

	input, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer input.Close()

	r := bufio.NewReader(input)
	sc := bufio.NewScanner(r)

	var n int = 0

	for sc.Scan() {

		if n > 0 {
			store.set(n, sc.Text())
		} else {
			num_ciph = sc.Text()
		}
		n++
	}

	if err := sc.Err(); err != nil {
		log.Fatalf("scan file error: %v", err)
		return
	}

	n_ciph, e := strconv.Atoi(num_ciph)

	if e != nil {
		fmt.Printf("Error: %s", err.Error())
		os.Exit(1)
	}

	body := make([]string, n_ciph)

	for i := 1; i <= n_ciph; i++ {
		body = append(body, store.getCipher(i))
	}

	ciphers = strings.Join(body, " ")

	err = os.WriteFile("output.txt", []byte(ciphers), 0666)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func New() *People {
	return &People{
		cip: make(map[int]string),
	}
}

func (p *People) set(key int, value string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.cip[key] = value

}

func (p *People) get(key int) string {
	p.mu.Lock()
	defer p.mu.Unlock()
	item := p.cip[key]

	return item
}

func alphaNum(letter rune) int {
	char := int(letter)
	char -= 64
	return char
}

func (p *People) getCipher(key int) string {
	var l int
	var num int
	var cipher int

	st := p.get(key)
	ss := strings.Split(st, ",")

	for k, value := range ss {

		if k == 0 {
			factor2 = alphaNum(rune(value[0]))
		}

		for k <= FNP {
			l += len(value)
		}

		for k > FNP && k <= date {
			for _, letter := range value {

				n, err := strconv.Atoi(string(letter))

				if err != nil {
					fmt.Printf("Error: %s", err.Error())
					os.Exit(1)
				}

				num += n
			}
		}
	}

	factor1 = num

	cipher = l + factor1*64 + factor2*256

	x := int64(cipher)
	cipher_hex := strconv.FormatInt(x, 16)

	return strings.ToUpper(cipher_hex[len(cipher_hex)-3:])
}
