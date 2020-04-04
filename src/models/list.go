package models

import (
	. "checklist/database"
	"github.com/jinzhu/gorm"
	"log"
)

type List struct {
	Model
	Title        string        `json:"title",gorm:"not null"`
	Category     string        `json:"category",gorm:"not null"`
	Items        []Item        `json:"items,omitempty"`
	ItemsSummary *ItemsSummary `json:"itemsSummary,omitempty",gorm:"-"`
}

func getAllListsWithSummary(list *[]List) (db *gorm.DB) {
	itemsSummaryQuery :=
		`SELECT * FROM
			(SELECT max(items.updated_at) as last_updated, items.list_id, count(*) as total, COUNT(*) - COUNT(done_at) as incomplete 
				FROM items 
				WHERE deleted_at IS null 
				GROUP BY list_id)
			t1
			FULL OUTER JOIN
			(SELECT * FROM lists where lists.deleted_at IS null)
			t2
			ON t1.list_id = t2.id`
	db = DB.Raw(itemsSummaryQuery)
	rows, err := db.Rows()
	if err != nil {
		log.Println(err)
		return db
	}
	for rows.Next() {
		var nextList List
		var itemsSummary ItemsSummary
		nextList.ItemsSummary = &itemsSummary
		err := rows.Scan(&itemsSummary.LastUpdated, &itemsSummary.ListID, &itemsSummary.Total, &itemsSummary.Incomplete,
			&nextList.ID, &nextList.CreatedAt, &nextList.UpdatedAt, &nextList.DeletedAt, &nextList.Title, &nextList.Category)
		if err != nil {
			log.Println(err)
		}
		if nextList.ItemsSummary.ListID == nil &&
			nextList.ItemsSummary.LastUpdated == nil &&
			nextList.ItemsSummary.Total == nil &&
			nextList.ItemsSummary.Incomplete == nil {
			nextList.ItemsSummary = nil
		}
		*list = append(*list, nextList)
	}
	if *list == nil {
		*list = []List{}
	}
	return db
}

func getAllListsWithSummaryAndItems(list *[]List) (db *gorm.DB) {
	var listCopy []List
	db = getAllListsWithSummary(&listCopy)
	var items []Item
	db = GetAllItems(&items)
	for i, li := range listCopy {
		for j, item := range items {
			if li.ID == item.ListID {
				listCopy[i].Items = append(listCopy[i].Items, items[j])
			}
		}
	}
	*list = listCopy
	if *list == nil {
		*list = []List{}
	}
	return db
}

func GetAllLists(list *[]List, withItems bool, withItemsSummary bool) (db *gorm.DB) {
	if withItems && withItemsSummary {
		return getAllListsWithSummaryAndItems(list)
	} else if withItems {
		return DB.Preload("Items").Find(list)
	} else if withItemsSummary {
		return getAllListsWithSummary(list)
	} else {
		return DB.Find(list)
	}
}

func GetListWithItems(list *List) (db *gorm.DB) {
	return DB.Preload("Items").Find(list)
}

func GetList(list *List) (db *gorm.DB) {
	return DB.Find(list)
}

func CreateList(list *List) (db *gorm.DB) {
	return DB.Create(list)
}

func UpdateList(list *List) (db *gorm.DB) {
	return DB.Save(list)
}

func DeleteList(list *List) (dbList *gorm.DB, dbItem *gorm.DB) {
	return DB.Delete(list), DB.Where("list_id = ?", list.ID).Delete(Item{})
}
