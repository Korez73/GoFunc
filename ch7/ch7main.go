package ch7

import "net/http"

// This func must be Exported, Capitalized, and comment added.

func Main() {
	//LogOutput("HI")

	//this will replace the call to LogOutput
	//should call the API???
	//sl := NewSimpleLogic()

	l := LoggerAdapter(LogOutput)
	ds := NewSimpleDataStore()
	logic := NewSimpleLogic(l, ds)
	c := NewController(l, logic)

	http.HandleFunc("/hello", c.SayHello)
	http.ListenAndServe(":8080", nil)
}
