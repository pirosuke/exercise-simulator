package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pirosuke/exercise-simulator/internal/models"
	"github.com/pirosuke/exercise-simulator/internal/simulator"
)

func fileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil
}

func main() {

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: exersim [flags]\n")
		flag.PrintDefaults()
	}

	planFilePath := flag.String("p", "", "Plans file path")
	flag.Parse()

	if !fileExists(*planFilePath) {
		fmt.Println("plans file does not exist")
		return
	}

	jsonContent, err := ioutil.ReadFile(*planFilePath)
	if err != nil {
		fmt.Println("failed to read plans file")
		return
	}

	plansConfig := new(models.PlansConfig)
	if err := json.Unmarshal(jsonContent, plansConfig); err != nil {
		fmt.Println("failed to read plans file")
		return
	}

	result := simulator.SimulatePlans(plansConfig, *planFilePath)

	resultJSON, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("failed to create result data")
		return
	}

	fmt.Println(string(resultJSON))
}
