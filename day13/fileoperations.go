package main 
import (
	
	"os"
)


func main(){
	file, _:=os.Create("test.txt")
	defer file.Close()

	file.WriteString("Hello Go. I am learning you! \n")

	data:=[]byte("Byte content \n")
	file.Write(data)
	


}