package service

import "time"

type ExportHistoryRepository interface {
	GetLatestExportedAt() (time.Time, error)
	Add(workbookID int, status string, exportedAt time.Time) error
}
