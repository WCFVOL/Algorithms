package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

//func Timer(f func(), second time.Duration) {
//	ticker := time.NewTicker(second * time.Second)
//	for range ticker.C {
//		f()
//	}
//}

const dir = "/Users/didi/go/src/github.com/WCFVOL"

func main() {

	//strs := strings.Split("key	value", "\t")
	//fmt.Println(strs)
	//fmt.Println("start")
	////time.Sleep(5*time.Second)
	//Recovery()
	//select {
	//}
	c := make(chan os.Signal, 5)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)
	for s:=range c {
		fmt.Println(s)
		time.Sleep(5*time.Second)
		os.Exit(0)
	}
}

func forTest(){
	c := make(chan os.Signal, 5)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)
	for s:=range c {
		fmt.Println(s)
		time.Sleep(5*time.Second)
		os.Exit(0)
	}
}

var jobOver chan int

func Recovery() {
	go recoveryMapAndAll(dir + "/map")
	//go recoveryMapAndAll(dir + "/all")
	//go recoveryShutdown()
	num := 0
	for {
		select {
		case <-jobOver:
			num++
			if num == 3 {
				return
			}
		}
	}
}

func recoveryMapAndAll(path string) {
	dirs, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
	}
	//p := datasource.GetPipe()
	for _, fi := range dirs {
		file, _ := os.Open(path+"/" + fi.Name())
		reader := bufio.NewReader(file)
		for {
			bline, ok, err := reader.ReadLine()
			if err != nil {
				fmt.Println(err)
				break
			}
			for ok{
				tmp := []byte("")
				tmp = BytesCombine(tmp,bline)
				_,ok,_ = reader.ReadLine()
				bline = tmp
			}
			strline := string(bline)
			//strings.Split(strline, "\t")
			strs := strings.Split(strline, "\t")
			key := strs[0]
			value := strs[1]
			fmt.Println(key,"\t",value)
			//p.Push(buffer.NewMessage(key, []byte(value)))
		}
	}
	jobOver <- 1
}

// []byte数组拼接
func BytesCombine(pBytes ...[]byte) []byte {
	return bytes.Join(pBytes, []byte(""))
}
