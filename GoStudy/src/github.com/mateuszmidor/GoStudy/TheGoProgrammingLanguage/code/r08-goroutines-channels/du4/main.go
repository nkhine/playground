// Project: kind of linux du -sh . tool, but showing the progress and with cancel option by pressing ENTER
// Usage: go run .
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

var verbose = flag.Bool("v", true, "show filesystem scan progress")
var done = make(chan struct{})

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// async read any key
	go func() {
		os.Stdin.Read(make([]byte, 1))
		fmt.Println("CANCELLED")
		close(done) // close the channel so the select in "cancelled" returns with true
	}()

	// async walk the directories
	fileSizes := make(chan int64)
	go func() {
		for _, root := range roots {
			walkDir(root, fileSizes)
		}
		close(fileSizes)
	}()

	// utility to print progress every 500ms
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}

	// calc total size printing progress every 500ms, with cancel option
	var nfiles, nbytes int64
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop // fileSizes channel was closed
			}
			nfiles++
			nbytes += size
			printDiskUsage(nfiles, nbytes)
		case <-tick:
		case <-done:
			//drain the channel so others finis
			for range fileSizes {
			}
		}
	}

	printDiskUsage(nfiles, nbytes)
}

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}
func walkDir(dir string, sizes chan<- int64) {
	if cancelled() {
		return
	}
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, sizes)
		} else {
			sizes <- entry.Size()
		}
	}
}

func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du4: %d\n", err)
		return nil
	}
	return entries
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files, %.1f GB\n", nfiles, float64(nbytes)/1e9)
}
