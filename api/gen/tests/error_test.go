package tests

import "github.com/defany/govk/pkg/random"

func (e *apiError) fillRandomly() {
	subcode := random.RandInt()
	l := random.RandIntn(random.MaxArrayLength + 1)
	requestParams := make([]RequestParam, l)
	for i := 0; i < l; i++ {
		requestParams[i].fillRandomly()
	}
	redirectURI := random.RandString()
	confirmationText := random.RandString()
	captchaSID := random.RandString()
	captchaImg := random.RandString()

	*e = apiError{
		ErrorCode:    random.RandInt(),
		ErrorSubcode: &subcode,
		ErrorMsg:     random.RandString(),
		ErrorText:    random.RandString(),
		ReqParams:    requestParams,
		RedirURI:     &redirectURI,
		ConfirmText:  &confirmationText,
		CaptchaSID:   &captchaSID,
		CaptchaImg:   &captchaImg,
	}
}

func (p *RequestParam) fillRandomly() {
	p.Key = random.RandString()
	p.Value = random.RandString()
}
