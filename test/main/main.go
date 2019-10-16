package main

import "fmt"

func main(){
	m := make(map[string]string)
	m["aaa"] = "132"
	if s,ok := m["aaa"];ok{
		fmt.Println(s,ok)
	}else {
		fmt.Println(s,ok)
	}

}
