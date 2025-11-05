package entity

type SummaryProject struct {
	ProjectID      int     `json:"project_id" gorm:"not null"`
	TotalTasks     int     `json:"total_tasks" gorm:"not null"`
	CompletedTasks int     `json:"completed_tasks" gorm:"not null"`
	Progress       float64 `json:"progress" gorm:"not null"`
}
