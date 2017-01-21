package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	// Trace ...
	Trace *log.Logger
	// Info ...
	Info *log.Logger
	// Warning ...
	Warning *log.Logger
	// Error ...
	Error *log.Logger
)

// Init ...
func Init(
	traceHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {

	Trace = log.New(traceHandle,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)

	var (
		mode = flag.String("mode", "dev", "Application mode")
	)

	flag.Parse()

	Info.Printf(fmt.Sprintf("App mode=\"%s\"", *mode))
}
