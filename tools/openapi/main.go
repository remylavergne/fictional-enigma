package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"sync"
)

var swaggerDef string = "https://raw.githubusercontent.com/openapitools/openapi-generator/master/modules/openapi-generator/src/test/resources/3_0/petstore.yaml"

func main() {
	log.Println("Open API Generator pour la boilerplate Afelio")

	os.MkdirAll("out/kotlin-multiplatform", 0755)

	var wg sync.WaitGroup = sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		// cmd := exec.Command("java", "-jar", "openapi-generator-cli.jar", "generate", "-g", "kotlin", "-i", swaggerDef, "--library", "multiplatform", "--model-name-suffix", "Dto", "-o", "/out/kotlin-multiplatform")
		cmd := exec.Command("/bin/sh", "-c", "java -jar openapi-generator-cli.jar generate -g kotlin -i https://raw.githubusercontent.com/openapitools/openapi-generator/master/modules/openapi-generator/src/test/resources/3_0/petstore.yaml --library multiplatform --model-name-suffix Dto -o out/kotlin-multiplatform")
		// cmd.Start()
		// cmd.Wait()
		out, err := cmd.CombinedOutput()
		log.Println("a")
		// err := cmd.Run()
		log.Println("b")
		if err != nil {
			fmt.Println("an error occurred.")
			log.Fatal(err)
		}
		log.Println(string(out))
		// log.Println(string(out))
	}()

	wg.Wait()
}
