package main

import (
	"log"
	"os"

	"github.com/appropriate/go-virtualboxclient/vboxwebsrv"
)

func main() {
	url := "http://127.0.0.1:18083"
	if len(os.Args) >= 2 {
		url = os.Args[1]
	}

	client := vboxwebsrv.NewVboxPortType(url, false, nil)
	response, err := client.IWebsessionManagerlogon(&vboxwebsrv.IWebsessionManagerlogon{})
	if err != nil {
		log.Fatalf("Unable to log on to vboxwebsrv: %v\n", err)
	}

	log.Printf("returnval=%v\n", response.Returnval)
}
