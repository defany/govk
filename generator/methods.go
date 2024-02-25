package generator

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Parameter struct {
	Property
	Name string `json:"name"`
}

type MethodJSON struct {
	Name            string              `json:"name"`
	Description     string              `json:"description"`
	AccessTokenType []string            `json:"access_token_type"`
	Parameters      []Parameter         `json:"parameters"`
	Responses       map[string]Property `json:"responses"`
	Errors          []Property          `json:"errors"`
}

type MethodsFile struct {
	Methods []MethodJSON `json:"methods"`
}

func GenerateMethods(methodsRaw []byte) {
	var file MethodsFile

	if err := json.Unmarshal(methodsRaw, &file); err != nil {
		panic(err.Error())
	}

	genners := make([]GennerWithoutTest, 0, len(file.Methods))
	resMap := make(map[string]string)

	for _, mJSON := range file.Methods {
		if response, ok := mJSON.Responses["response"]; ok {
			resMap[mJSON.Name] = getRefName(*response.Ref)
		}
		genners = append(genners, parseMethodGenner(mJSON)...)
	}

	for _, genner := range genners {
		dir := "gen/" + getMethodName(genner.GetName())
		err := os.MkdirAll(dir, 0777)
		if err != nil {
			panic(err.Error())
		}
		if _, err := os.Stat(dir + "/" + getMethodName(genner.GetName()) + ".go"); errors.Is(err, os.ErrNotExist) {
			f, err := os.OpenFile(dir+"/"+getMethodName(genner.GetName())+".go", os.O_CREATE|os.O_APPEND, os.ModePerm)
			if err != nil {
				panic(err.Error())
			}
			writeStartFile(f, getMethodName(genner.GetName()), "", nil, "github.com/defany/govk/api")
			writeStartMethodFile(f, getMethodName(genner.GetName()))
		}
		if _, err := os.Stat(dir + "/" + getFullMethodName(genner.GetName()) + ".go"); errors.Is(err, os.ErrNotExist) {
			f, err := os.OpenFile(dir+"/"+getFullMethodName(genner.GetName())+".go", os.O_CREATE|os.O_APPEND, os.ModePerm)
			if err != nil {
				panic(err.Error())
			}
			writeStartFile(f, getMethodName(genner.GetName()), "", nil, "github.com/defany/govk/api", fmt.Sprintf("github.com/defany/govk/gen/%s/models", getObjectName(resMap[genner.GetName()])))
			fmt.Fprint(f, genner.Gen())
		} else {
			f, err := os.OpenFile(dir+"/"+getFullMethodName(genner.GetName())+".go", os.O_APPEND, os.ModePerm)
			if err != nil {
				panic(err.Error())
			}

			fmt.Fprint(f, genner.Gen())
		}
	}
}

// TODO: поработать с импортами

type ParamNameNestedGenner interface {
	NameNestedGenner
	Param() Param
}

type Param struct {
	Name              string
	Type              string
	HasCustomType     bool
	IsRequired        bool
	ArrayNestingLevel int
	Limits            Limits
}

func (t SimpleType) Param() (p Param) {
	p.Name = t.Name
	p.ArrayNestingLevel = t.ArrayNestingLevel
	p.IsRequired = t.IsRequired
	p.Type = t.Type
	p.Limits = t.Limits

	return
}

func (e Enum) Param() (p Param) {
	p.Name = e.Name
	p.Type = e.ValuesType
	p.HasCustomType = true
	p.ArrayNestingLevel = e.ArrayNestingLevel
	p.IsRequired = e.IsRequired
	p.Limits = e.Limits

	return
}

func parseParamNestedGenner(param Parameter, arrayNestingLvl int) ParamNameNestedGenner {
	if param.Ref != nil {
		t := parseSimpleType(param.Name, param.Name, param.Property, arrayNestingLvl)
		t.withoutJSON = true
		return t
	}

	if param.Items != nil {
		// TODO: научиться прокидывать
		param.Property.Items.Limits.Add(param.Property.Limits)
		param.Property = *param.Property.Items
		return parseParamNestedGenner(param, arrayNestingLvl+1)
	}

	if param.Enum != nil {
		e := parseEnum(param.Name, param.Property, arrayNestingLvl)
		e.withoutJSON = true
		return e
	}

	t := parseSimpleType(param.Name, param.Name, param.Property, arrayNestingLvl)
	t.withoutJSON = true
	return t
}

