package main

import (
	"fmt"
	"net/http"
	"log"
	"os"
)

func main()  {
	url := "http://www.baidu.com"
	resp,err := http.Get(url)

	if err!=nil{
		fmt.Println(err.Error())
		log.Fatal(err)
	}

	getBody(resp)

}


func getBody(resp *http.Response)  {

	if resp.StatusCode == http.StatusOK{
		fmt.Println(resp.StatusCode)
	}

	defer resp.Body.Close()

	buf := make([]byte,1024)


	f,err1 := os.OpenFile("html.txt",os.O_RDWR|os.O_APPEND|os.O_CREATE,os.ModePerm)

	if err1!=nil{
		panic(err1)
		return
	}

	for {
		n,_ :=resp.Body.Read(buf)
		if 0==n{
			break
		}
		f.WriteString(string(buf[:n]))
	}

	defer f.Close()

}