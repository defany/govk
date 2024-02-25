package generator

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"
)

type ObjectsFile struct {
	Definitions map[string]Property
}

func GenerateObjects(objectsRaw []byte) {
	var file ObjectsFile

	if err := json.Unmarshal(objectsRaw, &file); err != nil {
		panic(err.Error())
	}

	nameGenners := make([]NameGennerWithoutTest, 0, len(file.Definitions))

	for name, prop := range file.Definitions {
		obj, _ := parseObjectNameGenner(name, prop, 0)
		if obj != nil {
			nameGenners = append(nameGenners, obj)
		}
	}

	sort.SliceStable(nameGenners, func(i, j int) bool {
		return nameGenners[i].GetName() < nameGenners[j].GetName()
	})
	for _, genner := range nameGenners {
		dir := "gen/" + getObjectName(genner.GetName()) + "/models"
		err := os.MkdirAll(dir, 0777)
		if err != nil {
			panic(err.Error())
		}

		if _, err := os.Stat(dir + "/" + getFullObjectName(genner.GetName()) + ".go"); errors.Is(err, os.ErrNotExist) {
			f, err := os.OpenFile(dir+"/"+getFullObjectName(genner.GetName())+".go", os.O_CREATE|os.O_APPEND, os.ModePerm)
			if err != nil {
				panic(err.Error())
			}
			writeStartFile(f, getObjectName(genner.GetName()), "", genner.GetImports(), "encoding/json")
			fmt.Fprint(f, "// suppress unused package warning\nvar _ *json.RawMessage\n\n")
			fmt.Fprint(f, genner.Gen())
		} else {
			f, err := os.OpenFile(dir+"/"+getFullObjectName(genner.GetName())+".go", os.O_APPEND, os.ModePerm)
			if err != nil {
				panic(err.Error())
			}

			fmt.Fprint(f, genner.Gen())
		}
	}
}

var SimpleReferences = make(map[string]RefType)

type RefType struct {
	ArrayNestingLevel int
	Type              string
}

type NestedGenner interface {
	nestedGen(nestingLvl int, objName string) (nestedGen string, additionalGen string)
}

type NameNestedGenner interface {
	Namer
	NestedGenner
}

func parseObjectNameGenner(name string, prop Property, arrayNestingLvl int) (NameGennerWithoutTest, []string) {
	if prop.Ref != nil {
		t := parseSimpleType(getObjectName(name), name, prop, arrayNestingLvl)
		SimpleReferences[getFullObjectName(t.Name)] = RefType{
			ArrayNestingLevel: arrayNestingLvl,
			Type:              t.Type,
		}
		return t, t.Imports
	}

	if prop.Items != nil {
		prop.Limits.Add(prop.Items.Limits)
		prop.Items.Description = prop.Description
		prop.Items.Required = prop.Required
		return parseObjectNameGenner(name, *prop.Items, arrayNestingLvl+1)
	}

	if prop.Enum != nil {
		e := parseEnum(name, prop, arrayNestingLvl)
		SimpleReferences[getFullObjectName(e.Name)] = RefType{
			ArrayNestingLevel: arrayNestingLvl,
			Type:              e.ValuesType,
		}
		return e, []string{}
	}

	if prop.AllOf != nil {
		return parseAllOf(name, prop, arrayNestingLvl)
	}

	if prop.OneOf != nil {
		return parseOneOf(name, prop, arrayNestingLvl)
	}

	if prop.Properties != nil {
		return parseObject(name, prop, arrayNestingLvl)
	}

	if prop.PatternProperties != nil {
		return parsePatternProperties(name, prop, arrayNestingLvl)
	}

	t := parseSimpleType(getObjectName(name), name, prop, arrayNestingLvl)
	SimpleReferences[getFullObjectName(t.Name)] = RefType{
		ArrayNestingLevel: arrayNestingLvl,
		Type:              t.Type,
	}
	return t, t.Imports
}

