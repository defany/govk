package main

import (
	"errors"
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
	genObjects("objects_test.go")
	genResponses("responses_test.go")
	genMethods()
	//genVK()
}

func genErrors(file string) {
	dir := "api/gen/tests"
	err := os.MkdirAll(dir, 0777)
	if err != nil {
		panic(err.Error())
	}
	var f *os.File
	if _, err := os.Stat("api/gen/tests/" + file); errors.Is(err, os.ErrNotExist) {
		f, err = os.OpenFile("api/gen/tests/"+file, os.O_CREATE|os.O_APPEND, os.ModePerm)
		if err != nil {
			panic(err.Error())
		}
	}
	generator.GenerateErrors(f, getRawFromAddr(errorsFile))
}

func genObjects(file string) {
	dir := "api/gen/tests"
	err := os.MkdirAll(dir, 0777)
	if err != nil {
		panic(err.Error())
	}
	var f *os.File
	if _, err := os.Stat("api/gen/tests/" + file); errors.Is(err, os.ErrNotExist) {
		f, err = os.OpenFile("api/gen/tests/"+file, os.O_CREATE|os.O_APPEND, os.ModePerm)
		if err != nil {
			panic(err.Error())
		}
	}

	generator.GenerateObjects(f, getRawFromAddr(objectsFile))
}

func genResponses(file string) {
	dir := "api/gen/tests"
	err := os.MkdirAll(dir, 0777)
	if err != nil {
		panic(err.Error())
	}
	var f *os.File
	if _, err := os.Stat("api/gen/tests/" + file); errors.Is(err, os.ErrNotExist) {
		f, err = os.OpenFile("api/gen/tests/"+file, os.O_CREATE|os.O_APPEND, os.ModePerm)
		if err != nil {
			panic(err.Error())
		}
	}

	generator.GenerateObjects(f, getRawFromAddr(responsesFile))
}

func genMethods() {
	generator.GenerateMethods(getRawFromAddr(methodsFile))
}

func genVK() {
	generator.GenerateVK(getRawFromAddr(methodsFile))
}
