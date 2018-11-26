package main

import (
	"fmt"
	"path/filepath"
	"os"
	"log"
	"net/http"
	"net"
)

func main()  {

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Power by aboc  mayinhua@gmail.com")
	fmt.Println(fmt.Sprintf("文件目录位于 %s",dir))
	//dir = "/home/aboc/document/EhMALL/泰山·秋色 (2)/泰山·秋色"
	http.Handle("/",http.FileServer(http.Dir(fmt.Sprintf("%s/",dir))))


	//fmt.Println(err)
	var findPort chan int
	findPort = make(chan int ,1)
	go func(){
		for port := 8080;port <=9000;port++{
			_,err = net.Dial("tcp",fmt.Sprintf(":%d",port))
			if err != nil{
				findPort <- port
				break
			}
		}
	}()
	serverPort := <-findPort
	startServer(serverPort)
}

func startServer(port int)error{
	fmt.Println(fmt.Sprintf("服务器启动成功，请访问 http://127.0.0.1:%d/",port))
	if err := http.ListenAndServe(fmt.Sprintf(":%d",port),nil);err != nil{
		return err
	}
	return nil
}
