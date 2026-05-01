package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func newProxy(target string) *httputil.ReverseProxy{
	parseUrl, _:=url.Parse(target)
	return httputil.NewSingleHostReverseProxy(parseUrl)
}


func auth(next http.Handler) http.Handler{

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token:=r.Header.Get("Authorization")

		if token !="secret-token"{
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return 
		}

		next.ServeHTTP(w, r)
	})
}


func  main()  {
	
	userService:=newProxy("http://127.0.0.1:9003")
	orderService:=newProxy("http://127.0.0.1:9004")


	mux:=http.NewServeMux()


	mux.Handle("/users/", auth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userService.ServeHTTP(w, r)
	})))


	mux.Handle("/orders", auth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		orderService.ServeHTTP(w, r)

	})))

	log.Println("Api running on : 8081")
	log.Fatal(http.ListenAndServe(":8081", mux))
}