type Method struct {
	Name         string
	VKDevName    string
	Description  string
	AccessTokens []string
	ErrorRefs    []string

	GenRequest  bool
	RequestName string
	FullName    string

	Params      []ParamNameNestedGenner
	ResponseRef *string
	SetFields   []SetField
}

type SetField struct {
	Name  string
	Value string
}

func (m Method) GetName() string {
	return m.VKDevName
}

func parseMethodGenner(mJSON MethodJSON) []GennerWithoutTest {
	var m Method

	m.Name = mJSON.Name
	m.VKDevName = mJSON.Name
	m.Description = mJSON.Description
	m.AccessTokens = mJSON.AccessTokenType

	// parse params
	m.Params = make([]ParamNameNestedGenner, 0, len(mJSON.Parameters))
	for _, parameter := range mJSON.Parameters {
		m.Params = append(m.Params, parseParamNestedGenner(parameter, 0))
	}

	// parse error references
	m.ErrorRefs = make([]string, 0, len(mJSON.Errors))
	for _, e := range mJSON.Errors {
		m.ErrorRefs = append(m.ErrorRefs, getRefName(*e.Ref))
	}

	// parse responses
	additionalMethods := m.parseResponses(mJSON.Responses)

	return append([]GennerWithoutTest{m}, additionalMethods...)
}

