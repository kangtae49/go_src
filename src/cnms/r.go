package cnms

import "fmt"


type LineFunc func(*LineInfo) 

type LineInfo struct {
     LINE string
     I uint64
     INPUT string
}


func GetFuncList() [] LineFunc {
	l := [] LineFunc{
		Func1,
		Func2,
	}
	return l
}


func Func1(info *LineInfo){
	fmt.Println(info.LINE, info.I, info.INPUT)
}

func Func2(info *LineInfo){
	fmt.Println(info)
}

