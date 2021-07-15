package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

var show bool

func main() {
	flag.BoolVar(&show, "show", false, "是否启动浏览器并显示")
	flag.Parse()
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("程序：aboc (http://www.phpec.org)")
	fmt.Println("如果使用存在问题，请访问 https://github.com/abocd/dirweb/releases 提交问题")
	fmt.Println(fmt.Sprintf("文件目录位于 %s", dir))
	http.Handle("/", http.FileServer(http.Dir(fmt.Sprintf("%s/", dir))))

	//fmt.Println(err)
	var findPort chan int
	findPort = make(chan int, 1)
	go func() {
		for port := 8080; port <= 9000; port++ {
			_, err = net.Dial("tcp", fmt.Sprintf(":%d", port))
			if err != nil {
				findPort <- port
				break
			}
		}
	}()
	serverPort := <-findPort
	go startServer(serverPort)
	url := fmt.Sprintf("http://127.0.0.1:%d/", serverPort)
	time.Sleep(3 * time.Second)
	if show {
		fmt.Println("准备打开浏览器...")
		runUrl(url)
	}
	<-findPort
}

func startServer(port int) error {
	fmt.Println(fmt.Sprintf("文件服务器启动成功，请访问 http://127.0.0.1:%d/", port))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		return err
	}
	return nil
}

func runUrl(url string) (err2 error, msg string) {
	cmd := exec.Command("explorer", url)

	err := cmd.Run()
	if err != nil {
		// fmt.Println("启动失败:", err)
		return err, "启动失败"
	} else {
		// fmt.Println("启动成功!")
		return nil, "启动成功"
	}
}