func parseNameNestedGenner(objName string, name string, prop Property, arrayNestingLvl int) (NameNestedGenner, []string) {
	if prop.Ref != nil {
		t := parseSimpleType(objName, name, prop, arrayNestingLvl)
		return t, t.Imports
	}

	if prop.Items != nil {
		prop.Limits.Add(prop.Items.Limits)
		prop.Items.Description = prop.Description
		prop.Items.Required = prop.Required
		return parseNameNestedGenner(objName, name, *prop.Items, arrayNestingLvl+1)
	}

	if prop.Enum != nil {
		return parseEnum(name, prop, arrayNestingLvl), []string{}
	}

	if prop.AllOf != nil {
		return parseAllOf(name, prop, arrayNestingLvl)
	}

	if prop.OneOf != nil {
		return parseOneOf(name, prop, arrayNestingLvl)
	}

	if prop.Properties != nil {
		return parseObject(objName, prop, arrayNestingLvl)
	}

	if prop.PatternProperties != nil {
		return parsePatternProperties(objName, prop, arrayNestingLvl)
	}

	t := parseSimpleType(objName, name, prop, arrayNestingLvl)
	return t, t.Imports
}

type Object struct {
	Name              string
	Description       string
	Fields            []NameNestedGenner
	ArrayNestingLevel int
	Imports           []string

	// fields for nested Object
	IsRequired bool
}

func (o Object) GetName() string {
	return o.Name
}

func (o Object) GetImports() []string {
	return removeDuplicateStr(o.Imports)
}

func parseObject(name string, prop Property, arrayNestingLvl int) (o Object, imp []string) {
	o.Name = name
	o.ArrayNestingLevel = arrayNestingLvl

	if prop.Description != nil {
		o.Description = *prop.Description
	}

	if prop.Required != nil {
		o.IsRequired = *prop.Required
	}

	if prop.Properties == nil {
		panic(name)
	}

	o.Fields = make([]NameNestedGenner, 0, len(*prop.Properties))

	for fieldName, fieldProp := range *prop.Properties {
		gen, im := parseNameNestedGenner(getObjectName(name), fieldName, fieldProp, 0)
		o.Fields = append(o.Fields, gen)
		if len(im) > 0 {
			o.Imports = append(o.Imports, im...)
		}
	}
	return o, o.Imports
}

func (o Object) Gen() (gen string) {
	genName := getFullObjectName(o.Name)

	sort.SliceStable(o.Fields, func(i, j int) bool {
		return o.Fields[i].GetName() < o.Fields[j].GetName()
	})

	var fieldsGen string

	for _, nestedGenner := range o.Fields {
		fGen, addGen := nestedGenner.nestedGen(1, genName)
		fieldsGen += fGen
		gen += addGen
	}

	// write comment
	if o.Description != "" {
		gen += fmt.Sprintf("// %s %s\n", genName, o.Description)
	}

	if _, easySkipped := easyJSONBlackList[o.Name]; easySkipped {
		gen += "//easyjson:skip\n"
	}
	gen += fmt.Sprintf("type %s %sstruct {\n%s}\n\n", genName, getArrayBrackets(o.ArrayNestingLevel), fieldsGen)
	return
}

func (o Object) nestedGen(nestingLvl int, objName string) (nestedGen, additionalGen string) {
	genName := getFullName(o.Name)
	newObjName := objName + genName
	if o.Name == "" {
		newObjName = objName
		nestingLvl--
	}

	sort.SliceStable(o.Fields, func(i, j int) bool {
		return o.Fields[i].GetName() < o.Fields[j].GetName()
	})

	var fieldsGen string

	for _, nestedGenner := range o.Fields {
		fGen, addGen := nestedGenner.nestedGen(nestingLvl+1, newObjName)
		fieldsGen += fGen
		additionalGen += addGen
	}

	tabs := getTabs(nestingLvl)

	// write comment
	if o.Description != "" {
		nestedGen += fmt.Sprintf("%s// %s\n", tabs, o.Description)
	}

	if o.Name == "" {
		// Unknown object
		// Like in AllOf or OneOf
		nestedGen += fieldsGen
		return
	}

	var preType, omitempty string
	if !o.IsRequired {
		omitempty = ",omitempty"
		preType = "*"
	}
	preType += getArrayBrackets(o.ArrayNestingLevel)

	genJSON := fmt.Sprintf("`json:%q`", o.Name+omitempty)

	nestedGen += fmt.Sprintf("%s%s %sstruct {\n%s%s} %s\n",
		tabs, genName, preType, fieldsGen, tabs, genJSON)

	return
}

