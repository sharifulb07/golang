package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func newproxy(target string) *httputil.ReverseProxy{

	parseUrl, err:=url.Parse(target)

	if err !=nil{
		log.Fatal(err)
	}

	proxy:=httputil.NewSingleHostReverseProxy(parseUrl)

	proxy.ModifyResponse=func(r *http.Response) error {
			r.Header.Set("X-Proxy", "Go-Reverse-Proxy")
		return nil 
	}

	return proxy
}



func startAPIServer(){

	mux := http.NewServeMux()

	mux.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Routing to api server ")
		
	})

	log.Println("api server is running on :9001")
	log.Fatal(http.ListenAndServe(":9001", nil))

}


func startWebServer(){

	mux := http.NewServeMux()

	mux .HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Routing to webserver")
		

	})

	log.Println("web server is running on :9002")
	log.Fatal(http.ListenAndServe(":9002", nil))

}
func main() {
	mux := http.NewServeMux()

	go startAPIServer()
	go startWebServer()

	apiServer := newproxy("http://127.0.0.1:9001")
	webServer := newproxy("http://127.0.0.1:9002")

	mux.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Routing to api server ")
		apiServer.ServeHTTP(w, r)
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Routing to webserver")
		webServer.ServeHTTP(w, r)

	})

	log.Println("Reverse Proxy is running on 8081")
	log.Fatal(http.ListenAndServe(":8081", mux))





}