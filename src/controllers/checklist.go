package controllers

import (
	"checklist/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
	"time"
)

func RetrieveItemsSummary(c *gin.Context) {
	listId, parseError := strconv.ParseUint(c.Param("listId"), 10, 32)
	if parseError != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	var itemsSummary models.ItemsSummary
	id := uint(listId)
	db := models.GetItemSummary(id, &itemsSummary)
	if db == nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	} else if db.Error != nil {
		c.AbortWithError(http.StatusInternalServerError, db.Error)
	} else {
		c.JSON(http.StatusOK, itemsSummary)
	}
}

func CreateList(c *gin.Context) {
	var list models.List
	bindError := c.BindJSON(&list)
	if bindError != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	db := models.CreateList(&list)
	if db.Error != nil {
		c.AbortWithError(http.StatusInternalServerError, db.Error)
	} else {
		c.JSON(http.StatusOK, list)
	}
}

func RetrieveLists(c *gin.Context) {
	var lists []models.List
	withItems := c.DefaultQuery("withItems", "false")
	withItemsSummary := c.DefaultQuery("withItemsSummary", "false")
	db := models.GetAllLists(&lists, "true" == withItems, "true" == withItemsSummary)
	if db.Error != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, lists)
	}
}

func RetrieveList(c *gin.Context) {
	var list models.List
	var db *gorm.DB
	id, parseError := strconv.ParseUint(c.Param("listId"), 10, 32)
	if parseError != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	list.ID = uint(id)
	withItems := c.DefaultQuery("withItems", "false")
	if "true" == withItems {
		db = models.GetListWithItems(&list)
	} else {
		db = models.GetList(&list)
	}
	if db.Error != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, list)
	}
}

func UpdateList(c *gin.Context) {
	var list models.List
	bindError := c.BindJSON(&list)
	if bindError != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	db := models.UpdateList(&list)
	if db.Error != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, list)
	}
}

func DeleteList(c *gin.Context) {
	var list models.List
	id, parseError := strconv.ParseUint(c.Param("listId"), 10, 32)
	if parseError != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	list.ID = uint(id)
	listDb, itemDb := models.DeleteList(&list)
	if listDb.Error != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"rowsAffected": listDb.RowsAffected + itemDb.RowsAffected,
		})
	}
}

func CreateItem(c *gin.Context) {
	var item models.Item
	listId, parseError := strconv.ParseUint(c.Param("listId"), 10, 32)
	if parseError != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	bindError := c.BindJSON(&item)
	if bindError != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	item.ListID = uint(listId)
	db := models.CreateItem(&item)
	if db.Error != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, item)
	}
}

func RetrieveItems(c *gin.Context) {
	var list models.List
	var items []models.Item
	id, parseError := strconv.ParseUint(c.Param("listId"), 10, 32)
	if parseError != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	list.ID = uint(id)
	db := models.GetAllItemsOnList(&list, &items)
	if db.Error != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else if len(items) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, items)
	}
}

func RetrieveItem(c *gin.Context) {
	var item models.Item
	itemId, itemIdParseError := strconv.ParseUint(c.Param("itemId"), 10, 32)
	if itemIdParseError != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	listId, listIdParseError := strconv.ParseUint(c.Param("listId"), 10, 32)
	if listIdParseError != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	item.ID = uint(itemId)
	item.ListID = uint(listId)
	db := models.GetItem(&item)
	if db.Error != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, item)
	}
}

func UpdateItem(c *gin.Context) {
	var item models.Item
	bindError := c.BindJSON(&item)
	if bindError != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	db := models.UpdateItem(&item)
	if db.Error != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, item)
	}
}

func DeleteItem(c *gin.Context) {
	var item models.Item
	listId, listIdParseError := strconv.ParseUint(c.Param("listId"), 10, 32)
	if listIdParseError != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	item.ListID = uint(listId)
	itemId, itemIdParseError := strconv.ParseUint(c.Param("itemId"), 10, 32)
	if itemIdParseError != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	item.ID = uint(itemId)
	db := models.DeleteItem(&item)
	if db.Error != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"rowsAffected": db.RowsAffected,
		})
	}
}

func CompleteItem(c *gin.Context) {
	var item models.Item
	listId, listIdParseError := strconv.ParseUint(c.Param("listId"), 10, 32)
	if listIdParseError != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	item.ListID = uint(listId)
	itemId, itemIdParseError := strconv.ParseUint(c.Param("itemId"), 10, 32)
	if itemIdParseError != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	item.ID = uint(itemId)
	doneAt := time.Now()
	db := models.CompleteItem(&item, doneAt)
	if db.Error != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"doneAt":       doneAt,
			"rowsAffected": db.RowsAffected,
		})
	}
}

func IncompleteItem(c *gin.Context) {
	var item models.Item
	listId, listIdParseError := strconv.ParseUint(c.Param("listId"), 10, 32)
	if listIdParseError != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	item.ListID = uint(listId)
	itemId, itemIdParseError := strconv.ParseUint(c.Param("itemId"), 10, 32)
	if itemIdParseError != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	item.ID = uint(itemId)
	db := models.IncompleteItem(&item)
	if db.Error != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"rowsAffected": db.RowsAffected,
		})
	}
}
