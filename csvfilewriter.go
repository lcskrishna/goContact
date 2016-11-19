package main

import (
	"os"
	"log"
	"fmt"
	"reflect"
)

//var data = [][] string {{"Line1","Hello Readers of :"},{"Line2","golangcode.com"}}

func main(){
	file, err := os.Create("contactlist.csv")
	checkError("Cannot create file",err)
	fmt.Println(reflect.TypeOf(file))

	/*writer:= csv.NewWriter(file)

	for _, value := range data {
		fmt.Println("Value of teh DAta is :",value)
		err:= writer.Write(value)
		checkError("Cannot Write into the file",err)

	}

	defer writer.Flush()
	*/
	fmt.Println("CSV File is Created")
}

func checkError(message string, err error){
	if err!= nil {
		log.Fatal(message,err)
	}

}