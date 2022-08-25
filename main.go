package main

import (
	cert "example/gencert/cert"
	"example/gencert/cert/html"
	"example/gencert/cert/pdf"
	"flag"
	"fmt"
	"os"
)

func main() {

	outputType := flag.String("type", "pdf", "Output type of the certificate.")
	file := flag.String("file", "", "csv file to parse")
	flag.Parse()

	if len(*file) <= 0 {
		fmt.Printf("Invalid file, got : %v\n", file)
		os.Exit(1)
	}

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

	certs, err := cert.ParseCsv(*file)
	if err != nil {
		fmt.Printf("Could not parse CSV file : %v\n", err)
		os.Exit(1)
	}

	for _, c := range certs {
		err = saver.Save(*c)
		if err != nil {
			fmt.Printf("Could not save Cert : %v\n", err)
		}
	}

}
