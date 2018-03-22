package main

import (
	f "fmt"
)

type Person struct {
	name string
	age  int
}

type Student struct {
	Person    //匿名字段
	specifity string
}

func Older(a Person, b Person) (Person, int) {
	if a.age > b.age {
		return a, a.age - b.age
	}
	return b, b.age - a.age
}

func main() {
	var tom Person
	tom.age = 18
	tom.name = "Blue"                     // 赋值初始化
	lucy := Person{name: "Lucy", age: 16} // 两个字段都写清楚的初始化
	joe := Person{"Joe", 20}              // 按照struct定义顺序初始化值

	tl_Older, tl_diff := Older(tom, lucy)
	tj_Older, tj_diff := Older(tom, joe)
	lj_Older, lj_diff := Older(lucy, joe)

	f.Printf("Of %s and %s, %s is older by %d years.\n", tom.name, lucy.name, tl_Older.name, tl_diff)
	f.Printf("Of %s and %s, %s is older by %d years.\n", tom.name, joe.name, tj_Older.name, tj_diff)
	f.Printf("Of %s and %s, %s is older by %d years.\n", lucy.name, joe.name, lj_Older.name, lj_diff)

}
