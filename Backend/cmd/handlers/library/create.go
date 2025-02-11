package library

import (
	"Home-Intranet-v2-Backend/internal/library/models"
	"Home-Intranet-v2-Backend/internal/platform/logger"
	"Home-Intranet-v2-Backend/internal/platform/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func (handler Handler) CreateBook(w http.ResponseWriter, request *http.Request) {
	var book models.Book

	byteData, err := io.ReadAll(request.Body)
	if err != nil {
		logger.Error(fmt.Sprintf("Issue reading request body. \nError: %+v", err.Error()))
		response.InternalServerError(w, err)
		return
	}

	err = json.Unmarshal(byteData, &book)
	if err != nil {
		logger.Error(fmt.Sprintf("Issue unmarshalling json. \nError: %+v", err.Error()))
		response.InternalServerError(w, err)
		return
	}

	if book.CheckedOut {
		book.CheckedOutTime = time.Now()
	}

	err = handler.Repository.Create(request.Context(), &book)
	if err != nil {
		logger.Error(fmt.Sprintf("Issue creating book. \nError: %+v", err.Error()))
		response.InternalServerError(w, err)
		return
	}

	for _, author := range book.Authors {
		if err = handler.Repository.Read(request.Context(), &author, bson.D{
			{Key: "first_name", Value: author.FirstName},
			{Key: "middle_name", Value: author.MiddleName},
			{Key: "last_name", Value: author.LastName},
		}); err != nil && err.Error() != "mongo: no documents in result" {
			logger.Error(fmt.Sprintf("Issue Finding Document. \nError: %+v", err))
			response.InternalServerError(w, err)
			return
		}

		if author.Model.ID.String() == `ObjectID("000000000000000000000000")` {
			if err = handler.Repository.Create(request.Context(), &author); err != nil {
				logger.Error(fmt.Sprintf("Issue creating author. \nError: %+v", err.Error()))
				response.InternalServerError(w, err)
				return
			}
		}
	}

	response.SuccessResponse(w, &book)
	return
}
