package authmodel

import (
	basemodel "gateway/api/model/base"
	"gateway/proto/auth"
)

type LoginResponse struct {
	Code     string        `json:"code"`
	Token    string        `json:"token"`
	User     *auth.User    `json:"user"`
	NextStep auth.NextStep `json:"next_step"`
}

type GenerateLoginSecretRequest struct {
	Data string `json:"data"`
}

type GenerateLoginSecretResponse struct {
	Code string      `json:"code"`
	Data LoginSecret `json:"data"`
}

type LoginSecret struct {
	Secret string `json:"secret"`
}

type ListEmployee struct {
	Code string    `json:"code"`
	Data Employees `json:"data"`
}

type Employees struct {
	*basemodel.PageMetadata
	Employees []*auth.User `json:"employees"`
}

type ListDepartment struct {
	Code string      `json:"code"`
	Data Departments `json:"data"`
}

type Departments struct {
	*basemodel.PageMetadata
	Departments []*auth.Department `json:"departments"`
}
