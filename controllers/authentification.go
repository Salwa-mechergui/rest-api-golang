package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//endpoint creation
func CreateEndpoint(w http.ResponseWriter, r *http.Request) {

	response, err := http.Post(os.Getenv("URLAUTH"), "application/json", r.Body)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
		w.Write(data)
		return
	}

}
