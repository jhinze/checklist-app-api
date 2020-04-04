package models

import (
	. "checklist/database"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

type ItemsSummary struct {
	ListID      *uint      `json:"listId,omitempty"`
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	Total       *uint      `json:"total,omitempty"`
	Incomplete  *uint      `json:"incomplete,omitempty"`
}

func GetItemSummary(listId uint, itemsSummary *ItemsSummary) (db *gorm.DB) {
	db = DB.Raw("SELECT COUNT(*) as total, COUNT(*) - COUNT(done_at) as incomplete FROM items WHERE list_id = ?", listId)
	countRow := db.Row()
	countRowErr := countRow.Scan(&itemsSummary.Total, &itemsSummary.Incomplete)
	if countRowErr != nil {
		log.Println(countRowErr)
		return nil
	}
	db = DB.Raw("SELECT updated_at FROM items WHERE list_id = ? ORDER BY updated_at DESC LIMIT 1", listId)
	lastUpdatedRow := db.Row()
	lastUpdatedErr := lastUpdatedRow.Scan(&itemsSummary.LastUpdated)
	if lastUpdatedErr != nil {
		log.Println(lastUpdatedErr)
		return nil
	}
	return db
}