func (m *Method) parseResponses(responses map[string]Property) []GennerWithoutTest {
	m.FullName = getFullMethodName(m.Name)
	m.RequestName = m.FullName + "Request"
	m.GenRequest = true

	if response, ok := responses["response"]; ok {
		ref := getRefName(*response.Ref)
		m.ResponseRef = &ref
	}

	if _, ok := responses["multiResponse"]; ok {
		return nil
	}

	if responseIntegerProp, ok := responses["responseInteger"]; ok {
		responseInteger := getRefName(*responseIntegerProp.Ref)
		responseArray := getRefName(*responses["responseArray"].Ref)

		countersMethod := Method{
			Name:         m.Name,
			VKDevName:    m.VKDevName,
			Description:  m.Description,
			AccessTokens: m.AccessTokens,
			ErrorRefs:    m.ErrorRefs,
			GenRequest:   true,
			RequestName:  m.FullName + "CountersRequest",
			FullName:     m.FullName + "Counters",
			Params:       m.Params,
			ResponseRef:  &responseArray,
			SetFields:    nil,
		}

		m.ResponseRef = &responseInteger

		notSecureMethod := Method{
			Name:         "setCounter",
			VKDevName:    m.VKDevName,
			Description:  m.Description,
			AccessTokens: m.AccessTokens,
			ErrorRefs:    m.ErrorRefs,
			GenRequest:   true,
			RequestName:  m.FullName + "NotSecureRequest",
			FullName:     m.FullName + "NotSecure",
			Params:       m.Params,
			ResponseRef:  &responseInteger,
			SetFields:    nil,
		}

		return []GennerWithoutTest{countersMethod, notSecureMethod}
	}

	if userIdsResponseProp, ok := responses["userIdsResponse"]; ok {
		userIdsResponse := getRefName(*userIdsResponseProp.Ref)

		if userIdsExtendedResponseProp, ok := responses["userIds_Extended_Response"]; ok {

			userIdsMethod := Method{
				Name:         m.Name,
				VKDevName:    m.VKDevName,
				Description:  m.Description,
				AccessTokens: m.AccessTokens,
				ErrorRefs:    m.ErrorRefs,
				GenRequest:   true,
				RequestName:  m.FullName + "UserIDsRequest",
				FullName:     m.FullName + "UserIDs",
				Params:       m.Params,
				ResponseRef:  &userIdsResponse,
				SetFields: []SetField{
					{
						Name:  "extended",
						Value: "0",
					},
				},
			}

			userIdsExtendedResponse := getRefName(*userIdsExtendedResponseProp.Ref)
			userIdsExtendedMethod := Method{
				Name:         m.Name,
				VKDevName:    m.VKDevName,
				Description:  m.Description,
				AccessTokens: m.AccessTokens,
				ErrorRefs:    m.ErrorRefs,
				GenRequest:   false,
				RequestName:  userIdsMethod.RequestName,
				FullName:     m.FullName + "ExtendedUserIDs",
				Params:       userIdsMethod.Params,
				ResponseRef:  &userIdsExtendedResponse,
				SetFields: []SetField{
					{
						Name:  "extended",
						Value: "1",
					},
				},
			}

			m.SetFields = []SetField{
				{
					Name:  "extended",
					Value: "0",
				},
			}

			extendedResponse := getRefName(*responses["extendedResponse"].Ref)
			extendedMethod := Method{
				Name:         m.Name,
				VKDevName:    m.VKDevName,
				Description:  m.Description,
				AccessTokens: m.AccessTokens,
				ErrorRefs:    m.ErrorRefs,
				GenRequest:   false,
				RequestName:  m.RequestName,
				FullName:     m.FullName + "Extended",
				Params:       m.Params,
				ResponseRef:  &extendedResponse,
				SetFields: []SetField{
					{
						Name:  "extended",
						Value: "1",
					},
				},
			}
			return []GennerWithoutTest{extendedMethod, userIdsMethod, userIdsExtendedMethod}
		}

		userIdsMethod := Method{
			Name:         m.Name,
			VKDevName:    m.VKDevName,
			Description:  m.Description,
			AccessTokens: m.AccessTokens,
			ErrorRefs:    m.ErrorRefs,
			GenRequest:   true,
			RequestName:  m.FullName + "UserIDsRequest",
			FullName:     m.FullName + "UserIDs",
			Params:       m.Params,
			ResponseRef:  &userIdsResponse,
			SetFields:    nil,
		}

		return []GennerWithoutTest{userIdsMethod}
	}

	if targetUidsResponseProp, ok := responses["targetUidsResponse"]; ok {
		targetUidsResponse := getRefName(*targetUidsResponseProp.Ref)

		targetUidsMethod := Method{
			Name:         m.Name,
			VKDevName:    m.VKDevName,
			Description:  m.Description,
			AccessTokens: m.AccessTokens,
			ErrorRefs:    m.ErrorRefs,
			GenRequest:   true,
			RequestName:  m.FullName + "TargetUIDsRequest",
			FullName:     m.FullName + "TargetUIDs",
			Params:       m.Params,
			ResponseRef:  &targetUidsResponse,
			SetFields:    nil,
		}

		return []GennerWithoutTest{targetUidsMethod}
	}

	if onlineMobileResponseProp, ok := responses["onlineMobileResponse"]; ok {
		m.SetFields = []SetField{
			{
				Name:  "online_mobile",
				Value: "0",
			},
		}

		onlineMobileResponse := getRefName(*onlineMobileResponseProp.Ref)

		onlineMobileMethod := Method{
			Name:         m.Name,
			VKDevName:    m.VKDevName,
			Description:  m.Description,
			AccessTokens: m.AccessTokens,
			ErrorRefs:    m.ErrorRefs,
			GenRequest:   false,
			RequestName:  m.RequestName,
			FullName:     m.FullName + "OnlineMobile",
			Params:       m.Params,
			ResponseRef:  &onlineMobileResponse,
			SetFields: []SetField{
				{
					Name:  "online_mobile",
					Value: "1",
				},
			},
		}

		return []GennerWithoutTest{onlineMobileMethod}
	}

	if extendedResponse, ok := responses["extendedResponse"]; ok {
		m.SetFields = []SetField{
			{
				Name:  "extended",
				Value: "0",
			},
		}
		extendedRef := getRefName(*extendedResponse.Ref)
		extendedM := Method{
			Name:         m.Name,
			VKDevName:    m.Name,
			Description:  m.Description,
			AccessTokens: m.AccessTokens,
			ErrorRefs:    m.ErrorRefs,
			GenRequest:   false,
			RequestName:  m.FullName + "Request",
			FullName:     m.FullName + "Extended",
			Params:       m.Params,
			ResponseRef:  &extendedRef,
			SetFields: []SetField{
				{
					Name:  "extended",
					Value: "1",
				},
			},
		}
		return []GennerWithoutTest{extendedM}
	}

	return nil
}

