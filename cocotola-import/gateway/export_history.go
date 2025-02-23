package gateway

import (
	"errors"
	"time"

	"gorm.io/gorm"

	"github.com/kujilabo/cocotola-1.23/cocotola-import/service"
	rsliberrors "github.com/kujilabo/cocotola-1.23/redstart/lib/errors"
)

type exportHistoryEntity struct {
	ID         int
	WorkbookID int
	Status     string
	ExportedAt time.Time
}

func (u *exportHistoryEntity) TableName() string {
	return "export_history"
}

type exportHistoryRepository struct {
	db *gorm.DB
}

func NewExportHistoryRepository(db *gorm.DB) service.ExportHistoryRepository {
	return &exportHistoryRepository{
		db: db,
	}
}

func (r *exportHistoryRepository) GetLatestExportedAt() (time.Time, error) {
	exportHistorEntity := exportHistoryEntity{}
	if result := r.db.Where("workbook_id = ?", 1).First(&exportHistorEntity).Order("exported_at desc"); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC), nil
		} else {
			return time.Time{}, rsliberrors.Errorf("Select: %w", result.Error)
		}
	}

	return exportHistorEntity.ExportedAt, nil
}

func (r *exportHistoryRepository) Add(workbookID int, status string, exportedAt time.Time) error {
	if result := r.db.Create(&exportHistoryEntity{
		WorkbookID: workbookID,
		Status:     status,
		ExportedAt: exportedAt,
	}); result.Error != nil {
		return rsliberrors.Errorf("Create: %w", result.Error)
	}
	return nil
}
