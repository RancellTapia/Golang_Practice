package main

import (
	"log"
	"net/http"
)

func main (){
	http.HandleFunc("/", func(wr http.ResponseWriter, r *http.Request) {
		wr.Write([]byte (
			`<html>
				<head>
					<title>Chat</title>
				</head>
				
					<h1>Let's chat</h1>
				
			</html>`,
		))
	})	

	// Start the web server

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Listen And Serve: ", err)
	}
}