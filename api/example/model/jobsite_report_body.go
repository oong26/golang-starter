package model

type JobsiteReportBody struct {
	ID						int `json:"id"`
	StakeholderTotal		int `json:"stakeholder_total"`
	CommunicationTotal		int `json:"communication_total"`
	SiteTension				string `json:"site_tension"`
}