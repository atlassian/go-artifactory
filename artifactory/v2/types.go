package v2

import (
	"github.com/go-resty/resty/v2"
)

type Service struct {
	client *resty.Client
}

type V2 struct {
	common Service

	Security *SecurityService
}
