package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"
	"log"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9573")
	if err != nil {
		fmt.Println("Error connecting:", err)
		os.Exit(1)
	}
	defer conn.Close()
	fmt.Println("Connecting to 127.0.0.1:9573")

	var wg sync.WaitGroup
	wg.Add(2)
	go handleWrite(conn, &wg)
	go handleRead(conn, &wg)
	wg.Wait()
}
func handleWrite(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, err := inputReader.ReadString(byte('\n'))
		if err != nil {
			log.Panic(err)
		}
		if input == "exit\n" {
			fmt.Println("good bye")
			os.Exit(1)
		}

		_, err = conn.Write([]byte(input))
		if err != nil {
			log.Panic(err)
		}
	}
}
func handleRead(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	reader := bufio.NewReader(conn)
	for {
		line, err := reader.ReadString(byte('\n'))
		if err != nil {
			fmt.Print("Error to read message because of ", err)
			return
		}
		fmt.Print(line)
	}
}
