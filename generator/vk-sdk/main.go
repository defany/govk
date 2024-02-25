package main

import (
	"flag"
	"github.com/defany/govk/generator"
	"io"
	"net/http"
	"os"
	"os/exec"
)

const (
	methodsFile   = `https://raw.githubusercontent.com/elias506/vk-api-schema/master/methods.json`
	errorsFile    = `https://raw.githubusercontent.com/elias506/vk-api-schema/master/errors.json`
	objectsFile   = `https://raw.githubusercontent.com/elias506/vk-api-schema/master/objects.json`
	responsesFile = `https://raw.githubusercontent.com/elias506/vk-api-schema/master/responses.json`
)

func getRawFromAddr(addr string) []byte {
	resp, err := http.Get(addr)

	if err != nil {
		panic(err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		panic("code is no OK: " + resp.Status)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err.Error())
	}

	return body
}

func goFmt(path string) {
	if *noformat {
		return
	}

	cmd := exec.Command("go", "fmt", path)

	cmd.Run()
}

var noformat = flag.Bool("noformat", false, "do not run 'gofmt -w' on output file")

func main() {
	flag.Parse()

	genErrors("error_codes.go")
	genObjects()
	genResponses()
	genMethods("methods.go")
}

func genErrors(file string) {
	defer goFmt(file)

	e, err := os.Create(file)

	if err != nil {
		panic(err.Error())
	}
	defer e.Close()
	generator.GenerateErrors(e, getRawFromAddr(errorsFile))
}

func genObjects() {
	generator.GenerateObjects(getRawFromAddr(objectsFile))
}

func genResponses() {
	generator.GenerateObjects(getRawFromAddr(responsesFile))
}

func genMethods(file string) {
	generator.GenerateMethods(getRawFromAddr(methodsFile))
}
