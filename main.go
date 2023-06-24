package main

import (
	"bufio"
	filesys "csvDB/FileSys"
	"fmt"
	"net"
	"os"
	"time"
)

var test []string = []string{"hdsjl"}

const (
	maxThreads = 4
	port       = ":1337"
	masterDir  = "./Databases"
)

func main() {
	update := NewUpdate()
	kA := KeepAlive{}
	listener, _ := net.Listen("tcp", port)
	threads := new(threads)

	var master *Master
	if !filesys.IsDir(masterDir) {
		//make your new Database
		fmt.Printf("Please enter in the following order initial:[databasename tablename value{seperated by \":\"}]\n >> ")
		reader := bufio.NewReader(os.Stdin)
		bytes, _ := reader.ReadBytes('\n')
		str := ""
		for i := 0; i < len(bytes); i++ {
			str += (string(bytes[i]))
		}
		if str[len(str)-1] == '\n' {
			newStr := ""
			for i := 0; i < len(str)-1; i++ {
				newStr += string(str[i])
			}
			str = newStr
		}
		Sstr := filesys.Split(str, " ")
		if len(Sstr) >= 3 {
			master = CreateMaster(masterDir, &update, Sstr[0], Sstr[1], filesys.Split(Sstr[2], ":"))
		} else {
			update.Stop()
			fmt.Println("error: to few arguments")
		}
	} else {
		master = LoadMaster(masterDir, &update)
	}

	for i := 0; i < maxThreads; i++ {
		go thread(threads, &update, master, &kA, i)
		time.Sleep(time.Millisecond * 5)
	}
	go update.Upd(&kA)

	fmt.Println("sever initialized")
	for !update.stop {
		conn, _ := listener.Accept()
		threads.add(&conn)
	}

	fmt.Println("Server Stopped")
	//time.Sleep(time.Second * 2)
	kA.Wait()
}
