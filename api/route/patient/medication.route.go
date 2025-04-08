package patientroute

import (
	patientcontroller "gateway/api/controller/patient"
	"gateway/api/route/common"
)

type MedicationRoute struct {
	MedicationController *patientcontroller.MedicationController
}

func NewMedicationRoute(patientController *patientcontroller.MedicationController) *MedicationRoute {
	return &MedicationRoute{MedicationController: patientController}
}

func (r *MedicationRoute) Register(props *common.RouterProps) {
	router := props.App.Group(props.Prefix)
	router.Post("/upsert-medication", r.MedicationController.UpsertMedication)
	router.Get("/list-medication", r.MedicationController.ListMedication)
}

func (r *MedicationRoute) SubPath() string {
	return "/patient/medication"
}
