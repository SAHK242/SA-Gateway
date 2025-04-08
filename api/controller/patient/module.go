package patientcontroller

import (
	"go.uber.org/fx"
)

var Module = fx.Provide(
	NewPatientController,
	NewMedicationController,
)
