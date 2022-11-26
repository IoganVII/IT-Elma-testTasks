package main

import (
	"finalFantasyTask/arrayOperationService"
	"net/http"
)

func main() {
	err := http.ListenAndServe("localHost:8080", arrayOperationService.NewService())
	if err != nil {
		return
	}
}
