package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("Мы загадали число от 1 до 100.")
	fmt.Println("Попробуйте угадать")
	//fmt.Println(target)

	reader := bufio.NewReader(os.Stdin)

	success := false
	for x := 0; x < 10; x++ {
		fmt.Println("У вас осталось", 10-x, "попыток")

		fmt.Print("Ввведите свое число: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		if guess < target {
			fmt.Println("Неудача, ваше число МЕНЬШЕ")
		} else if guess > target {
			fmt.Println("Неудача. Ваше число БОЛЬШЕ")
		} else {
			success = true
			fmt.Println("Круто! Вы угадали")
			break
		}
	}
	if !success {
		fmt.Println("К сожалению, вы не угадали. Число было ", target)
	}

}
