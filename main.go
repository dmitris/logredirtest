package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/dmitris/stdmock"
)

func wrapFoobar() {
	log.SetOutput(ioutil.Discard)
	stdmock.Foobar()
	log.SetOutput(os.Stderr)
}

func main() {
	log.Println("before...")
	// stdmock.Foobar()
	wrapFoobar()
	log.Println("after...")
}
