package controller

import (
	"encoding/json"
	"net/http"

	"github.com/giovanifranz/testes-go/internal/infra/database"
	"github.com/giovanifranz/testes-go/internal/usecase"
)

func (b *BaseHandler) CreateClientHandler(w http.ResponseWriter, r *http.Request) {

	var content usecase.CreateClientImputDTO

	err := json.NewDecoder(r.Body).Decode(&content)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	repository := database.NewClientRepository(b.db)
	uc := usecase.NewClientClientUseCase(repository)
	_, err = uc.Execute(content)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
