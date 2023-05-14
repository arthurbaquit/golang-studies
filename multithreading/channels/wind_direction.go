package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Iremos fazer um pipeline de canais para calcular a direção do vento.
//

var (
	windReportRegex = regexp.MustCompile(`\d* METAR.*EGLL \d*Z [A-Z ]*(\d{5}KT|VRB\d{2}KT).*=`)
	tafValidation   = regexp.MustCompile(`.*TAF.*`)
	comment         = regexp.MustCompile(`\w*#.*`)
	metarClose      = regexp.MustCompile(`.*=`)
	variableWind    = regexp.MustCompile(`.*VRB\d{2}KT`)
	validWind       = regexp.MustCompile(`\d{5}KT`)
	windDirOnly     = regexp.MustCompile(`(\d{3})\d{2}KT`)
	windDist        [8]int
)

func parseToArray(reportChannel chan string, metarChannel chan []string) {
	for report := range reportChannel {
		lines := strings.Split(report, "\n")
		metarSlice := make([]string, 0, len(lines))
		metarStr := ""
		for _, line := range lines {
			if tafValidation.MatchString(line) {
				break
			}
			if !comment.MatchString(line) {
				metarStr += strings.Trim(line, " ")
			}
			if metarClose.MatchString(line) {
				metarSlice = append(metarSlice, metarStr)
				metarStr = ""
			}
		}
		metarChannel <- metarSlice
	}
	close(metarChannel)
}

func extractWindDirection(metarChannel chan []string, windChannel chan []string) {
	for metars := range metarChannel {
		winds := make([]string, 0, len(metars))
		for _, metar := range metars {
			if windReportRegex.MatchString(metar) {
				winds = append(winds, windReportRegex.FindAllStringSubmatch(metar, -1)[0][1])
			}
		}
		windChannel <- winds
	}
	close(windChannel)
}

func mineWindDistribution(windChannel chan []string, resultsChannel chan [8]int) {
	for winds := range windChannel {
		for _, wind := range winds {
			if variableWind.MatchString(wind) {
				for i := 0; i < 8; i++ {
					windDist[i]++
				}
			} else if validWind.MatchString(wind) {
				windStr := windDirOnly.FindAllStringSubmatch(wind, -1)[0][1]

				if d, err := strconv.ParseFloat(windStr, 64); err == nil {
					dirIndex := int(math.Round(d/45.0)) % 8
					windDist[dirIndex]++
				}
			}
		}
	}
	resultsChannel <- windDist

	close(resultsChannel)
}

func main() {
	// primeiro vamos ter um channel para receber o dado e parsea para array
	// um channel que extrai as informações necessárias do array
	// Por fim, um que cria a distribuição.

	reportChannel := make(chan string)
	metarChannel := make(chan []string)
	windChannel := make(chan []string)
	resultsChannel := make(chan [8]int)

	go parseToArray(reportChannel, metarChannel)
	go extractWindDirection(metarChannel, windChannel)
	go mineWindDistribution(windChannel, resultsChannel)

	filePath, _ := filepath.Abs("./multithreading/channels/windReports")
	files, _ := ioutil.ReadDir(filePath)
	startTime := time.Now()
	for _, file := range files {
		dat, err := ioutil.ReadFile(filePath + "/" + file.Name())
		fmt.Println(file.Name())
		if err != nil {
			panic(err)
		}
		reportChannel <- string(dat)
	}
	close(reportChannel)
	results := <-resultsChannel
	fmt.Printf("Total time: %s\n", time.Since(startTime))
	fmt.Println("Wind distribution: ", results)

}
