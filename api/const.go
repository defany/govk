package api

const UserAgent = "govk"

const (
	retries = 3
)

const (
	Version = "5.199"
	Url     = "https://api.vk.com/method/"
)

const (
	maxGroupApiRequestsCalls = 25
)

const (
	// TODO: add if fasthttp will support ctx in req
	ctxParam = ":ctx"

	versionParam = "v"

	tokenParam = "access_token"

	langParam = "lang"
)

// encodings
const (
	gzip = "gzip"
)
