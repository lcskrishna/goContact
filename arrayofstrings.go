package main 
import(
	"encoding/csv"
	"os"
	"log"
	"fmt"
)

func main() {
	file,err := os.Open("contactlist.csv")
	checkError("Cannot read file",err)

		// Get a CSV Reader Object.
	reader := csv.NewReader(file)
	reader.Comma = ','

	records,err:= reader.ReadAll()
	checkError("Read Error", err)
	
	file.Close()

	var err1 = os.Remove("contactlist.csv")
	checkError("Delete Error", err1)

	csvfile, err2 := os.Create("contactlist.csv")
	checkError("Creation error",err2)
	writer := csv.NewWriter(csvfile)

	
	data := [][]string{{"3","Nikita","9803185718","testemail@yahoo.com","UT","0"}}

	for _, record := range records {
		fmt.Println("ENtered")
		fmt.Println("Value of teh DAta is :",record)
		err:= writer.Write(record)
		checkError("Cannot Write into the file",err)
	}

	for _,value := range data {
		fmt.Println("New Data")
		err:= writer.Write(value)
		checkError("Cannot Write into the file",err)
	}

	 writer.Flush()
	 csvfile.Close()

}

func checkError(message string, err error){
	if err!= nil {
		log.Fatal(message,err)
	}

}