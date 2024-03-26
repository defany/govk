package generator

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
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
		dir := "api/gen/" + strings.Split(genner.GetName(), ".")[0]
		err := os.MkdirAll(dir, 0777)
		if err != nil {
			panic(err.Error())
		}
		if _, err := os.Stat(dir + "/" + strings.Split(genner.GetName(), ".")[0] + ".go"); errors.Is(err, os.ErrNotExist) {
			f, err := os.OpenFile(dir+"/"+strings.Split(genner.GetName(), ".")[0]+".go", os.O_CREATE|os.O_APPEND, os.ModePerm)
			if err != nil {
				panic(err.Error())
			}
			writeStartFile(f, "requests", "", "github.com/defany/govk/api")
			writeStartMethodFile(f, strings.Split(genner.GetName(), ".")[0])
		}

		if _, err := os.Stat("api/gen/tests/" + strings.Split(genner.GetName(), ".")[0] + "_test" + ".go"); errors.Is(err, os.ErrNotExist) {
			f, err := os.OpenFile("api/gen/tests/"+strings.Split(genner.GetName(), ".")[0]+"_test"+".go", os.O_CREATE|os.O_APPEND, os.ModePerm)
			if err != nil {
				panic(err.Error())
			}
			writeStartFile(f, "tests", "", "github.com/defany/govk/api/gen/models", fmt.Sprintf("github.com/defany/govk/api/gen/%s", getMethodName(genner.GetName())), "github.com/stretchr/testify/assert", "github.com/stretchr/testify/require", "github.com/defany/govk/vk", "testing", "encoding/json")
			fmt.Fprint(f, genner.TestGen())
		} else {
			f, err := os.OpenFile("api/gen/tests/"+strings.Split(genner.GetName(), ".")[0]+"_test"+".go", os.O_APPEND, os.ModePerm)
			if err != nil {
				panic(err.Error())
			}
			fmt.Fprint(f, genner.TestGen())
		}

		if _, err := os.Stat(dir + "/" + getFullMethodName(genner.GetName()) + ".go"); errors.Is(err, os.ErrNotExist) {
			f, err := os.OpenFile(dir+"/"+getFullMethodName(genner.GetName())+".go", os.O_CREATE|os.O_APPEND, os.ModePerm)
			if err != nil {
				panic(err.Error())
			}
			writeStartFile(f, "requests", "", "github.com/defany/govk/api", "github.com/defany/govk/api/gen/models")
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
		t := parseSimpleType(param.Name, param.Property, arrayNestingLvl)
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

	t := parseSimpleType(param.Name, param.Property, arrayNestingLvl)
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
	Value bool
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
						Value: false,
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
						Value: true,
					},
				},
			}

			m.SetFields = []SetField{
				{
					Name:  "extended",
					Value: false,
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
						Value: true,
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
				Value: false,
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
					Value: true,
				},
			},
		}

		return []GennerWithoutTest{onlineMobileMethod}
	}

	if extendedResponse, ok := responses["extendedResponse"]; ok {
		m.SetFields = []SetField{
			{
				Name:  "extended",
				Value: false,
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
					Value: true,
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
		genResp = fmt.Sprintf("(resp %s, err error)", "models."+getFullObjectName(*m.ResponseRef))
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
	body += fmt.Sprintf("\treq := api.NewRequest[%s](%s.api)\n\n", "models."+getFullObjectName(*m.ResponseRef), strings.ToLower(string([]rune(m.RequestName)[0:1])))

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
	var paramType string
	if !isGoType(param.Type) {
		paramType = "models." + param.Type
	} else {
		paramType = param.Type
	}
	return fmt.Sprintf("func (%s %s) With%s(%[1]s_%[5]s %[6]s%[4]s) %[2]s{\n\t%[1]s[\"%[5]s\"] = %[1]s_%[5]s\n\treturn %[1]s\n}\n\n", strings.ToLower(string([]rune(m.RequestName)[0:1])), m.RequestName, getFullObjectName(param.Name), paramType, param.Name, getArrayBrackets(param.ArrayNestingLevel))
}

func (m Method) TestGen() (testGen string) {
	if m.GenRequest && len(m.Params) > 0 {
		testGen += m.testGenRequest()
	}

	testGen += m.successTestGen()
	//testGen += m.apiErrorTestGen()
	//testGen += m.errorTestGen()

	return
}

func (m Method) successTestGen() (testGen string) {
	testGen += fmt.Sprintf("func TestVK%sSuccess(t *testing.T) {\n", m.FullName)

	if len(m.Params) > 0 {
		testGen += fmt.Sprintf("\tparams := requests.New%s()\n",
			m.RequestName)
		testGen += fmt.Sprintf("\t%s%s(&params)\n", testMethodName, m.RequestName)
	}

	for _, st := range m.SetFields {
		testGen += fmt.Sprintf("\tparams.With%s(%t)\n", upFirstAny(getFullObjectName(st.Name)), st.Value)
	}

	if m.ResponseRef != nil {
		resp := getFullObjectName(*m.ResponseRef)
		testGen += fmt.Sprintf("\tvar expected models.%s\n", resp)
		testGen += fmt.Sprintf("\t%s%s(&expected)\n", testMethodName, resp)
		testGen += "\texpectedJSON, err := json.Marshal(expected)\n"
		testGen += "\trequire.NoError(t, err)\n"
	} else {
		testGen += fmt.Sprintf("\texpectedJSON := []byte(%q)", "{}")
	}

	testGen += "\ttoken := randString()\n"
	testGen += fmt.Sprintf("\tvk, err := govk.NewVK(token)\n")
	testGen += "\tassert.NoError(t, err)\n"
	if len(m.Params) > 0 {
		testGen += fmt.Sprintf("\tvk.Api.WithHTTP(NewTestClient(t, %q, params.Params(), expectedJSON))\n", m.Name)
	} else {
		testGen += fmt.Sprintf("\tvk.Api.WithHTTP(NewTestClient(t, %q, nil, expectedJSON))\n", m.Name)
	}

	if len(m.Params) > 0 {
		testGen += fmt.Sprintf("\tresp, err := vk.Api.%s.%s(params)\n", upFirstAny(strings.Split(m.VKDevName, ".")[0]), m.FullName)
	} else {
		testGen += fmt.Sprintf("\tresp, err := vk.Api.%s.%s()\n", upFirstAny(strings.Split(m.VKDevName, ".")[0]), m.FullName)
	}

	testGen += "\tassert.EqualValues(t, expected, resp)\n"
	testGen += "\tassert.NoError(t, err)\n"

	testGen += "}\n\n"

	return
}

//func (m Method) apiErrorTestGen() (testGen string) {
//	testGen += fmt.Sprintf("func TestVK%sApiError(t *testing.T) {\n", m.FullName)
//
//	testGen += fmt.Sprintf("\tvar expected apiError\n")
//	testGen += "\texpected.fillRandomly()\n"
//	testGen += "\texpectedJSON, err := json.Marshal(expected)\n"
//	testGen += "\trequire.NoError(t, err)\n"
//
//	testGen += "\ttoken := randString()\n"
//	testGen += fmt.Sprintf("\tvk, err := govk.NewVK(token)\n")
//	testGen += fmt.Sprintf("\tvk.Api.WithHTTP(NewApiErrorTestClient(t, %q, expectedJSON))\n", m.Name)
//	testGen += "\tassert.NoError(t, err)\n"
//
//	testGen += fmt.Sprintf("\tresp, err := vk.Api.%s.%s()\n", upFirstAny(getMethodName(m.VKDevName)), m.FullName)
//	// TODO: доделать
//	testGen += "\tassert.Empty(t, resp)\n"
//	testGen += "\tassert.Equal(t, &expected, resp)\n"
//	testGen += "\tassert.NoError(t, err)\n"
//
//	testGen += "}\n\n"
//
//	return
//}
//
//func (m Method) errorTestGen() (testGen string) {
//	testGen += fmt.Sprintf("func TestVK%sError(t *testing.T) {\n", m.FullName)
//
//	testGen += "\texpected := errors.New(randString())\n"
//	testGen += "\ttoken := randString()\n"
//	testGen += fmt.Sprintf("\tvk, err := govk.NewVK(token)\n")
//	testGen += fmt.Sprintf("\tvk.Api.WithHTTP(NewErrorTestClient(expected))\n")
//
//	testGen += fmt.Sprintf("\tresp, err := vk.Api.%s.%s()\n", upFirstAny(getMethodName(m.VKDevName)), m.FullName)
//
//	testGen += "\tassert.Empty(t, resp)\n"
//	testGen += "\tassert.Nil(t, resp)\n"
//	testGen += "\tassert.ErrorIs(t, err, expected)\n"
//
//	testGen += "}\n\n"
//
//	return
//}

func (m Method) testGenRequest() (testGen string) {
	var fieldsGen string

	for _, pGenner := range m.Params {
		if isGoType(pGenner.Param().Type) {
			if pGenner.Param().ArrayNestingLevel < 1 {
				fieldsGen += fmt.Sprintf("\tr.With%s(%s)\n", getFullName(pGenner.GetName()), getRandSetter(pGenner.Param().Type))
			} else {
				fieldsGen += fmt.Sprintf("\tl%s := randIntn(maxArrayLength + 1)\n", getFullName(pGenner.GetName()))
				fieldsGen += fmt.Sprintf("\tr.With%s(%sl%[1]s))\n", getFullName(pGenner.GetName()), getRandSetter(getArrayBrackets(pGenner.Param().ArrayNestingLevel)+pGenner.Param().Type))
			}
		} else {
			if pGenner.Param().ArrayNestingLevel < 1 {
				fieldsGen += fmt.Sprintf("\t%s := new(models.%s)\n", getFullName(pGenner.GetName()), pGenner.Param().Type)
				fieldsGen += fmt.Sprintf("\tfillRandomly%s(%s)\n", pGenner.Param().Type, getFullName(pGenner.GetName()))
				fieldsGen += fmt.Sprintf("\tr.With%s(*%[1]s)\n", getFullName(pGenner.GetName()))
			} else {
				fieldsGen += fmt.Sprintf("\t%s := new(%smodels.%s)\n", getFullName(pGenner.GetName()), getArrayBrackets(pGenner.Param().ArrayNestingLevel), pGenner.Param().Type)
				fieldsGen += fmt.Sprintf("\tl%s := randIntn(maxArrayLength + 1)\n", getFullName(pGenner.GetName()))
				fieldsGen += fmt.Sprintf("\t*%s = make(%smodels.%s, l%[1]s)\n", getFullName(pGenner.GetName()), getArrayBrackets(pGenner.Param().ArrayNestingLevel), pGenner.Param().Type)
				fieldsGen += fmt.Sprintf("\tfor i0 := 0; i0 < l%[2]s; i0++ {\n\t\tfillRandomly%[1]s(&(*%s)[i0])\n\t}\n", pGenner.Param().Type, getFullName(pGenner.GetName()))
				fieldsGen += fmt.Sprintf("\tr.With%s(*%[1]s)\n", getFullName(pGenner.GetName()))
			}
		}
	}

	testGen += fmt.Sprintf("func %s%s(r *requests.%s) {\n%s}\n\n", testMethodName, m.RequestName, m.RequestName, fieldsGen)

	return
}
