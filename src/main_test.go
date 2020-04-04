package main

import (
	"checklist/database"
	"checklist/routes"
	"encoding/json"
	"github.com/jinzhu/gorm"
	"net/http"
	"net/http/httptest"
	"testing"

	mocket "github.com/selvatico/go-mocket"
	"github.com/stretchr/testify/assert"
)

func setupMockDb() *gorm.DB {
	mocket.Catcher.Register() // Safe register. Allowed multiple calls to save
	mocket.Catcher.Logging = true
	db, err := gorm.Open(mocket.DriverName, "connection_string") // Can be any connection string
	if err != nil {
		return nil
	}
	db.LogMode(true)
	return db
}

func TestResponses(t *testing.T) {
	database.DB = setupMockDb()
	router := routes.SetupRouter()

	t.Run("GET /v1/list?withItems=true", func(t *testing.T) {
		var mockListsResponse []map[string]interface{}
		listErr := json.Unmarshal([]byte(listsResponse), &mockListsResponse)
		if listErr != nil {
			t.Error(listErr)
		}
		var mockItemsResponse []map[string]interface{}
		itemErr := json.Unmarshal([]byte(itemsResponse), &mockItemsResponse)
		if itemErr != nil {
			t.Error(listErr)
		}
		mocket.Catcher.Reset()
		mocket.Catcher.NewMock().WithQuery(`SELECT * FROM "lists"`).WithReply(mockListsResponse)
		mocket.Catcher.NewMock().WithArgs(int64(47), int64(46)).WithReply(mockItemsResponse)

		response := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/v1/list?withItems=true", nil)
		router.ServeHTTP(response, req)

		assert.Equal(t, 200, response.Code)
	})

	t.Run("GET /v1/list/47/item/", func(t *testing.T) {
		var mockItemsResponse []map[string]interface{}
		itemErr := json.Unmarshal([]byte(list47sItems), &mockItemsResponse)
		if itemErr != nil {
			t.Error(itemErr)
		}
		mocket.Catcher.Reset().NewMock().WithArgs(int64(47)).WithReply(mockItemsResponse)

		response := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/v1/list/47/item/", nil)
		router.ServeHTTP(response, req)

		assert.Equal(t, 200, response.Code)
	})

}

var listsResponse = `
[
    {
        "id": 46,
        "createdAt": "2020-03-06T17:43:42.956558Z",
        "updatedAt": "2020-03-06T17:43:42.956558Z",
        "title": "foos",
        "category": "",
        "items": null
    },
    {
        "id": 47,
        "createdAt": "2020-03-06T19:17:45.08012Z",
        "updatedAt": "2020-03-06T19:17:45.08012Z",
        "title": "bars",
        "category": "",
        "items": null
    }
]
`

var itemsResponse = `
[
	{
		"id": 20,
		"createdAt": "2020-03-06T17:43:42.961355Z",
		"updatedAt": "2020-03-06T17:43:42.961355Z",
		"listId": 46,
		"description": "item 1",
		"done": 0
	},
	{
		"id": 21,
		"createdAt": "2020-03-06T17:43:42.966391Z",
		"updatedAt": "2020-03-06T17:43:42.966391Z",
		"listId": 46,
		"description": "item 2",
		"done": 0
	},
	{
		"id": 22,
		"createdAt": "2020-03-06T19:17:45.09156Z",
		"updatedAt": "2020-03-06T19:17:45.09156Z",
		"listId": 47,
		"description": "item 1",
		"done": 0
	},
	{
		"id": 23,
		"createdAt": "2020-03-06T19:17:45.097794Z",
		"updatedAt": "2020-03-06T19:17:45.097794Z",
		"listId": 47,
		"description": "item 2",
		"done": 0
	}
]
`

var list47sItems = `
[
	{
		"id": 22,
		"createdAt": "2020-03-06T19:17:45.09156Z",
		"updatedAt": "2020-03-06T19:17:45.09156Z",
		"listId": 47,
		"description": "item 1",
		"done": 0
	},
	{
		"id": 23,
		"createdAt": "2020-03-06T19:17:45.097794Z",
		"updatedAt": "2020-03-06T19:17:45.097794Z",
		"listId": 47,
		"description": "item 2",
		"done": 0
	}
]
`
