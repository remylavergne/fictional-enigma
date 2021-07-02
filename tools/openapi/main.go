package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
)

var swaggerDef string = "https://raw.githubusercontent.com/openapitools/openapi-generator/master/modules/openapi-generator/src/test/resources/3_0/petstore.yaml"
var wg sync.WaitGroup = sync.WaitGroup{}

func main() {
	log.Println("Open API Generator pour la boilerplate Afelio")

	parseArgs()

	err := os.MkdirAll("out/kotlin-multiplatform", 0755)
	if err != nil {
		panic(err)
	}

	generatorCmd := "java -jar openapi-generator-cli.jar generate -g kotlin -i " + swaggerDef + " --library multiplatform --model-name-suffix Dto -o out/kotlin-multiplatform"
	wg.Add(1)
	go func() {
		defer wg.Done()

		cmd := exec.Command("/bin/sh", "-c", generatorCmd)
		out, err := cmd.CombinedOutput()
		if err != nil {
			log.Fatal(err)
		}

		log.Println(string(out))
	}()

	wg.Wait()

	moveFilesToSharedModule()
}

func parseArgs() {
	args := os.Args[1:]

	for i, arg := range args {
		if arg == "-1" && i+1 < len(args) {
			swaggerDef = args[i+1]
		}
	}
}

func moveFilesToSharedModule() {
	errr := filepath.Walk("./out/kotlin-multiplatform/src",
		func(subpath string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			// Ignore top level path
			if subpath == "./out/kotlin-multiplatform/src" {
				return nil
			}

			// Change path
			log.Println("Path / file found:", subpath)
			newPath := strings.Replace(subpath, "out/kotlin-multiplatform/src", "../../shared/src", 1)

			log.Println("New path / file to be created:", newPath)

			if info.IsDir() {
				err := os.MkdirAll(newPath, 0755)
				if err != nil {
					panic(err)
				}
				log.Println("Directory", info.Name(), "created to", newPath)
			} else {
				read, err := ioutil.ReadFile(subpath)
				if err != nil {
					panic(err)
				}

				err = ioutil.WriteFile(newPath, []byte(read), 0755)
				if err != nil {
					panic(err)
				}

				// filesMoved += 1
				log.Println("File", info.Name(), "moved to", newPath)
			}

			return nil
		})
	if errr != nil {
		log.Println(errr)
	}
}
