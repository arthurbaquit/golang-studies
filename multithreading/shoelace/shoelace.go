package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Point2D struct {
	X float64
	Y float64
}

const numberOfThreads = 4

var (
	pointRegex = regexp.MustCompile(`\((\d*),(\d*)\)`)
	wg         = sync.WaitGroup{}
)

func findArea(pointChannel chan string) {
	for pointStr := range pointChannel {
		var points []Point2D
		for _, p := range pointRegex.FindAllStringSubmatch(pointStr, -1) {
			x, _ := strconv.ParseFloat(p[1], 64)
			y, _ := strconv.ParseFloat(p[2], 64)
			points = append(points, Point2D{x, y})
		}
		area := 0.0
		for i := 0; i < len(points); i++ {
			area += points[i].X*points[(i+1)%len(points)].Y - points[i].Y*points[(i+1)%len(points)].X
		}
		area = math.Abs(area) / 2.0
		fmt.Println(area)
	}
	wg.Done()
}

func main() {
	filePath, _ := filepath.Abs("multithreading/shoelace/polygon.txt")
	data, _ := ioutil.ReadFile(filePath)
	pointChannel := make(chan string)
	for i := 0; i < numberOfThreads; i++ {
		wg.Add(1)
		go findArea(pointChannel)
	}
	startTime := time.Now()
	for _, line := range strings.Split(string(data), "\n") {
		pointChannel <- line
	}
	close(pointChannel)
	wg.Wait()
	fmt.Println("Time elapsed: ", time.Since(startTime))

}
