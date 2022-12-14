package usecase

import "github.com/giovanifranz/testes-go/internal/entity"

type CreateClientImputDTO struct {
	Name  string
	Email string
}

type CreateClientOutputDTO struct {
	ID     string
	Name   string
	Email  string
	Points int
}

type CreateClientUseCase struct {
	ClientRepository entity.ClientRepositoryInterface
}

func NewClientClientUseCase(clientRepository entity.ClientRepositoryInterface) *CreateClientUseCase {
	return &CreateClientUseCase{
		ClientRepository: clientRepository,
	}
}

func (c *CreateClientUseCase) Execute(input CreateClientImputDTO) (*CreateClientOutputDTO, error) {
	client, err := entity.NewClient(input.Name, input.Email)

	if err != nil {
		return nil, err
	}

	err = c.ClientRepository.Save(client)

	if err != nil {
		return nil, err
	}

	return &CreateClientOutputDTO{
		ID:     client.ID,
		Name:   client.Name,
		Email:  client.Email,
		Points: client.Points,
	}, nil
}
