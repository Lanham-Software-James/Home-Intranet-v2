package library

import (
	"Home-Intranet-v2-Backend/internal/library/models"
	"Home-Intranet-v2-Backend/internal/platform/logger"
	"Home-Intranet-v2-Backend/internal/platform/response"
	"fmt"
	"net/http"
)

func (handler Handler) CreateBook(w http.ResponseWriter, request *http.Request) {
	var book = models.Book{
		Title:           "Test Book",
		AuthorFirstName: "Jimmy",
		AuthorLastName:  "James",
	}
	logger.Debug("Made it!")

	err := handler.Repository.Create(request.Context(), &book)

	// logger.Debug(fmt.Sprintf("Raw data from List: %+v", data))
	if err != nil {
		logger.Error(fmt.Sprintf("Issue retriving books. \nError: %s", err.Error()))
		response.InternalServerError(w, err)
		return
	}

	response.SuccessResponse(w, &book)
	return
}
