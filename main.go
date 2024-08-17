/*
10個のランダムな整数を含むスライスを作成し、それらの数値の平均値、最大値、最小値を計算して表示するプログラムを書いてください。
次に、キーが文字列（例： "Even", "Odd"）、値がそのカテゴリに該当する数のリストであるマップを作成し、
スライス内の整数を偶数と奇数に分類してそのマップに格納し、結果を表示するプログラムを書いてください。
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func MakeRandom() []int {
	rand.Seed(time.Now().UnixNano())
	sl1 := []int{}

	for i := 1; i <= 10; i++ {
		sl1 = append(sl1, rand.Intn(100))
	}
	return sl1
}

func Ave(numbers []int) int {
	sum := 0
	for _, k := range numbers {
		sum += k
	}

	av := sum / len(numbers)
	return av
}

func maxnumber(numbers []int) int {
	num := numbers[0]

	for i := 0; i < len(numbers); i++ {
		if num < numbers[i] {
			num = numbers[i]
		}
	}
	return num
}

func containsValue(m map[string][]int, target int) bool {
	for _, values := range m {
		for _, value := range values {
			if value == target {
				return true
			}
		}
		break
	}
	return false
}

func MakeMap() map[string][]int {
	m1 := make(map[string][]int)
	EvenNumbers := []int{}
	OddNumbers := []int{}
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			EvenNumbers = append(EvenNumbers, i)
		} else {
			OddNumbers = append(OddNumbers, i)
		}
	}

	m1["Even"] = EvenNumbers
	m1["Odd"] = OddNumbers

	return m1

}
func Search(sl2 []int, m1 map[string][]int) map[string][]int {
	m2 := make(map[string][]int)
	EvenNumbers := []int{}
	OddNumbers := []int{}

	for _, v := range sl2 {
		if containsValue(m1, v) == true {
			EvenNumbers = append(EvenNumbers, v)
			m2["Even"] = EvenNumbers
		} else {
			OddNumbers = append(OddNumbers, v)
			m2["Odd"] = OddNumbers
		}
	}
	return m2
}

func main() {

	fmt.Println(MakeRandom())
	fmt.Println(Ave(MakeRandom()))
	fmt.Println(maxnumber(MakeRandom()))
	fmt.Println(Search(MakeRandom(), MakeMap()))
}
