package main

import "fmt"
import "cnms"

func main() {

	info := new(cnms.LineInfo)
	info.LINE = "line"
	info.I = 0
	info.INPUT = "ef_input"

	for _, f:= range cnms.GetFuncList(){
		f(info)
	}

	fmt.Println("completed")

}