type AllOf struct {
	Name              string
	Description       string
	Fields            []NestedGenner
	ArrayNestingLevel int
	Imports           []string
	// fields for nested AllOf
	IsRequired bool
}

func (ao AllOf) GetName() string {
	return ao.Name
}

func (ao AllOf) GetImports() []string {
	return removeDuplicateStr(ao.Imports)
}

func parseAllOf(name string, prop Property, arrayNestingLvl int) (ao AllOf, imp []string) {
	ao.Name = name
	ao.ArrayNestingLevel = arrayNestingLvl

	if prop.Description != nil {
		ao.Description = *prop.Description
	}

	if prop.Required != nil {
		ao.IsRequired = *prop.Required
	}

	if prop.AllOf == nil {
		panic(name)
	}

	ao.Fields = make([]NestedGenner, 0, len(*prop.AllOf))

	for _, allOfProp := range *prop.AllOf {
		required := true
		allOfProp.Required = &required
		gen, im := parseNameNestedGenner(getObjectName(name), name, allOfProp, 0)
		ao.Fields = append(ao.Fields, gen)
		if len(im) > 0 {
			ao.Imports = append(ao.Imports, im...)
		}
	}

	return ao, ao.Imports
}

func (ao AllOf) Gen() (gen string) {
	genName := getFullObjectName(ao.Name)

	var fieldsGen string

	for _, nestedGenner := range ao.Fields {
		fGen, addGen := nestedGenner.nestedGen(1, genName)
		fieldsGen += fGen
		gen += addGen
	}

	// write comment
	if ao.Description != "" {
		gen += fmt.Sprintf("// %s %s\n", genName, ao.Description)
	}

	if _, easySkipped := easyJSONBlackList[ao.Name]; easySkipped {
		gen += "//easyjson:skip\n"
	}

	gen += fmt.Sprintf("type %s %sstruct {\n%s}\n\n", genName, getArrayBrackets(ao.ArrayNestingLevel), fieldsGen)

	return
}

func (ao AllOf) nestedGen(nestingLvl int, objName string) (nestedGen, additionalGen string) {
	genName := getFullName(ao.Name)

	var fieldsGen string

	for _, nestedGenner := range ao.Fields {
		fGen, addGen := nestedGenner.nestedGen(nestingLvl, objName+genName)
		fieldsGen += fGen
		additionalGen += addGen
	}

	tabs := getTabs(nestingLvl)

	// write comment
	if ao.Description != "" {
		nestedGen += fmt.Sprintf("%s// %s\n", tabs, ao.Description)
	}

	if ao.Name == "" {
		// Unknown object
		// Like in AllOf or OneOf
		nestedGen += fieldsGen
		return
	}

	var preType, omitempty string
	if !ao.IsRequired {
		omitempty = ",omitempty"
		preType = "*"
	}
	preType += getArrayBrackets(ao.ArrayNestingLevel)

	genJSON := fmt.Sprintf("`json:%q`", ao.Name+omitempty)

	nestedGen += fmt.Sprintf("%s%s %sstruct {\n%s%s} %s\n",
		tabs, genName, preType, fieldsGen, tabs, genJSON)

	return
}

type OneOf struct {
	Name              string
	Description       string
	Fields            []NestedGenner
	ArrayNestingLevel int

	// fields for nested OneOf
	Imports    []string
	IsRequired bool
	isNested   bool
}

func (of OneOf) GetName() string {
	return of.Name
}

func (of OneOf) GetImports() []string {
	return removeDuplicateStr(of.Imports)
}

