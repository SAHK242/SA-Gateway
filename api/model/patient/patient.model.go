package patientmodel

import (
	basemodel "gateway/api/model/base"
	"gateway/proto/patient"
)

type GetPatientResponse struct {
	Code          string                 `json:"code"`
	PatientDetail *patient.PatientDetail `json:"patientDetail"`
}

type ListPatientResponse struct {
	Code string                  `json:"code"`
	Data ListPatientResponseData `json:"data"`
}

type ListPatientResponseData struct {
	*basemodel.PageMetadata
	Patients []*patient.PatientDetail `json:"patients"`
}
