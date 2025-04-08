package patientmodel

import (
	basemodel "gateway/api/model/base"
	"gateway/proto/patient"
)

type GetPatientResponse struct {
	Code    string           `json:"code"`
	Patient *patient.Patient `json:"patient"`
}

type ListPatientResponse struct {
	Code string                  `json:"code"`
	Data ListPatientResponseData `json:"data"`
}

type ListPatientResponseData struct {
	*basemodel.PageMetadata
	Patients []*patient.Patient `json:"patients"`
}

type GetMedicalHistoryResponse struct {
	Code           string              `json:"code"`
	MedicalHistory *MedicalHistoryData `json:"medical_history"`
}

type MedicalHistoryData struct {
	*basemodel.PageMetadata
	MedicalHistories []*patient.MedicalHistory `json:"medical_histories"`
}

type ListMedicationResponse struct {
	Code string                     `json:"code"`
	Data ListMedicationResponseData `json:"data"`
}

type ListMedicationResponseData struct {
	*basemodel.PageMetadata
	Medications []*patient.Medication `json:"medications"`
}

type GetMedicalHistoryDetailResponse struct {
	Code string                               `json:"code"`
	Data *GetMedicalHistoryDetailResponseData `json:"data"`
}

type GetMedicalHistoryDetailResponseData struct {
	MedicalHistory       *patient.MedicalHistory        `json:"medical_history"`
	MedicalTreatments    []*patient.MedicalTreatment    `json:"medical_treatments"`
	MedicalSurgeries     []*patient.MedicalSurgery      `json:"medical_surgeries"`
	MedicalPrescriptions []*patient.MedicalPrescription `json:"medical_prescriptions"`
}
