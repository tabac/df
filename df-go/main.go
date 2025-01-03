package main

/*
#cgo LDFLAGS: -L./lib -ldf_rs_ffi
#include "./lib/df.h"
#include <stdlib.h>
*/
import "C"

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/tabac/df/client"
	"github.com/tabac/df/server"
)

func main() {
	if len(os.Args) < 2 {
		usage()

		return
	}

	switch os.Args[1] {
	case "client":
		var (
			runs    int    = 5
			network string = "tcp"
		)

		if len(os.Args) == 3 {
			network = os.Args[2]
		}

		client, err := client.New(network)
		if err != nil {
			log.Fatal(err)
		}

		var wg sync.WaitGroup

		for i := 0; i < runs; i++ {
			wg.Add(1)

			go func(id int) {
				defer wg.Done()

				err = client.Run(id)
				if err != nil {
					log.Fatal(err)
				}
			}(i)
		}

		wg.Wait()
	case "server":
		var network string = "tcp"

		if len(os.Args) == 3 {
			network = os.Args[2]
		}

		server := server.New(network)

		err := server.Run()
		if err != nil {
			log.Fatal(err)
		}
	case "ffi-tcp":
		C.start_tcp_ffi()
	case "ffi-unix":
		C.start_unix_ffi()
	}
}

func usage() {
	fmt.Println("df-go [client/server]")
}
