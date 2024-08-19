package main

import "fmt"

type Person struct {
	Name    string
	Age     int
	Address string
}

type PersonList []*Person

// Personを生成する関数
func CreatePerson(name string, age int, address string) *Person {
	return &Person{Name: name, Age: age, Address: address}
}

// GreetメソッドをPerson構造体に定義
func (p *Person) Greet() {
	fmt.Printf("こんにちは、私は%sです。住んでいる場所は%sです。年齢は%v歳です\n", p.Name, p.Address, p.Age)
}

func main() {
	// Personインスタンスを生成
	p1 := CreatePerson("太郎", 25, "東京")
	p2 := CreatePerson("次郎", 30, "大阪")
	p3 := CreatePerson("三郎", 35, "福岡")
	p4 := CreatePerson("四郎", 40, "名古屋")

	// PersonListに追加
	users := PersonList{}
	users = append(users, p1, p2, p3, p4)

	// 各PersonについてGreetメソッドを呼び出す
	for _, person := range users {
		person.Greet()
	}
}
