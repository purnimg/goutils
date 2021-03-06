/* Threaded Simple TCP client
*/


package main

import (
				"fmt"
				"net"
				"os"
)

func main(){
				service := ":2100"
				tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
				checkErr(err)
				
				listener, err := net.ListenTCP("tcp4", tcpAddr)
				checkErr(err)

				for{
								conn, err := listener.Accept()
								if err!= nil {
												continue
								}
								//run as a goroutine
								go handleClient(conn)
				}
}

func handleClient(conn net.Conn){
				//close connection on exit
				defer conn.Close()

				var buf [512]byte
				for {
								//read up to 512 bytes
								n, err := conn.Read(buf[0:])
								if err!=nil{
												return
								}
								fmt.Println(string(buf[0:]))
								//write the n bytes read
								_, err2 := conn.Write(buf[0:n])
								if err2!=nil{
												return
								}
				}
}

func checkErr(err error){
				if err!=nil {
								fmt.Println("Fatal Error: %s", err.Error())
								os.Exit(1)
				}
}

