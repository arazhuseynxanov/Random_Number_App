package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	target := numRand(1, 100)
	fmt.Print("1 ile 100 Arasi Ededi Tapin")
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < 10; i++ {
		fmt.Println(10-i, "Qeder Sansiniz Qaldi!")
		fmt.Print("Sansinizi Sinayin:")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		input = strings.TrimSpace(input)
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println(err)
		}
		if num > target {
			fmt.Println("Texmin Elediniyiz Reqem Boyukdur....Bundan Kicik Reqemi Daxil Edin!!")
		} else if num < target {
			fmt.Println("Texmin Elediniyiz Reqem Kicik....Bundan Boyuk Reqemi Daxil Edin!!")
		} else {
			fmt.Println("Cavab Dogru", i, "Seferde Tapdiniz!")
			break
		}
	}
}

func numRand(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
