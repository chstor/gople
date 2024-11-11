package main

import (
	"fmt"
	"io"
	 "os"
	"net/http"
	"strings"
	 "time"
)
//go.dev godoc.org
func main(){
	start := time.Now()
	ch := make(chan string)
	for _,url := range os.Args[1:]{
		if !strings.HasPrefix(url,"https://") {
			url = "https://" + url
			go fetch(url,ch)
		}
	}
	for range os.Args[1:]{
		fmt.Println(<-ch)
	}
	fmt.Printf("%2.fs elpased",time.Since(start).Seconds())
}

func fetch(url string,ch chan <- string){
	start := time.Now() 
	resp,err := http.Get(url)
	if err != nil{
		ch <- fmt.Sprint(err)
		return
	}
	nbytes,err  := io.Copy(io.Discard, resp.Body)
	if err != nil{
		ch <- fmt.Sprint(err)
		return 
	}
	resp.Body.Close()
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2f %7d %s",secs,nbytes,url)
}