package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func check(err error, message string) {
	if err != nil{
		panic(err)
	}
	fmt.Printf("%s\n", message)
}

type ClientJob struct{
	name string
	conn net.Conn
}

func generateResponses( ClientJobs chan ClientJob){
	for{
		//wait for the next job to come off the queue.
		ClientJob := <-ClientJobs

		//Do something that keeps the cpu buys for a whole second.
		for start := time.Now(); time.Now().Sub(start) < time.Second;{

		}
		//send back the response.
		ClientJob.conn.Write([]byte("Hello, " + ClientJob.name))
	}
}

func main() {
	ClientJobs := make(chan ClientJob)
	go generateResponses(ClientJobs)

	ln, err := net.Listen("tcp", ":8080")
	check(err, "server is ready.")

	for {
		conn, err:= ln.Accept()
		check(err, "Accepted connection.")

		go func(){
			buf:= bufio.NewReader(conn)

			for {
				name,err:= buf.ReadString('\n')

				if err != nil{
					fmt.Printf("Client disconnected.\n")
					break
				}

				ClientJobs <- ClientJob{name, conn}
			}
		}()
	}
}