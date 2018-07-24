package main

import (
	"log"
	"net/http"

	"github.com/jspc/loadtest"
)

func main() {
	idb, err := NewInfluxdbCollector("http://localhost:8086", "magnum")
	if err != nil {
		panic(err)
	}

	collectors := []Collector{
		idb,
	}

	c := make(chan loadtest.Output)

	go func() {
		for o := range c {
			for _, collector := range collectors {
				err = collector.Push(o)
				if err != nil {
					log.Print(err)
				}
			}
		}
	}()

	api := API{
		OutputChan: c,
	}

	panic(http.ListenAndServe(":8082", api))
}