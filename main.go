package main

import "fmt"

func main(){
	var s []int
	var n int
	var aux int
	var total int = 0
	fmt.Scan(&n)
	for i:=0; i < n; i++{
		fmt.Scan(&aux)
		s = append(s, aux)
		total = total + s[i]
		//fmt.Println(v)
	}
	fmt.Println(total)
}