package patientcontroller

import (
	basemodel "gateway/api/model/base"
	patientmodel "gateway/api/model/patient"
	apiutil "gateway/api/util"
	authgrpcclient "gateway/grpc/client/auth"
	patientgrpcclient "gateway/grpc/client/patient"
	grpcutil "gateway/grpc/util"
	"gateway/proto/gcommon"
	"gateway/proto/patient"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

type MedicationController struct {
	*patientgrpcclient.MedicationGrpcClient
	*authgrpcclient.AuthGrpcClient
	redis *redis.Client
}

func NewMedicationController(medicationgrpcclient *patientgrpcclient.MedicationGrpcClient, redis *redis.Client, authgrpcclient *authgrpcclient.AuthGrpcClient) *MedicationController {
	return &MedicationController{MedicationGrpcClient: medicationgrpcclient, redis: redis, AuthGrpcClient: authgrpcclient}
}

// UpsertMedication
// @Summary     Upsert Medication
// @Description Upsert Medication
// @Tags        SA - Medication Management
// @Accept      json
// @Produce     json
// @Param       _ body   patient.UpsertMedicationRequest true "Upsert medication Request"
// @Success     200                          {object} basemodel.ApiEmptyResponse
// @Failure     400,401,403,500              {object} basemodel.ApiError
// @Router      /patient/medication/upsert-medication [post]
// @Security    JWT
func (c *MedicationController) UpsertMedication(ctx *fiber.Ctx) error {
	md := grpcutil.WithModuleACLMetadata(ctx, gcommon.Module_MODULE_AUTH)

	req := new(patient.UpsertMedicationRequest)

	if err := ctx.BodyParser(req); err != nil {
		return apiutil.AsApiError(err)
	}

	res, err := c.MedicationGrpcClient.UpsertMedication(req, md)

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

// ListMedication
// @Summary     List Medication
// @Description List Medication
// @Tags        SA - Medication Management
// @Accept      json
// @Produce     json
// @Param       page            query    int    false "Page number"
// @Param       size            query    int    false "Page size"
// @Param       sort            query    string false "Sort field. Default to name,asc" Enums(name, created_at)
// @Param       paging_ignored  query    bool   false "Ignore pagination and return all data"
// @Param       search          query    string false "Search by name"
// @Success     200             {object} patientmodel.ListMedicationResponse
// @Failure     400,401,403,500 {object} basemodel.ApiError
// @Router      /patient/medication/list-medication [get]
// @Security    JWT
func (c *MedicationController) ListMedication(ctx *fiber.Ctx) error {
	md := grpcutil.WithModuleACLMetadata(ctx, gcommon.Module_MODULE_AUTH)

	req := &patient.ListMedicationRequest{
		Pageable: apiutil.AsPageable(ctx, "name,asc"),
		Search:   ctx.Query("search"),
	}

	res, err := c.MedicationGrpcClient.ListMedication(req, md)

	if err != nil {
		return apiutil.AsApiError(err)
	}

	if apiutil.HasGrpcError(res.Error) {
		return apiutil.AsApiError(err)
	}

	return apiutil.AsServerResponse(ctx, &patientmodel.ListMedicationResponse{
		Code: apiutil.AsSuccessCode(),
		Data: patientmodel.ListMedicationResponseData{
			PageMetadata: apiutil.AsPageMetadata(res.PageMetadata, req.Pageable, res.Medications),
			Medications:  res.Medications,
		},
	})
}
