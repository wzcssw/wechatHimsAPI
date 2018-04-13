package model

import "time"

type PatientInfo struct {
	ID                uint64 `gorm:"primary_key"`
	ExamineFinishFlag int
	HospitalNo        string
	FilmNo            string
	PatientName       string
	PatientSex        int
	PatientPhone      string
	ProjectID         int
	ProjectName       string
	BodyDescription   string
	DiagnoseInfo      string
	State             int
	AppointmentAt     time.Time
	OrderCode         string
	CanPrint          int
	IsPrint           int
	EnPatientName     string
}
