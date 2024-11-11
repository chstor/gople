package main

import(
	"bufio"
	"fmt"
	"os"
)

type lineinfo struct{
	text string
	name string
}
func main(){
	cnt := make(map[lineinfo]int)
	
	files := os.Args[1:]
	if len(files) == 0{
		countLines(os.Stdin,cnt)
	}else{
		for _,arg := range files{
			f,err  := os.Open(arg)
			if err != nil{
				fmt.Fprintf(os.Stderr,"err : %v\n",err)
				continue
			}
			countLines(f,cnt)
			f.Close()
		}
		countLines(os.Stdin,cnt)
	}
	dup := make(map[string]bool)
	for lineinfo,v := range cnt{
		if v > 1 && !dup[lineinfo.name]{
			fmt.Println(lineinfo.name)
			dup[lineinfo.name] = true
		}
	}
}

func countLines(f *os.File,cnt map[lineinfo]int){
	input := bufio.NewScanner(f)
	for input.Scan(){
		cnt[lineinfo{input.Text(),f.Name()}] ++
	}
}