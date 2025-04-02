package patientcontroller

import (
	basemodel "gateway/api/model/base"
	patientmodel "gateway/api/model/patient"
	apiutil "gateway/api/util"
	patientgrpcclient "gateway/grpc/client/patient"
	grpcutil "gateway/grpc/util"
	"gateway/proto/gcommon"
	"gateway/proto/patient"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

type PatientController struct {
	*patientgrpcclient.PatientGrpcClient
	redis *redis.Client
}

func NewPatientController(patientgrpcclient *patientgrpcclient.PatientGrpcClient, redis *redis.Client) *PatientController {
	return &PatientController{PatientGrpcClient: patientgrpcclient, redis: redis}
}

// UpsertPatient
// @Summary     Upsert Patient
// @Description Upsert Patient
// @Tags        SA - Patient Management
// @Accept      json
// @Produce     json
// @Param       _ body   patient.UpsertPatientRequest true "Upsert patient Request"
// @Success     200                          {object} basemodel.ApiEmptyResponse
// @Failure     400,401,403,500              {object} basemodel.ApiError
// @Router      /patient/upsert-patient [post]
// @Security    JWT
func (c *PatientController) UpsertPatient(ctx *fiber.Ctx) error {
	md := grpcutil.WithModuleACLMetadata(ctx, gcommon.Module_MODULE_AUTH)

	req := new(patient.UpsertPatientRequest)

	if err := ctx.BodyParser(req); err != nil {
		return apiutil.AsApiError(err)
	}

	res, err := c.PatientGrpcClient.UpsertPatient(req, md)

	if err != nil {
		return apiutil.AsApiError(err)
	}

	if apiutil.HasGrpcError(res.Error) {
		return apiutil.AsApiError(err)
	}

	return apiutil.AsServerResponse(ctx, &basemodel.ApiEmptyResponse{
		Code:    apiutil.AsSuccessCode(),
		Message: "OK",
	})
}

// GetPatient
// @Summary     Get Patient
// @Description Get Patient
// @Tags        SA - Patient Management
// @Accept      json
// @Produce     json
// @Param       id path   string true "Patient ID"
// @Success     200                          {object} patientmodel.GetPatientResponse
// @Failure     400,401,403,500              {object} basemodel.ApiError
// @Router      /patient/get-patient/{id} [get]
// @Security    JWT
func (c *PatientController) GetPatient(ctx *fiber.Ctx) error {
	md := grpcutil.WithModuleACLMetadata(ctx, gcommon.Module_MODULE_AUTH)

	req := &gcommon.IdRequest{
		Id: ctx.Params("id"),
	}

	res, err := c.PatientGrpcClient.GetPatient(req, md)

	if err != nil {
		return apiutil.AsApiError(err)
	}

	if apiutil.HasGrpcError(res.Error) {
		return apiutil.AsApiError(err)
	}

	return apiutil.AsServerResponse(ctx, res)
}

// ListPatient
// @Summary     List Employee
// @Description List Employee
// @Tags        SA - Patient Management
// @Accept      json
// @Produce     json
// @Param       page            query    int    false "Page number"
// @Param       size            query    int    false "Page size"
// @Param       sort            query    string false "Sort field. Default to name,asc" Enums(name)
// @Param       paging_ignored  query    bool   false "Ignore pagination and return all data"
// @Param       search          query    string false "Search by name"
// @Success     200             {object} patientmodel.ListPatientResponse
// @Failure     400,401,403,500 {object} basemodel.ApiError
// @Router      /auth/list-patient [get]
// @Security    JWT
func (c *PatientController) ListPatient(ctx *fiber.Ctx) error {
	md := grpcutil.WithModuleACLMetadata(ctx, gcommon.Module_MODULE_AUTH)

	req := c.buildListRequest(ctx)

	res, err := c.PatientGrpcClient.ListPatient(req, md)

	if err != nil {
		return apiutil.AsApiError(err)
	}

	if apiutil.HasGrpcError(res.Error) {
		return apiutil.AsApiError(err)
	}

	return apiutil.AsServerResponse(ctx, &patientmodel.ListPatientResponse{
		Code: apiutil.AsSuccessCode(),
		Data: patientmodel.ListPatientResponseData{
			PageMetadata: apiutil.AsPageMetadata(res.PageMetadata, req.Pageable, res.PatientDetails),
			Patients:     res.PatientDetails,
		},
	})

}

func (c *PatientController) buildListRequest(ctx *fiber.Ctx) *patient.ListPatientRequest {
	req := &patient.ListPatientRequest{
		Pageable: apiutil.AsPageable(ctx, "first_name,asc"),
		Search:   ctx.Query("search"),
	}
	return req
}
