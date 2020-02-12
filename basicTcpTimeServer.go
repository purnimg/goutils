/* Time server
*/

package main

import (
        "fmt"
        "net"
        "os"
       "time"
)

func main(){
        service := ":1200"
        tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
        checkError(err)

        listener, err := net.ListenTCP("tcp4", tcpAddr)
        checkError(err)

        for{ //keep listening on the connection
                conn, err := listener.Accept()
                if err!= nil{
                      continue
                }

                daytime := time.Now().String()
                conn.Write([]byte(daytime))  //dont care about retvalue 
                conn.Close()
        }
}

func checkError(err error){
     if err!=nil{
           fmt.Fprintf(os.Stderr, "fatal error: %s", err.Error())
           os.Exit(1)
		 }
}
