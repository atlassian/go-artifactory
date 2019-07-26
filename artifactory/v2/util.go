package v2

import "github.com/go-resty/resty/v2"

func String(v string) *string { return &v }

func NewV2(client *resty.Client) *V2 {
	v := &V2{}
	v.common.client = client

	v.Security = (*SecurityService)(&v.common)

	return v
}
