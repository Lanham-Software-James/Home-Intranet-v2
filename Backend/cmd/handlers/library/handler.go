package library

import (
	"Home-Intranet-v2-Backend/internal/platform/repository"
)

// Handler is used to allow us to pass our data persistance objects as mocks for better testing
type Handler struct {
	Repository *repository.Repository
}
