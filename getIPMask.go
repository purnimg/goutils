/* 
IP: Takes IP mask and subnet number and prints network details
*/

package main

import (
				"fmt"
				"net"
				"os"
				"strconv"
)

func main(){
				if len(os.Args) !=4 {
								fmt.Fprintf(os.Stderr, "Usage: %s dotted-ip-addr ones bits \n", os.Args[0])
								os.Exit(1)
				}

				dotAddr := os.Args[1]
				ones, _ := strconv.Atoi(os.Args[2])
				bits, _ := strconv.Atoi(os.Args[3])

				addr := net.ParseIP(dotAddr)

				if addr == nil {
								fmt.Println("Invalid Address")
								os.Exit(1)
				}

			
				if addr == nil {
								fmt.Println("Invalid Address")
				} else {
								fmt.Println("The address is ", addr.String())
				}
				os.Exit(0)
}

