// Package library contains all the controllers for the library functionality
package library

import (
	"Home-Intranet-v2-Backend/internal/library/models"
	"Home-Intranet-v2-Backend/internal/platform/logger"
	"Home-Intranet-v2-Backend/internal/platform/response"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// ListBooks returns a list of books based on the parameters the user enter
func (handler Handler) ListBooks(w http.ResponseWriter, request *http.Request) {
	values := request.URL.Query()

	sortColumn := strings.ToLower(values.Get("sort-col"))
	sortDirectionString := strings.ToLower(values.Get("sort-dir"))
	offsetString := values.Get("offset")
	limitString := values.Get("limit")

	// TODO: Handle Seach and Filter

	// Build sort map
	if sortColumn == "" {
		sortColumn = "title"
	}

	sortDirection := 1
	if sortDirectionString == "desc" {
		sortDirection = -1
	}

	sort := map[string]string{
		"shelf":    "1",
		sortColumn: strconv.Itoa(sortDirection),
	}

	// Get default offset value
	if offsetString == "" {
		offsetString = "0"
	}

	// Get default limit value
	if limitString == "" {
		limitString = "20"
	}

	// Convert to ints
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

	// Query
	data, err := handler.Repository.List(request.Context(), &models.Book{}, map[string]string{}, sort, offset, limit)
	if err != nil {
		logger.Error(fmt.Sprintf("Issue retriving books. \nError: %s", err.Error()))
		response.InternalServerError(w, err)
		return
	}

	var books []models.Book
	err = json.Unmarshal(data, &books)
	if err != nil {
		logger.Error(fmt.Sprintf("Error unmarshaling data: %s", err.Error()))
		response.BadRequest(w, err)
		return
	}

	response.SuccessResponse(w, books)
	return
}
