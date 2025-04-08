package patientroute

import (
	patientcontroller "gateway/api/controller/patient"
	"gateway/api/route/common"
)

type PatientRoute struct {
	PatientController *patientcontroller.PatientController
}

func NewPatientRoute(patientController *patientcontroller.PatientController) *PatientRoute {
	return &PatientRoute{PatientController: patientController}
}

func (r *PatientRoute) Register(props *common.RouterProps) {
	router := props.App.Group(props.Prefix)

	router.Get("/list-patient", r.PatientController.ListPatient)
	router.Post("/upsert-patient", r.PatientController.UpsertPatient)
	router.Get("/get-patient/:id", r.PatientController.GetPatient)
	router.Post("/medical", r.PatientController.UpsertMedicalRecord)
	router.Post("/medical/treatment", r.PatientController.UpsertMedicalTreatment)
	router.Post("/medical/surgery", r.PatientController.UpsertMedicalSurgery)
	router.Post("/medical/prescription", r.PatientController.UpsertMedicalPrescription)
	router.Get("/medical/history", r.PatientController.GetMedicalHistory)
	router.Get("/medical/history/:id", r.PatientController.GetMedicalHistoryDetail)
}

func (r *PatientRoute) SubPath() string {
	return "/patient"
}
