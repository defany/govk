package tests

import "github.com/defany/govk/pkg/random"

func (e *apiError) fillRandomly() {
	subcode := random.Int()
	l := random.IntDiapason(random.MaxArrayLength + 1)
	requestParams := make([]RequestParam, l)
	for i := 0; i < l; i++ {
		requestParams[i].fillRandomly()
	}
	redirectURI := random.String()
	confirmationText := random.String()
	captchaSID := random.String()
	captchaImg := random.String()

	*e = apiError{
		ErrorCode:    random.Int(),
		ErrorSubcode: &subcode,
		ErrorMsg:     random.String(),
		ErrorText:    random.String(),
		ReqParams:    requestParams,
		RedirURI:     &redirectURI,
		ConfirmText:  &confirmationText,
		CaptchaSID:   &captchaSID,
		CaptchaImg:   &captchaImg,
	}
}

func (p *RequestParam) fillRandomly() {
	p.Key = random.String()
	p.Value = random.String()
}
