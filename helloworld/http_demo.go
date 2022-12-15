package main

import "fmt"

type person struct {
	Name string
}

func main() {
	//http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
	//	fmt.Println("hello")
	//})
	//var svr http.Server
	//svr.Addr = ":9999"
	//svr.ListenAndServe()
	//http.ListenAndServe(":9999", nil)
	var p person
	p.Name = "zhangsan"

	fmt.Println(p)
}