func parseOneOf(name string, prop Property, arrayNestingLvl int) (of OneOf, imp []string) {
	of.Name = name
	of.ArrayNestingLevel = arrayNestingLvl

	if prop.Description != nil {
		of.Description = *prop.Description
	}

	if prop.Required != nil {
		of.IsRequired = *prop.Required
	}

	if prop.OneOf == nil {
		panic(name)
	}

	of.Fields = make([]NestedGenner, 0, len(*prop.OneOf))

	for _, oneOfProp := range *prop.OneOf {
		gen, im := parseNameNestedGenner(getObjectName(name), name, oneOfProp, 0)
		of.Fields = append(of.Fields, gen)
		if len(im) > 0 {
			of.Imports = append(of.Imports, im...)
		}
	}

	return of, of.Imports
}

func (of OneOf) Gen() (gen string) {
	var genName string

	if of.isNested {
		genName = of.Name
	} else {
		genName = getFullObjectName(of.Name)
	}

	// write comment
	if of.Description != "" {
		gen += fmt.Sprintf("// %s %s\n", genName, of.Description)
	}

	// write main object
	gen += "//easyjson:skip\n"
	gen += fmt.Sprintf("type %s struct{\n", genName)
	gen += "\traw []byte\n"
	gen += "}\n\n"

	// write marshal/unmarshaler method
	gen += of.genMarshaler()
	gen += of.genUnmarshaler()

	// write json map getter
	gen += of.genRawGetter()

	return
}

func (of OneOf) genMarshaler() (gen string) {
	var genName string

	if of.isNested {
		genName = of.Name
	} else {
		genName = getFullObjectName(of.Name)
	}

	gen += fmt.Sprintf("func (o *%s) MarshalJSON() ([]byte, error) {\n", genName)

	gen += fmt.Sprintf("\treturn o.raw, nil\n")
	gen += "}\n\n"

	return
}

func (of OneOf) genUnmarshaler() (gen string) {
	var genName string

	if of.isNested {
		genName = of.Name
	} else {
		genName = getFullObjectName(of.Name)
	}

	gen += fmt.Sprintf("func (o *%s) UnmarshalJSON(body []byte) (err error) {\n", genName)

	gen += "\to.raw = body\n"
	//gen += "\treturn json.Unmarshal(body, &o.raw)\n"
	gen += "\treturn nil\n"

	gen += "}\n\n"

	return
}

func (of OneOf) genRawGetter() (gen string) {
	var genName string

	if of.isNested {
		genName = of.Name
	} else {
		genName = getFullObjectName(of.Name)
	}

	gen += fmt.Sprintf("func (o %s) Raw() []byte {\n", genName)
	gen += fmt.Sprintf("\treturn o.raw\n")
	gen += "}\n\n"

	return
}

func (of OneOf) nestedGen(tabsCount int, objName string) (gen, additionalGen string) {
	genName := getFullName(of.Name)
	newStructType := objName + genName

	mainOneOf := OneOf{
		Name:              newStructType,
		Description:       "",
		Fields:            of.Fields,
		ArrayNestingLevel: of.ArrayNestingLevel,
		isNested:          true,
	}

	additionalGen += mainOneOf.Gen()

	tabs := getTabs(tabsCount)

	// write comment
	if of.Description != "" {
		gen += fmt.Sprintf("%s// %s\n", tabs, of.Description)
	}

	var preType, omitempty string
	if !of.IsRequired {
		omitempty = ",omitempty"
		preType = "*"
	}
	preType += getArrayBrackets(of.ArrayNestingLevel)

	genJSON := fmt.Sprintf("`json:%q`", of.Name+omitempty)

	// write main object
	gen += fmt.Sprintf("%s%s %s%s %s\n", tabs, genName, preType, newStructType, genJSON)

	return
}

type SimpleType struct {
	Name              string
	Description       string
	Type              string
	Limits            Limits
	ArrayNestingLevel int
	IsRequired        bool

	withoutJSON bool
	isNesting   bool
	Imports     []string
}

func (t SimpleType) GetName() string {
	return t.Name
}

func (t SimpleType) GetImports() []string {
	return removeDuplicateStr(t.Imports)
}

