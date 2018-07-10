package main

import (
	"os"
)

func main() {
	a := App{}

	// use db connection creds stored in env vars
	a.Initialize(
		os.Getenv("postgres"),
		os.Getenv("debian23"),
		os.Getenv("gobase"),
		os.Getenv("no"),
	)

	a.Run(":8080")
}
