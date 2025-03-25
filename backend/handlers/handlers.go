package handlers

import (
	"github.com/alexgiura/h4u/repository"
)

// Use the Interface for the Repository
type Handler struct {
	Repo repository.RepositoryInterface
}

// NewHandler initializes a generic handler
func NewHandler(repo repository.RepositoryInterface) *Handler {
	return &Handler{Repo: repo}
}
