// Clock1 is a TCP server that periodically writes the time.
package main

import (
	"fmt"
)

func add(a, b int) int {
	return a + b
}

func main1() {

	var a int = 10          // 声明变量a并初始化为10
	var b float64 = 3.14    // 声明变量b并初始化为3.14
	var c string = "hello"  // 声明变量c并初始化为"hello"
	var d bool = true       // 声明变量d并初始化为true
	fmt.Println(a, b, c, d) // 输出：10 3.14 hello true
	// 2. 数组初始化
	var arr1 [5]int = [5]int{1, 2, 3, 4, 5} // 声明一个长度为5的数组arr1并初始化为{1, 2, 3, 4, 5}
	var arr2 = [5]int{1, 2, 3, 4, 5}        // 省略数组类型，编译器会自动推断类型
	fmt.Println(arr1, arr2)                 // 输出：[1 2 3 4 5] [1 2 3 4 5]
	// 3. 切片初始化
	var slice1 []int = []int{1, 2, 3, 4, 5} // 声明一个切片slice1并初始化为{1, 2, 3, 4, 5}
	var slice2 = []int{1, 2, 3, 4, 5}       // 省略切片类型，编译器会自动推断类型
	fmt.Println(slice1, slice2)             // 输出：[1 2 3 4 5] [1 2 3 4 5]
	// 4. 字典初始化
	var map1 map[string]int = map[string]int{"a": 1, "b": 2, "c": 3} // 声明一个字典map1并初始化为{"a": 1, "b": 2, "c": 3}
	var map2 = map[string]int{"a": 1, "b": 2, "c": 3}                // 省略字典类型，编译器会自动推断类型
	fmt.Println(map1, map2)                                          // 输出：map[a:1 b:2 c:3] map[a:1 b:2 c:3]
	// 5. 结构体初始化
	type Person struct { // 定义一个结构体Person
		Name    string // 结构体字段Name
		Age     int    // 结构体字段Age
		Gender  string // 结构体字段Gender
		Address string // 结构体字段Address
	}

}
