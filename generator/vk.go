package generator

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

func GenerateVK(methodsRaw []byte) {
	var file MethodsFile

	if err := json.Unmarshal(methodsRaw, &file); err != nil {
		panic(err.Error())
	}

	addProvider := make([]string, 40)
	addFields := make([]string, 40)
	addImports := make([]string, 40)
	for _, mJSON := range file.Methods {
		if slices.Contains(addImports, fmt.Sprintf("%sReqs \"github.com/defany/govk/api/gen/%[1]s\"\n", getMethodName(mJSON.Name))) {
			continue
		}
		addImports = append(addImports, fmt.Sprintf("%sReqs \"github.com/defany/govk/api/gen/%[1]s\"\n", getMethodName(mJSON.Name)))
		addFields = append(addFields, fmt.Sprintf("%s: %sReqs.New%[1]s(api),\n", upFirstAny(strings.Split(mJSON.Name, ".")[0]), getMethodName(mJSON.Name)))
		addProvider = append(addProvider, fmt.Sprintf("%s *%sReqs.%[1]s", upFirstAny(strings.Split(mJSON.Name, ".")[0]), getMethodName(mJSON.Name)))
	}
	dir := "vk/"
	err := os.MkdirAll(dir, 0777)
	if err != nil {
		panic(err.Error())
	}
	f, err := os.OpenFile(dir+"/vk.go", os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(err.Error())
	}
	writeFile(f, addFields, addProvider, addImports)
}

func writeFile(f io.Writer, addFields, addProvider, addImports []string) {
	fmt.Fprint(f, "package govk\n\n")
	fmt.Fprint(f, "import (\n\t\"github.com/defany/govk/api\"\n\t\"github.com/defany/govk/updates\"\n\tapiModel \"github.com/defany/govk/api/model\"\n")
	for _, addImport := range addImports {
		fmt.Fprint(f, "\t"+addImport)
	}
	fmt.Fprint(f, ")\n")

	fmt.Fprint(f, "type VK struct {\n\tApi *apiModel.ApiProvider\n\tUpdates *updates.Updates\n}\n")
	fmt.Fprint(f, "func NewVK(tokens ...string) (*VK, error) {\n\tapi, err := api.NewAPI(tokens...)\n\tif err != nil {\n\t\treturn nil, err\n\t}\n")
	fmt.Fprint(f, "\tapiModel := &apiModel.ApiProvider{\n\t\tApi: api,\n")
	for _, addField := range addFields {
		fmt.Fprint(f, fmt.Sprintf("\t\t%s", addField))
	}
	fmt.Fprint(f, "\t}")
	fmt.Fprint(f, "\n\n\tupdates := updates.NewUpdates(apiModel)\n\n\t")
	fmt.Fprint(f, "return &VK{\n\t\tApi: apiModel,\n\t\tUpdates: updates,\n\t}, nil\n}")
	fmt.Fprint(f, "\ntype apiModel struct {\n\t")
	for _, provider := range addProvider {
		fmt.Fprint(f, "\t"+provider+"\n")
	}
	fmt.Fprint(f, "}")
}
