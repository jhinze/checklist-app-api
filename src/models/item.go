package models

import (
	. "checklist/database"
	"github.com/jinzhu/gorm"
	"time"
)

type Item struct {
	Model
	ListID      uint       `json:"listId"`
	Description string     `json:"description",gorm:"not null"`
	DoneAt      *time.Time `json:"doneAt"`
}

func GetAllItemsOnList(list *List, items *[]Item) (db *gorm.DB) {
	return DB.Where("list_id = ?", list.ID).Find(items)
}

func GetItem(item *Item) (db *gorm.DB) {
	return DB.Find(item)
}

func GetAllItems(items *[]Item) (db *gorm.DB) {
	return DB.Where("deleted_at IS null").Find(items)
}

func CreateItem(item *Item) (db *gorm.DB) {
	return DB.Create(item)
}

func UpdateItem(item *Item) (db *gorm.DB) {
	return DB.Save(item)
}

func DeleteItem(item *Item) (db *gorm.DB) {
	return DB.Where("id = ? AND list_id = ?", item.ID, item.ListID).Delete(Item{})
}

func CompleteItem(item *Item, time time.Time) (db *gorm.DB) {
	return DB.Model(Item{}).Where("id = ? AND list_id = ?", item.ID, item.ListID).Update("doneAt", time)
}

func IncompleteItem(item *Item) (db *gorm.DB) {
	return DB.Model(Item{}).Where("id = ? AND list_id = ?", item.ID, item.ListID).Update("doneAt", nil)
}
