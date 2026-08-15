package main 
import (
	"fmt"
	"log"
	"net/http"
)



func main(){
	http.HandleFunc("/hello", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}



func handler(w http.ResponseWriter, r *http.Request){

	if(r.Method!=http.MethodGet){
		http.Error(w, "Method not allowerd", http.StatusMethodNotAllowed)
		return 
	}

	name:=r.URL.Query().Get("name")


	if len(name)==0{
		http.Error(w,"Name is required", http.StatusBadRequest)
		return 
	}

	if len(name)>50{
		http.Error(w,"Name is too long", http.StatusBadRequest)
		return 
	}

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Hello, %s!", name)
}
