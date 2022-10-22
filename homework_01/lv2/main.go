package main

import (
	"fmt"
)

func trans(s string)int{
	var sum int= 1
	for _,i:=range s{
		sum=sum*int(i-rune('A')+1)
		//fmt.Println(int(i-rune('A')+1))
	}
	return sum%47
}

func main(){
	var name1 string
	var name2 string
	fmt.Scanln(&name1)
	fmt.Scanln(&name2)
	if trans(name1)==trans(name2) {
		fmt.Println("Go")
	} else {
		fmt.Println("Stay")
	}
}