const docLink = "https://dev.vk.com/method/"

func (m Method) Gen() (gen string) {

	if m.Description != "" {
		gen += fmt.Sprintf("// %s %s\n", m.FullName, m.Description)
	} else {
		gen += fmt.Sprintf("// %s ...\n", m.FullName)
	}

	var genResp string
	if m.ResponseRef != nil {
		genResp = fmt.Sprintf("(resp %s, err error)", getObjectName(*m.ResponseRef)+"."+getFullObjectName(*m.ResponseRef))
	} else {
		genResp = "err error"
	}

	if m.GenRequest && len(m.Params) > 0 {
		gen += m.genRequest()
	}

	if m.AccessTokens != nil {
		gen += fmt.Sprintf("// May execute with listed access token types:\n//    [ %s ]\n", strings.Join(m.AccessTokens, ", "))
	}

	if len(m.ErrorRefs) != 0 {
		gen += fmt.Sprintf("// When executing method, may return one of global or with listed codes API errors:\n//    [ %s ]\n", buildPossibleErrors(m.ErrorRefs))
	} else {
		gen += fmt.Sprintf("// When executing method, may return one of global API errors.\n")
	}

	gen += fmt.Sprintf("//\n// %s%s\n", docLink, m.VKDevName)

	genBody := m.genBody()

	gen += fmt.Sprintf("func (%s *%s) %s(params ...api.MethodParams) %s {\n%s}\n\n", strings.ToLower(string([]rune(m.RequestName)[0:1])), upFirstAny(strings.Split(m.Name, ".")[0]), m.FullName, genResp, genBody)

	return
}

func (m Method) genBody() (body string) {
	body += fmt.Sprintf("\treq := api.NewRequest[%s](%s.api)\n\n", getObjectName(*m.ResponseRef)+"."+getFullObjectName(*m.ResponseRef), strings.ToLower(string([]rune(m.RequestName)[0:1])))

	body += fmt.Sprintf("\tres, err := req.Execute(\"%s\", api.ParamsOrNil(params))\n\tif err != nil {\n\t\treturn res, err\n\t}\n\n\treturn res, nil\n", m.Name)

	return
}

func (m Method) genRequest() (gen string) {
	gen += fmt.Sprintf("type %s api.Params\n\n", m.RequestName)
	gen += fmt.Sprintf("func New%[1]s() %[1]s {\n\tparams := make(%[1]s, %d)\n\treturn params\n}\n\n", m.RequestName, len(m.Params)+len(m.SetFields)+1)

	for _, pGenner := range m.Params {
		gen += m.genFieldParam(pGenner.Param())
	}

	gen += fmt.Sprintf("func (%s %s) Params() api.Params {\n\treturn api.Params(%[1]s)\n}\n\n", strings.ToLower(string([]rune(m.RequestName)[0:1])), m.RequestName)
	return
}

func (m Method) genFieldParam(param Param) string {
	if !isGoType(param.Type) {
		str := strings.Split(param.Type, ".")[0]
		re := regexp.MustCompile("[A-Z][a-z]*")
		slice := re.FindAllString(str, -1)
		addImport("gen/"+strings.Split(m.Name, ".")[0]+"/"+m.FullName+".go", fmt.Sprintf("\t%s \"github.com/defany/govk/gen/%s/models\"\n", str, strings.ToLower(slice[0])))
	}
	return fmt.Sprintf("func (%s %s) With%s(%[1]s_%[5]s %[6]s%[4]s) %[2]s{\n\t%[1]s[\"%[5]s\"] = %[1]s_%[5]s\n\treturn %[1]s\n}\n\n", strings.ToLower(string([]rune(m.RequestName)[0:1])), m.RequestName, getFullObjectName(param.Name), param.Type, param.Name, getArrayBrackets(param.ArrayNestingLevel))
}
