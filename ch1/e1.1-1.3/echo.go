package main

import(
	"fmt"
	"os"
)

func main(){

	for idx, str := range os.Args{
		fmt.Println(idx,str)
	}

}