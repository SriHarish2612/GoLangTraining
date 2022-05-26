package main

import (
	"Movie/API"
	"Movie/WebJob"
)

func main() {
	API.Start()
	WebJob.Job()
}
