package main

import (
	"github.com/skateboard/tls-client/service"
	"sync"
)

func main() {
	go service.StartService()

	// low cpu usage when using wait groups!
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
