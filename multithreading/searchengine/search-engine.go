package main

import (
	"io/ioutil"
	"strings"
	"sync"
)

var (
	matches []string
	wg      sync.WaitGroup
	mutex   sync.Mutex
)

func lookAtFolder(folderPath, searchFile string) {
	defer wg.Done()
	files, err := ioutil.ReadDir(folderPath)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if file.IsDir() {
			wg.Add(1)
			go lookAtFolder(folderPath+"/"+file.Name(), searchFile)
		} else {
			if strings.Contains(file.Name(), searchFile) {
				mutex.Lock()
				matches = append(matches, folderPath+"/"+file.Name())
				mutex.Unlock()
				println(folderPath + "/" + file.Name())
			}
		}
	}
}

func main() {
	wg.Add(1)
	go lookAtFolder("/Users/arthurbaquit/estudos/golang-studies", "README.md")
	wg.Wait()
	// Notice that, in Boid, we didn't use wg. The reason for that is, in boid, we
	// didn't need the sync between the goroutines. As soon as one point is finished,
	// we want to compute it and write it in the screen. However, in this case, we
	// need to wait for all the goroutines to finish before printing the results.
	// Therefore, the waitgroup is used to sync the data.
}
