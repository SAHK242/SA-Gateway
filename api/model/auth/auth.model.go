package authmodel

import (
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
