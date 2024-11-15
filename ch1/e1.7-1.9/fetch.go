package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main(){
	for _,url := range os.Args[1:]{
		if !strings.HasPrefix(url,"http://") {
			url = "http://" + url
		}
		resp,err  := http.Get(url)
		if err != nil{
			fmt.Fprintf(os.Stderr,"fetch : %v",err)
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil{
			fmt.Fprintf(os.Stderr,"fetch : %v",err)
		}
		resp.Body.Close()
		fmt.Println("Status:", resp.Status)
	}
}