// Package library contains all the controllers for the library functionality
package library

import (
	"Home-Intranet-v2-Backend/internal/library/models"
	"Home-Intranet-v2-Backend/internal/platform/logger"
	"Home-Intranet-v2-Backend/internal/platform/response"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

// ListBooks returns a list of books based on the parameters the user enter
func (handler Handler) ListBooks(w http.ResponseWriter, request *http.Request) {
	values := request.URL.Query()

	sortColumn := strings.ToLower(values.Get("sort-col"))
	sortDirectionString := strings.ToLower(values.Get("sort-dir"))
	offsetString := values.Get("offset")
	limitString := values.Get("limit")

	// TODO: Handle filter, search

	if sortColumn == "" {
		sortColumn = "title"
	}

	sort := 1
	if sortDirectionString == "desc" {
		sort = -1
	}

	if offsetString == "" {
		offsetString = "0"
	}

	if limitString == "" {
		limitString = "20"
	}

	offset, err := strconv.ParseInt(offsetString, 10, 64)
	if err != nil {
		logger.Error(fmt.Sprintf("Error converting offset to int: %v", err))
		response.BadRequest(w, err)
		return
	}

	limit, err := strconv.ParseInt(limitString, 10, 64)
	if err != nil {
		logger.Error(fmt.Sprintf("Error converting offset to int: %v", err))
		response.BadRequest(w, err)
		return
	}

	data, err := handler.Repository.List(request.Context(), &models.Book{}, bson.D{}, bson.D{{Key: "shelf", Value: 1}, {Key: sortColumn, Value: sort}}, offset, limit)
	if err != nil {
		logger.Error(fmt.Sprintf("Issue retriving books. \nError: %s", err.Error()))
		response.InternalServerError(w, err)
		return
	}

	var books []models.Book
	for _, element := range data {
		var book models.Book

		bsonData, err := bson.Marshal(element)
		if err != nil {
			logger.Error(fmt.Sprintf("Error marshaling data: %s", err.Error()))
			continue
		}

		err = bson.Unmarshal(bsonData, &book)
		if err != nil {
			logger.Error(fmt.Sprintf("Error unmarshaling data: %s", err.Error()))
			continue
		}

		books = append(books, book)
	}

	response.SuccessResponse(w, books)
	return
}
