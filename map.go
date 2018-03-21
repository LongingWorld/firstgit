package main

import (
	"fmt"
)

func main() {
	// 声明一个key是字符串，值为int的字典,这种方式的声明需要在使用之前使用make初始化
	var map_test map[string]int
	map_test = make(map[string]int)
	map_test["Heng"] = 2
	map_test["Jin"] = 5
	map_test["Yu"] = 0

	fmt.Println("Heng Jin Yu is:", map_test["Heng"], map_test["Jin"], map_test["Yu"], "map_test len is:", len(map_test))
	map_test["Heng"] = 5
	map_test["Jin"] = 2
	fmt.Println("Heng Jin Yu is:", map_test["Heng"], map_test["Jin"], map_test["Yu"], "map_test len is:", len(map_test))
	map_value, ok := map_test["Zhou"] //map内置有判断是否存在key的方式
	if ok {
		fmt.Println("Zhou is in the map and it's value is ", map_value)
	} else {
		fmt.Println("We have no value associated with Zhou in the map")
	}
	//for配合range可以用于读取slice和map的数据
	for k, v := range map_test {
		fmt.Println("map_test's key :", k)
		fmt.Println("map_test's value :", v)
	}
	//由于 Go 支持 “多值返回”, 而对于“声明而未被调用”的变量, 编译器会报错, 在这种情况下, 可以使用_来丢弃不需要的返回值
	for _, v1 := range map_test {
		fmt.Println("map_test's value :", v1)
	}
	var sum int
	sum = 121
	for sum < 1000 { //省略;  相当于while循环
		sum += sum
		fmt.Println(sum)
	}
	/*for ; sum < 1000; {
	sum += sum
	}*/
	fmt.Println(sum)

	for i := 0; i < 5; i++ {
		defer fmt.Println("defer first in last out:", i)
	}
}
