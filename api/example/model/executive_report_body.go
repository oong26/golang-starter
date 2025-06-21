package model

type ExecutiveReportBody struct {
	RawData		[]JobsiteReportBody `json:"raw_data"`
	AICommand	string `json:"ai_command"`
}