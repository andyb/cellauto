package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	log.Print("Started")
	ca := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	log.Printf("%v", ca)
	for i := 0; i < 15; i++ {
		ngca := nextGen(applyRules, ca)
		log.Printf("%v", ngca)
		ca = ngca
	}

}

func nextGen(fn func(int, int, int) int, origin []int) []int {
	nextGen := make([]int, len(origin))

	for i := 1; i < len(origin)-1; i++ {
		nextGen[i] = fn(origin[i-1], origin[i], origin[i+1])
	}

	return nextGen
}

func applyRules(left int, curr int, right int) int {
	res, _ := strconv.ParseInt(fmt.Sprintf("%v%v%v", left, curr, right), 2, 2)
	return int(res)
}
