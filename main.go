package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func main() {

	http.HandleFunc("/vnstat", func(w http.ResponseWriter, r *http.Request) {
		cmd := exec.Command("vnstat", "-s")
		stdout, _ := cmd.Output()
		fmt.Fprintf(w, "%s", stdout)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
