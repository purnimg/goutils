/* ProtocolBuffer 
 */

package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"os"
	"person"
)

func main() {
	name := person.Person_Name{
		Family: "Gunturu",
		Personal: "Purnima"}

	email1 := person.Person_Email{
		Kind: "Work",
		Address: "xxx1@gmail.com"}
	email2 := person.Person_Email{
		Kind: "Home",
		Address: "xxx2@gmail.com"}

	emails := []*person.Person_Email{&email1, &email2}
	p := person.Person{
		Name : &name,
		Email: emails,
	}

	fmt.Println(p)

	data, err := proto. Marshal(&p)
	checkError(err)
	newP := person.Person{}
	err = proto.Unmarshal(data, &newP)
	checkError(err)
	fmt.Println(newP)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
