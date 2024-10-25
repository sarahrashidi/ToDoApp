package main

import (
	"fmt"
	"net/http"
)

var shortGolang = "Watch Go crash course"
var fullGolang = "Watch Go Full Course"
var rewardDessert = "Reward myself with a game"
var taskItems = []string{shortGolang, fullGolang, rewardDessert}

func main() {

	http.HandleFunc("/", helloUser)
	http.HandleFunc("/show-tasks", showTasks)

	http.ListenAndServe(":8080", nil)

}

func helloUser(writer http.ResponseWriter, request *http.Request) {
	var greeting = "Hello user! Welcome to Todolist App!"
	fmt.Fprintln(writer, greeting)

}

func showTasks(writer http.ResponseWriter, request *http.Request) {
	for _, task := range taskItems {
		fmt.Fprintln(writer, task)

	}

}
