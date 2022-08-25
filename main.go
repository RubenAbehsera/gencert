package main

import (
	cert "example/gencert/cert"
	"example/gencert/html"
	"example/gencert/pdf"
	"flag"
	"fmt"
	"os"
)

func main() {

	outputType := flag.String("type", "pdf", "Output type of the certificate.")
	flag.Parse()

	var saver cert.Saver
	var err error
	switch *outputType {
	case "html":
		// Generate HTML
		saver, err = html.New("output")
	case "pdf":
		// Generate PDF
		saver, err = pdf.New("output")
	default:
		fmt.Printf("Unknown output type. Got:  %v\n", *outputType)
	}

	if err != nil {
		fmt.Printf("Could not create generator : '%v'", err)
		os.Exit(1)
	}

	c, err := cert.New("Golang programming", "Bob Dylan", "2018-06-21")
	if err != nil {
		fmt.Printf("Error during certificate creation : %v", err)
		os.Exit(1)
	}

	err = saver.Save(*c)
	if err != nil {
		return
	}

}
