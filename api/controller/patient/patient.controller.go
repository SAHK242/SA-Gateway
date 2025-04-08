package patientcontroller

import (
	basemodel "gateway/api/model/base"
	patientmodel "gateway/api/model/patient"
	apiutil "gateway/api/util"
	authgrpcclient "gateway/grpc/client/auth"
	patientgrpcclient "gateway/grpc/client/patient"
	grpcutil "gateway/grpc/util"
	"gateway/proto/auth"
	"gateway/proto/gcommon"
	"gateway/proto/patient"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"strings"
)

type PatientController struct {
	*patientgrpcclient.PatientGrpcClient
	*authgrpcclient.EmployeeGrpcClient
	redis *redis.Client
}

func NewPatientController(patientgrpcclient *patientgrpcclient.PatientGrpcClient, redis *redis.Client, employeegrpcclient *authgrpcclient.EmployeeGrpcClient) *PatientController {
	return &PatientController{PatientGrpcClient: patientgrpcclient, redis: redis, EmployeeGrpcClient: employeegrpcclient}
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
// @Router      /patient/list-patient [get]
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

	userIds := make([]string, 0)

	userIdsMap := make(map[string]bool)

	for _, pat := range res.Patients {
		if _, ok := userIdsMap[pat.CreatedBy.Id]; !ok {
			userIds = append(userIds, pat.CreatedBy.Id)
			userIdsMap[pat.CreatedBy.Id] = true
		}
		if _, ok := userIdsMap[pat.UpdatedBy.Id]; !ok {
			userIds = append(userIds, pat.UpdatedBy.Id)
			userIdsMap[pat.UpdatedBy.Id] = true
		}
	}

	employeeIds := strings.Join(userIds, ",")

	users, err := c.EmployeeGrpcClient.ListEmployee(&auth.ListEmployeeRequest{
		EmployeeId: employeeIds,
		Pageable:   apiutil.AsPageable(ctx, "first_name,asc"),
	}, md)

	if err != nil {
		return apiutil.AsApiError(err)
	}

	if apiutil.HasGrpcError(users.Error) {
		return apiutil.AsApiError(err)
	}

	employeeMap := make(map[string]*auth.User)
	for _, emp := range users.Employees {
		employeeMap[emp.Id] = emp
	}

	for _, pat := range res.Patients {
		if emp, ok := employeeMap[pat.CreatedBy.Id]; ok {
			pat.CreatedBy.FirstName = emp.FirstName
			pat.CreatedBy.LastName = emp.LastName
			pat.CreatedBy.Gender = emp.Gender
			pat.CreatedBy.Code = emp.Code
			pat.CreatedBy.PhoneNumber = emp.PhoneNumber
		}
		if emp, ok := employeeMap[pat.UpdatedBy.Id]; ok {
			pat.UpdatedBy.FirstName = emp.FirstName
			pat.UpdatedBy.LastName = emp.LastName
			pat.UpdatedBy.Gender = emp.Gender
			pat.UpdatedBy.Code = emp.Code
			pat.UpdatedBy.PhoneNumber = emp.PhoneNumber
		}
	}

	return apiutil.AsServerResponse(ctx, &patientmodel.ListPatientResponse{
		Code: apiutil.AsSuccessCode(),
		Data: patientmodel.ListPatientResponseData{
			PageMetadata: apiutil.AsPageMetadata(res.PageMetadata, req.Pageable, res.Patients),
			Patients:     res.Patients,
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

// UpsertMedicalRecord
// @Summary     Upsert Medical Record
// @Description Upsert Medical Record
// @Tags        SA - Patient Management
// @Accept      json
// @Produce     json
// @Param       _ body   patient.UpsertMedicalRecordRequest true "Upsert medical record Request"
// @Success     200                          {object} basemodel.ApiEmptyResponse
// @Failure     400,401,403,500              {object} basemodel.ApiError
// @Router      /patient/medical [post]
// @Security    JWT
func (c *PatientController) UpsertMedicalRecord(ctx *fiber.Ctx) error {
	md := grpcutil.WithModuleACLMetadata(ctx, gcommon.Module_MODULE_AUTH)

	req := new(patient.UpsertMedicalRecordRequest)

	if err := ctx.BodyParser(req); err != nil {
		return apiutil.AsApiError(err)
	}

	res, err := c.PatientGrpcClient.UpsertMedicalRecord(req, md)

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

// UpsertMedicalTreatment
// @Summary     Upsert Medical Treatment
// @Description Upsert Medical Treatment
// @Tags        SA - Patient Management
// @Accept      json
// @Produce     json
// @Param       _ body   patient.UpsertMedicalTreatmentRequest true "Upsert medical treatment Request"
// @Success     200                          {object} basemodel.ApiEmptyResponse
// @Failure     400,401,403,500              {object} basemodel.ApiError
// @Router      /patient/medical/treatment [post]
// @Security    JWT
func (c *PatientController) UpsertMedicalTreatment(ctx *fiber.Ctx) error {
	md := grpcutil.WithModuleACLMetadata(ctx, gcommon.Module_MODULE_AUTH)

	req := new(patient.UpsertMedicalTreatmentRequest)

	if err := ctx.BodyParser(req); err != nil {
		return apiutil.AsApiError(err)
	}

	res, err := c.PatientGrpcClient.UpsertMedicalTreatment(req, md)

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

// UpsertMedicalSurgery
// @Summary     Upsert Medical Surgery
// @Description Upsert Medical Surgery
// @Tags        SA - Patient Management
// @Accept      json
// @Produce     json
// @Param       _ body   patient.UpsertMedicalSurgeryRequest true "Upsert medical surgery Request"
// @Success     200                          {object} basemodel.ApiEmptyResponse
// @Failure     400,401,403,500              {object} basemodel.ApiError
// @Router      /patient/medical/surgery [post]
// @Security    JWT
func (c *PatientController) UpsertMedicalSurgery(ctx *fiber.Ctx) error {
	md := grpcutil.WithModuleACLMetadata(ctx, gcommon.Module_MODULE_AUTH)

	req := new(patient.UpsertMedicalSurgeryRequest)

	if err := ctx.BodyParser(req); err != nil {
		return apiutil.AsApiError(err)
	}

	res, err := c.PatientGrpcClient.UpsertMedicalSurgery(req, md)

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

// UpsertMedicalPrescription
// @Summary     Upsert Medical Prescription
// @Description Upsert Medical Prescription
// @Tags        SA - Patient Management
// @Accept      json
// @Produce     json
// @Param       _ body   patient.UpsertMedicalPrescriptionRequest true "Upsert medical prescription Request"
// @Success     200                          {object} basemodel.ApiEmptyResponse
// @Failure     400,401,403,500              {object} basemodel.ApiError
// @Router      /patient/medical/prescription [post]
// @Security    JWT
func (c *PatientController) UpsertMedicalPrescription(ctx *fiber.Ctx) error {
	md := grpcutil.WithModuleACLMetadata(ctx, gcommon.Module_MODULE_AUTH)

	req := new(patient.UpsertMedicalPrescriptionRequest)

	if err := ctx.BodyParser(req); err != nil {
		return apiutil.AsApiError(err)
	}

	res, err := c.PatientGrpcClient.UpsertMedicalPrescription(req, md)

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

// GetMedicalHistory
// @Summary     Get Medical History
// @Description Get Medical History
// @Tags        SA - Patient Management
// @Accept      json
// @Produce     json
// @Param       page            query    int    false "Page number"
// @Param       size            query    int    false "Page size"
// @Param       paging_ignored  query    bool   false "Ignore pagination and return all data"
// @Param       patient_id      query    string true "Patient ID"
// @Param       from_date        query    string false "From date"
// @Param       to_date        query    string false "To date"
// @Param       created_by      query    string false "Created by, doctor ID or nurse ID"
// @Success     200             {object} patientmodel.GetMedicalHistoryResponse
// @Failure     400,401,403,500 {object} basemodel.ApiError
// @Router      /patient/medical/history [get]
// @Security    JWT
func (c *PatientController) GetMedicalHistory(ctx *fiber.Ctx) error {
	md := grpcutil.WithModuleACLMetadata(ctx, gcommon.Module_MODULE_AUTH)

	req := &patient.GetMedicalHistoryRequest{
		Pageable:  apiutil.AsPageable(ctx, "created_at,desc"),
		PatientId: ctx.Query("patient_id", ""),
		CreatedBy: ctx.Query("created_by", ""),
		DateRange: &gcommon.DateRange{
			FromDate: int64(ctx.QueryInt("from_date", 0)),
			ToDate:   int64(ctx.QueryInt("to_date", 0)),
		},
	}

	res, err := c.PatientGrpcClient.GetMedicalHistory(req, md)

	if err != nil {
		return apiutil.AsApiError(err)
	}

	if apiutil.HasGrpcError(res.Error) {
		return apiutil.AsApiError(err)
	}

	userIds := make([]string, 0)
	userIdsMap := make(map[string]bool)
	for _, pat := range res.MedicalHistories {
		if _, ok := userIdsMap[pat.CreatedBy.Id]; !ok {
			userIds = append(userIds, pat.CreatedBy.Id)
			userIdsMap[pat.CreatedBy.Id] = true
		}
		if _, ok := userIdsMap[pat.UpdatedBy.Id]; !ok {
			userIds = append(userIds, pat.UpdatedBy.Id)
			userIdsMap[pat.UpdatedBy.Id] = true
		}
	}

	employeeIds := strings.Join(userIds, ",")
	users, err := c.EmployeeGrpcClient.ListEmployee(&auth.ListEmployeeRequest{
		EmployeeId: employeeIds,
		Pageable:   apiutil.AsPageable(ctx, "first_name,asc"),
	}, md)

	if err != nil {
		return apiutil.AsApiError(err)
	}

	if apiutil.HasGrpcError(users.Error) {
		return apiutil.AsApiError(err)
	}

	employeeMap := make(map[string]*auth.User)
	for _, emp := range users.Employees {
		employeeMap[emp.Id] = emp
	}

	for _, pat := range res.MedicalHistories {
		if emp, ok := employeeMap[pat.CreatedBy.Id]; ok {
			pat.CreatedBy.FirstName = emp.FirstName
			pat.CreatedBy.LastName = emp.LastName
			pat.CreatedBy.Gender = emp.Gender
			pat.CreatedBy.Code = emp.Code
			pat.CreatedBy.PhoneNumber = emp.PhoneNumber
		}
		if emp, ok := employeeMap[pat.UpdatedBy.Id]; ok {
			pat.UpdatedBy.FirstName = emp.FirstName
			pat.UpdatedBy.LastName = emp.LastName
			pat.UpdatedBy.Gender = emp.Gender
			pat.UpdatedBy.Code = emp.Code
			pat.UpdatedBy.PhoneNumber = emp.PhoneNumber
		}
	}

	return apiutil.AsServerResponse(ctx, &patientmodel.GetMedicalHistoryResponse{
		Code: apiutil.AsSuccessCode(),
		MedicalHistory: &patientmodel.MedicalHistoryData{
			PageMetadata:     apiutil.AsPageMetadata(res.PageMetadata, req.Pageable, res.MedicalHistories),
			MedicalHistories: res.MedicalHistories,
		},
	})
}

// GetMedicalHistoryDetail
// @Summary     Get Medical History Detail
// @Description Get Medical History Detail
// @Tags        SA - Patient Management
// @Accept      json
// @Produce     json
// @Param       id path   string true "Medical History ID"
// @Success     200                          {object} patientmodel.GetMedicalHistoryDetailResponse
// @Failure     400,401,403,500              {object} basemodel.ApiError
// @Router      /patient/medical/history/{id} [get]
// @Security    JWT
func (c *PatientController) GetMedicalHistoryDetail(ctx *fiber.Ctx) error {
	md := grpcutil.WithModuleACLMetadata(ctx, gcommon.Module_MODULE_AUTH)

	req := &gcommon.IdRequest{
		Id: ctx.Params("id"),
	}

	res, err := c.PatientGrpcClient.GetMedicalHistoryDetail(req, md)

	if err != nil {
		return apiutil.AsApiError(err)
	}

	if apiutil.HasGrpcError(res.Error) {
		return apiutil.AsApiError(err)
	}

	userIds := make([]string, 0)
	userIdsMap := make(map[string]bool)
	pat := res.MedicalHistory
	if _, ok := userIdsMap[pat.CreatedBy.Id]; !ok {
		userIds = append(userIds, pat.CreatedBy.Id)
		userIdsMap[pat.CreatedBy.Id] = true
	}
	if _, ok := userIdsMap[pat.UpdatedBy.Id]; !ok {
		userIds = append(userIds, pat.UpdatedBy.Id)
		userIdsMap[pat.UpdatedBy.Id] = true
	}

	for _, pat2 := range res.MedicalTreatments {
		if _, ok := userIdsMap[pat2.CreatedBy.Id]; !ok {
			userIds = append(userIds, pat2.CreatedBy.Id)
			userIdsMap[pat2.CreatedBy.Id] = true
		}
		if _, ok := userIdsMap[pat2.UpdatedBy.Id]; !ok {
			userIds = append(userIds, pat2.UpdatedBy.Id)
			userIdsMap[pat2.UpdatedBy.Id] = true
		}
	}

	for _, pat2 := range res.MedicalSurgeries {
		if _, ok := userIdsMap[pat2.CreatedBy.Id]; !ok {
			userIds = append(userIds, pat2.CreatedBy.Id)
			userIdsMap[pat2.CreatedBy.Id] = true
		}
		if _, ok := userIdsMap[pat2.UpdatedBy.Id]; !ok {
			userIds = append(userIds, pat2.UpdatedBy.Id)
			userIdsMap[pat2.UpdatedBy.Id] = true
		}
	}

	for _, pat2 := range res.MedicalPrescriptions {
		if _, ok := userIdsMap[pat2.CreatedBy.Id]; !ok {
			userIds = append(userIds, pat2.CreatedBy.Id)
			userIdsMap[pat2.CreatedBy.Id] = true
		}
	}

	employeeIds := strings.Join(userIds, ",")
	users, err := c.EmployeeGrpcClient.ListEmployee(&auth.ListEmployeeRequest{
		EmployeeId: employeeIds,
		Pageable:   apiutil.AsPageable(ctx, "first_name,asc"),
	}, md)

	if err != nil {
		return apiutil.AsApiError(err)
	}

	if apiutil.HasGrpcError(users.Error) {
		return apiutil.AsApiError(err)
	}

	employeeMap := make(map[string]*auth.User)
	for _, emp := range users.Employees {
		employeeMap[emp.Id] = emp
	}

	if emp, ok := employeeMap[pat.CreatedBy.Id]; ok {
		pat.CreatedBy.FirstName = emp.FirstName
		pat.CreatedBy.LastName = emp.LastName
		pat.CreatedBy.Gender = emp.Gender
		pat.CreatedBy.Code = emp.Code
		pat.CreatedBy.PhoneNumber = emp.PhoneNumber
	}

	if emp, ok := employeeMap[pat.UpdatedBy.Id]; ok {
		pat.UpdatedBy.FirstName = emp.FirstName
		pat.UpdatedBy.LastName = emp.LastName
		pat.UpdatedBy.Gender = emp.Gender
		pat.UpdatedBy.Code = emp.Code
		pat.UpdatedBy.PhoneNumber = emp.PhoneNumber
	}

	for _, pat2 := range res.MedicalTreatments {
		if emp, ok := employeeMap[pat2.CreatedBy.Id]; ok {
			pat2.CreatedBy.FirstName = emp.FirstName
			pat2.CreatedBy.LastName = emp.LastName
			pat2.CreatedBy.Gender = emp.Gender
			pat2.CreatedBy.Code = emp.Code
			pat2.CreatedBy.PhoneNumber = emp.PhoneNumber
		}
		if emp, ok := employeeMap[pat2.UpdatedBy.Id]; ok {
			pat2.UpdatedBy.FirstName = emp.FirstName
			pat2.UpdatedBy.LastName = emp.LastName
			pat2.UpdatedBy.Gender = emp.Gender
			pat2.UpdatedBy.Code = emp.Code
			pat2.UpdatedBy.PhoneNumber = emp.PhoneNumber
		}
	}

	for _, pat2 := range res.MedicalSurgeries {
		if emp, ok := employeeMap[pat2.CreatedBy.Id]; ok {
			pat2.CreatedBy.FirstName = emp.FirstName
			pat2.CreatedBy.LastName = emp.LastName
			pat2.CreatedBy.Gender = emp.Gender
			pat2.CreatedBy.Code = emp.Code
			pat2.CreatedBy.PhoneNumber = emp.PhoneNumber
		}
		if emp, ok := employeeMap[pat2.UpdatedBy.Id]; ok {
			pat2.UpdatedBy.FirstName = emp.FirstName
			pat2.UpdatedBy.LastName = emp.LastName
			pat2.UpdatedBy.Gender = emp.Gender
			pat2.UpdatedBy.Code = emp.Code
			pat2.UpdatedBy.PhoneNumber = emp.PhoneNumber
		}
	}

	for _, pat2 := range res.MedicalPrescriptions {
		if emp, ok := employeeMap[pat2.CreatedBy.Id]; ok {
			pat2.CreatedBy.FirstName = emp.FirstName
			pat2.CreatedBy.LastName = emp.LastName
			pat2.CreatedBy.Gender = emp.Gender
			pat2.CreatedBy.Code = emp.Code
			pat2.CreatedBy.PhoneNumber = emp.PhoneNumber
		}
	}

	return apiutil.AsServerResponse(ctx, &patientmodel.GetMedicalHistoryDetailResponse{
		Code: apiutil.AsSuccessCode(),
		Data: &patientmodel.GetMedicalHistoryDetailResponseData{
			MedicalHistory:       res.MedicalHistory,
			MedicalTreatments:    res.MedicalTreatments,
			MedicalSurgeries:     res.MedicalSurgeries,
			MedicalPrescriptions: res.MedicalPrescriptions,
		},
	})
}
