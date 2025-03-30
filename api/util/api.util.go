package apiutil

import (
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc/status"

	"gateway/api/model/base"
	"gateway/proto/gcommon"
)

func AsPageable(ctx *fiber.Ctx, defaultSort string) *gcommon.Pageable {
	page := StringToInt32(ctx.Query("page"), 0)
	size := StringToInt32(ctx.Query("size"), 10)
	sort := ctx.Query("sort", defaultSort)
	pagingIgnored := StringToBool(ctx.Query("paging_ignored"), false)

	pageable := &gcommon.Pageable{
		Page:          page,
		Size:          size,
		Sort:          sort,
		PagingIgnored: pagingIgnored,
	}

	return pageable
}

func AsDefaultPageable() *gcommon.Pageable {
	return &gcommon.Pageable{
		Page:          0,
		Size:          10,
		PagingIgnored: false,
	}
}

func AsNonPagingPageable(sort string) *gcommon.Pageable {
	return &gcommon.Pageable{
		Sort:          sort,
		PagingIgnored: true,
	}
}

func AsApiError(err error) *basemodel.ApiError {
	if e, ok := status.FromError(err); ok {
		return &basemodel.ApiError{
			Code:    gcommon.Code_CODE_UNKNOWN_ERROR.String(),
			Message: e.Message(),
		}
	} else {
		return &basemodel.ApiError{
			Code:    gcommon.Code_CODE_UNKNOWN_ERROR.String(),
			Message: err.Error(),
		}
	}
}

func HasGrpcError(err *gcommon.Error) bool {
	return err != nil && err.Code != gcommon.Code_CODE_SUCCESS
}

func AsGrpcError(err *gcommon.Error) *basemodel.ApiError {
	return &basemodel.ApiError{
		Code:    AsResponseCode(err),
		Message: AsResponseMessage(err),
		Details: err.Details,
	}
}

func AsForbiddenError() *basemodel.ApiError {
	return &basemodel.ApiError{
		Code:    gcommon.Code_CODE_FORBIDDEN.String(),
		Message: "Permission Denied",
	}
}

func AsServerResponse(ctx *fiber.Ctx, res interface{}) error {
	return AsServerResponseWithCode(ctx, fiber.StatusOK, res)
}

func AsServerResponseWithCode(ctx *fiber.Ctx, code int, res interface{}) error {
	return ctx.Status(code).JSON(res)
}

func AsEmptyResponse(ctx *fiber.Ctx, grpcError *gcommon.Error, err error) error {
	if err != nil {
		return AsApiError(err)
	}

	if HasGrpcError(grpcError) {
		return AsGrpcError(grpcError)
	}

	return ctx.JSON(&basemodel.ApiEmptyResponse{
		Code:    AsSuccessCode(),
		Message: "OK",
	})
}

func AsIDResponse(ctx *fiber.Ctx, grpcError *gcommon.Error, err error, id string) error {
	if err != nil {
		return AsApiError(err)
	}

	if HasGrpcError(grpcError) {
		return AsGrpcError(grpcError)
	}

	return ctx.JSON(&basemodel.ApiIDResponse{
		Code:    AsSuccessCode(),
		Message: "OK",
		ID:      id,
	})
}

func asListMetadata[T any](data []T) *basemodel.PageMetadata {
	return &basemodel.PageMetadata{
		Page:          1,
		Size:          int32(len(data)),
		TotalItems:    int32(len(data)),
		TotalElements: int64(len(data)),
		TotalPages:    1,
		HasNext:       false,
		HasPrevious:   false,
	}
}

func AsPageMetadata[T any](metadata *gcommon.PageMetadata, pageable *gcommon.Pageable, data []T) *basemodel.PageMetadata {
	if pageable.PagingIgnored || metadata == nil {
		return asListMetadata(data)
	}

	return &basemodel.PageMetadata{
		Page:          metadata.Page,
		Size:          metadata.Size,
		TotalItems:    int32(len(data)),
		TotalElements: metadata.TotalElements,
		TotalPages:    metadata.TotalPages,
		HasNext:       metadata.HasNext,
		HasPrevious:   metadata.HasPrevious,
	}
}

func AsEmptyRestPageMetadata() *basemodel.PageMetadata {
	return &basemodel.PageMetadata{
		Page:          0,
		Size:          0,
		TotalItems:    0,
		TotalElements: 0,
		TotalPages:    0,
		HasNext:       false,
		HasPrevious:   false,
	}
}

func AsEmptyGrpcPageMetadata() *gcommon.PageMetadata {
	return &gcommon.PageMetadata{
		Page:          0,
		Size:          0,
		TotalElements: 0,
		TotalPages:    0,
		HasNext:       false,
		HasPrevious:   false,
	}
}

func AsResponseCode(error *gcommon.Error) string {
	if error == nil {
		return gcommon.Code_CODE_UNSPECIFIED.String()
	}

	return error.Code.String()
}

func AsSuccessCode() string {
	return gcommon.Code_CODE_SUCCESS.String()
}

func AsResponseMessage(error *gcommon.Error) string {
	if error.Message == "" {
		return "unexpected error"
	} else {
		return error.Message
	}
}
