package basemodel

import "gateway/proto/gcommon"

type ApiError struct {
	Code    string                 `json:"code"`
	Message string                 `json:"message"`
	Details []*gcommon.ErrorDetail `json:"details"`
}

func (e ApiError) Error() string {
	return e.Message
}

type ApiEmptyResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type ApiIDResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	ID      string `json:"id"`
}

type PageMetadata struct {
	Page          int32 `json:"page"`
	Size          int32 `json:"size"`
	TotalItems    int32 `json:"totalItems"`    // Num items in current page
	TotalElements int64 `json:"totalElements"` // Num items in all pages
	TotalPages    int32 `json:"totalPages"`
	HasNext       bool  `json:"hasNext"`
	HasPrevious   bool  `json:"hasPrevious"`
}

type ApiIdResponse struct {
	Code string `json:"code"`
	Id   string `json:"id"`
}

type ApiExistenceResponse struct {
	Code  string `json:"code"`
	Exist bool   `json:"exist"`
}

type ApiDuplicationCheckingResponse struct {
	Code       string `json:"code"`
	Duplicated bool   `json:"duplicated"`
}
