package main

import "fmt"

func variableTest(){
	var a = 2
	var b =10
	c :=5
	fmt.Println(a,b,c)
	a,b,c = c,a,b
	fmt.Println(a,b,c)
}

func GetString() (a,b,c string){
	return "a","b","c"
}

func constantTest(){
	const(
		a=1
		b,c,d=0.2,0.3,"d"
	)

	const(
		c0 = iota
		c1
		c2
	)
	fmt.Println(a,b,c,d,c0,c1,c2)
}

func main(){
	fmt.Println("hello go")
	variableTest()
	fmt.Println(GetString())
	constantTest()
}