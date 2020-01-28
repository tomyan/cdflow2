package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func getEnv() map[string]string {
	result := make(map[string]string)
	for _, item := range os.Environ() {
		kv := strings.SplitN(item, "=", 2)
		result[kv[0]] = kv[1]
	}
	return result
}

func main() {

	if len(os.Args) == 3 && os.Args[1] == "workspace" && os.Args[2] == "list" {
		fmt.Println("* default")
		fmt.Println("  existing-workspace")
	} else {
		fmt.Println("message to stdout")
	}

	encoder := json.NewEncoder(os.Stderr)

	fileContents, err := ioutil.ReadFile("/code/mapped-dir-test")
	if err != nil {
		log.Fatalln("could not read file:", err)
	}

	var input bytes.Buffer
	io.Copy(os.Stdin, &input)

	dir, err := os.Getwd()
	if err != nil {
		log.Fatalln("could not get working directory:", err)
	}

	encoder.Encode(map[string]interface{}{
		"args": os.Args[1:], "env": getEnv(), "input": input.String(),
		"cwd": dir, "file": string(fileContents),
	})

	if len(os.Args) == 3 && os.Args[1] == "init" && os.Args[2] == "/code/infra" {
		if err := ioutil.WriteFile("build-output-test", []byte("build output"), 0644); err != nil {
			log.Fatalln("could not write file:", err)
		}
	}

	if err := ioutil.WriteFile("/code/source-dir-write-test", []byte("source output"), 0644); err == nil {
		log.Fatalln("was able to write to file:", err)
	}
}
