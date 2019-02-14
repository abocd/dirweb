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
	fmt.Println("程序：aboc (http://www.phpec.org)")
	fmt.Println("如果使用存在问题，请访问 https://github.com/abocd/dirweb 提交问题")
	fmt.Println(fmt.Sprintf("文件目录位于 %s",dir))
	http.Handle("/",http.FileServer(http.Dir(fmt.Sprintf("%s/",dir))))


	//fmt.Println(err)
	var findPort chan int
	findPort = make(chan int ,1)
	go func(){
		var havePort = false
		for port := 8080;port <=9000;port++{
			_,err = net.Dial("tcp",fmt.Sprintf(":%d",port))
			if err != nil{
				findPort <- port
				havePort = true
				break
			}
		}
		if !havePort{
			fmt.Println("沒有找到合适的端口用来启动服务器，请关掉重新打开")
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
