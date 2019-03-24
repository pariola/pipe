package main

import (
	"flag"
	"fmt"
	"net/http"
	"os/exec"
	"runtime"
)

func main() {
	if runtime.GOOS == "windows" {
		fmt.Println("Can't Execute this on a windows machine")
	} else {
		execute()
	}
}

func execute() {
	scriptPath := flag.String("script", "./push.sh", "a string")
	port := flag.String("port", "3000", "a string")
	flag.Parse()

	http.HandleFunc("/push", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			fmt.Fprint(w, "OK")
			fmt.Println("Push Recieved")

			out, _ := exec.Command(*scriptPath).Output()
			fmt.Println(string(out))
		} else {
			fmt.Println("Forbidden")
		}

	})

	fmt.Println("Server starting!!")
	http.ListenAndServe(":"+*port, nil)
	fmt.Println("Server started!!")
}
