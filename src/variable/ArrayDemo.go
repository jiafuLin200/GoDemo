package main

import "fmt"

func modify(array [5]int) {
	array[0] = 10 // 试图修改数组的第一个元素
	fmt.Println("In modify(), array values:", array)
}

func testModify() {
	array := [5]int{1, 2, 3, 4, 5} // 定义并初始化一个数组
	modify(array)                  // 传递给一个函数,并试图在函数体内修改这个数组内容
	fmt.Println("In testModify(), array values:", array)
}

func testSlice01() {
	// 先定义一个数组
	var myArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// 基于数组创建一个数组切片
	var mySlice []int = myArray[:5]
	fmt.Println("Elements of myArray: ")
	for _, v := range myArray {
		fmt.Print(v, " ")
	}
	fmt.Println("\nElements of mySlice: ")
	for _, v := range mySlice {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func testSlice02() {
	mySlice := make([]int, 5, 10)
	fmt.Println("len(mySlice):", len(mySlice))
	fmt.Println("cap(mySlice):", cap(mySlice))
}

func main() {
	testModify()
	testSlice01()
	testSlice02()
}
