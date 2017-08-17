package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func main() {
	http.HandleFunc("/devices", ifconfigHandler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("incoming request: ", r)
		w.Write([]byte("Hello there! \n Go to /devices to see list of network devices"))
	})
	fmt.Print("Server started on port : 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

var execCommand = exec.Command

func ListDevices() ([]byte, error) {
	cmd := execCommand("ifconfig")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("\n Error while execution command: %s , error: %s", "ifconfig", err)
		return []byte{}, err
	}
	//log.Println(string(out))
	return out, nil

}
func ifconfigHandler(w http.ResponseWriter, r *http.Request) {
	out, err := ListDevices()
	if err != nil {
		log.Println("Execution 'ifconfig' error: ", err)
		http.Error(w, fmt.Sprintf("Execution 'ifconfig' error: %s", err), http.StatusInternalServerError)
	}
	w.Write(out)
}
