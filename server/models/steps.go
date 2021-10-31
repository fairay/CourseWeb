package models

import "api/recipes/repository"

type StepM struct {
	rep    repository.StepsRep
	models *Models
}

func NewStep(rep repository.StepsRep, models *Models) *StepM {
	return &StepM{rep, models}
}