func parseSimpleType(objName string, name string, prop Property, arrayNestingLvl int) (t SimpleType) {
	t.Name = name
	t.ArrayNestingLevel = arrayNestingLvl
	t.Limits = prop.Limits

	if prop.Description != nil {
		t.Description = *prop.Description
	}

	if prop.Required != nil {
		t.IsRequired = *prop.Required
	}

	if prop.Ref != nil {
		if strings.Contains(getRefName(*prop.Ref), "_") {
			els := strings.Split(getRefName(*prop.Ref), "_")
			if objName == els[0] {
				t.Type = getFullName(getRefName(*prop.Ref))
				return
			}
			t.Type = getFullName(getRefName(*prop.Ref)) + "." + getFullName(getRefName(*prop.Ref))
			t.Imports = append(t.Imports, getFullName(getRefName(*prop.Ref))+" github.com/defany/govk/gen/"+strings.Split(getRefName(*prop.Ref), "_")[0]+"/models")
			return
		}
		t.Imports = append(t.Imports, getFullName(getRefName(*prop.Ref))+" github.com/defany/govk/gen/"+strings.Split(getRefName(*prop.Ref), "_")[0]+"/models")
		t.Type = getRefName(*prop.Ref)
		return
	}

	if prop.Type != nil {
		if *prop.Type == "object" {
			panic(name)
		}

		if _, ok := (*prop.Type).([]interface{}); ok {
			t.Type = "string"
		} else {
			t.Type = getSimpleType((*prop.Type).(string))
		}
		return
	}

	panic(name)
}

func (t SimpleType) Gen() (gen string) {
	var genName string

	if t.isNesting {
		genName = t.Name
	} else {
		genName = getFullObjectName(t.Name)
	}

	// write comment
	if t.Description != "" {
		gen += fmt.Sprintf("// %s %s\n", genName, t.Description)
	}
	gen += t.Limits.gen(1)

	if _, easySkipped := easyJSONBlackList[t.Name]; easySkipped {
		gen += "//easyjson:skip\n"
	}

	// write main object
	gen += fmt.Sprintf("type %s %s%s\n\n", genName, getArrayBrackets(t.ArrayNestingLevel), getFullObjectName(t.Type))

	return
}

func (t SimpleType) nestedGen(nestingLvl int, objName string) (gen, _ string) {
	tabs := getTabs(nestingLvl)
	genType := getFullObjectName(t.Type)
	genName := getFullName(t.Name)

	if t.Name == "" {
		if !t.IsRequired {
			// OneOf style
			if t.ArrayNestingLevel > 0 {
				genName = strings.ReplaceAll(getArrayBrackets(t.ArrayNestingLevel), "[]", "Array") + genType
			} else {
				genName = upFirstAny(genType)
			}
		}
		t.withoutJSON = true
	}

	var preType, omitempty string
	if !t.IsRequired {
		omitempty = ",omitempty"
		preType = "*"
	}
	preType += getArrayBrackets(t.ArrayNestingLevel)

	if genName == "2faRequired" {
		genName = "TwoFaRequired"
	}

	// write comment
	if t.Description != "" {
		gen += fmt.Sprintf("%s// %s\n", tabs, t.Description)
	}
	gen += t.Limits.gen(nestingLvl)

	// write nested object
	if genName != "" {
		gen += fmt.Sprintf("%s%s %s%s", tabs, genName, preType, genType)
	} else {
		gen += fmt.Sprintf("%s%s%s", tabs, preType, genType)
	}

	if !t.withoutJSON {
		gen += fmt.Sprintf(" `json:%q`", t.Name+omitempty)
	}

	gen += "\n"
	return
}

type Enum struct {
	Name        string
	Description string
	ValuesType  string
	EnumValues  []interface{}
	EnumNames   []string
	Limits      Limits

	IsRequired        bool
	ArrayNestingLevel int

	isNested    bool // global or nested
	withoutJSON bool
}

func (e Enum) GetName() string {
	return e.Name
}

func (e Enum) GetImports() []string {
	return nil
}

