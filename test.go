package main

import (
	"flag"
	"fmt"
	humanize "github.com/dustin/go-humanize"
	"math/rand"
	"time"
)

const (
	buffersize    = 1024 * 10
	storedbuffers = 500
	intsize       = 8
	workingset    = storedbuffers * buffersize * intsize
)

var (
	randomdata       = false
	randomsize       = false
	maxalloc   int64 = 1024
)

func main() {
	flag.BoolVar(&randomdata, "rand", randomdata, "Fill buffers with random data after allocation.")
	flag.BoolVar(&randomsize, "randsize", randomsize, "Create buffers with random (bounded) size.")
	flag.Int64Var(&maxalloc, "mem", maxalloc, "The program will stop after this many MB have been allocated.")
	flag.Parse()
	rand.Seed(time.Now().Unix())
	maxalloc *= 1024 * 1024

	var i int64 = 0
	var allocated int64 = 0
	var buffers [storedbuffers][]int
	for allocated*intsize < maxalloc {
		var size int64
		if randomsize {
			size = int64(rand.Int() % buffersize)
		} else {
			size = buffersize
		}
		newbuffer := make([]int, size)
		buffers[i%storedbuffers] = newbuffer
		allocated += size
		if randomdata {
			for j := range newbuffer {
				newbuffer[j] = rand.Int()
			}
		}
		fmt.Printf("Allocated %v buffers, total storage: %v.\n", i, humanize.Comma(allocated*intsize))
		i++
	}
	var realWorkingSet int64 = workingset
	if randomsize {
		// Divide by two because the allocated buffer sizes are random, so this is the average
		realWorkingSet /= 2
	}
	fmt.Printf("working set %v.\n", humanize.Comma(realWorkingSet))
}