func parseEnum(name string, prop Property, arrayNestingLvl int) (e Enum) {
	e.Name = name
	e.ArrayNestingLevel = arrayNestingLvl
	e.Limits = prop.Limits

	if prop.Description != nil {
		e.Description = *prop.Description
	}

	if prop.Required != nil {
		e.IsRequired = *prop.Required
	}

	if prop.Enum == nil {
		panic(name)
	}

	e.EnumValues = *prop.Enum

	if prop.EnumNames != nil {
		e.EnumNames = *prop.EnumNames
	}

	// TODO: костыль
	if vType := fmt.Sprintf("%T", e.EnumValues[0]); vType == "float64" {
		e.ValuesType = "int"
	} else {
		e.ValuesType = vType
	}

	return
}

func (e Enum) Gen() (gen string) {
	var genName string

	if e.isNested {
		genName = e.Name
	} else {
		genName = getFullObjectName(e.Name)
	}

	//write comment
	if e.Description != "" {
		gen += fmt.Sprintf("// %s %s\n", genName, e.Description)
	}

	if _, easySkipped := easyJSONBlackList[e.Name]; easySkipped {
		gen += "//easyjson:skip\n"
	}

	// write main object
	gen += fmt.Sprintf("type %s %s\n\n", genName, e.ValuesType)

	consts := make([]string, 0, len(e.EnumValues))

	for i, v := range e.EnumValues {
		var suffix string
		if len(e.EnumNames) == 0 {
			suffix = fmt.Sprintf("%v", v)
		} else {
			suffix = editEnumSpace(e.EnumNames[i])
		}

		var genValue string
		if e.ValuesType == "string" {
			genValue = fmt.Sprintf("%q", v)
		} else {
			genValue = fmt.Sprintf("%v", v)
		}

		c := fmt.Sprintf("\t%s%s %s = %s", genName, getFullName(suffix), genName, genValue)

		consts = append(consts, c)
	}

	genConsts := strings.Join(consts, "\n")

	gen += fmt.Sprintf("const (\n%s\n)\n\n", genConsts)

	return
}

func (e Enum) nestedGen(nestingLvl int, objName string) (nestedGen, additionalGen string) {
	genName := getFullName(e.Name)

	newStructType := objName + genName

	mainEnum := Enum{
		Name:        newStructType,
		Description: "",
		ValuesType:  e.ValuesType,
		EnumValues:  e.EnumValues,
		EnumNames:   e.EnumNames,
		isNested:    true,
	}

	additionalGen = mainEnum.Gen()

	// generate nested enum field
	tabs := getTabs(nestingLvl)

	// write comment
	if e.Description != "" {
		nestedGen += fmt.Sprintf("%s// %s\n", tabs, e.Description)
	}
	nestedGen += e.Limits.gen(nestingLvl)

	var preType, omitempty string
	if !e.IsRequired {
		omitempty = ",omitempty"
		preType = "*"
	}
	preType += getArrayBrackets(e.ArrayNestingLevel)

	// write nested enum
	nestedGen += fmt.Sprintf("%s%s %s%s", tabs, genName, preType, newStructType)

	if !e.withoutJSON {
		nestedGen += fmt.Sprintf("`json:%q`", e.Name+omitempty)
	}

	nestedGen += "\n"

	return
}

func parsePatternProperties(name string, prop Property, arrayNestingLvl int) (t SimpleType, imp []string) {
	t.Name = name
	t.ArrayNestingLevel = arrayNestingLvl
	t.Limits = prop.Limits

	if prop.Description != nil {
		t.Description = *prop.Description
	}

	if prop.Required != nil {
		t.IsRequired = *prop.Required
	}

	if prop.PatternProperties == nil {
		panic(name)
	}

	ref := (*prop.PatternProperties)["^[0-9]+$"]["$ref"]
	els := strings.Split(getRefName(ref), "_")
	if name != els[0] {
		t.Imports = append(t.Imports, getFullName(getRefName(ref))+" github.com/defany/govk/gen/"+strings.Split(getRefName(ref), "_")[0]+"/models")
		t.Type = "map[string]" + getRefName(ref) + "." + getRefName(ref)
		return t, t.Imports
	}
	t.Type = "map[string]" + getRefName(ref)

	return t, t.Imports
}
