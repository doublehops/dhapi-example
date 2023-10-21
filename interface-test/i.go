package main

import (
	"context"
	"encoding/json"
	"fmt"
)

type service interface {
	GetByID(ctx context.Context, author Model, ID int32) error
}

type AuthorService struct{}

func (as AuthorService) GetByID(ctx context.Context, author Model, ID int32) error {
	return nil
}

type Model interface {
	Unmarshal(data []byte) error
}

type AuthorModel struct{}

func (am *AuthorModel) Unmarshal(data []byte) error {
	return json.Unmarshal(data, am)
}

func ProcessUpdate(model Model, service service) error {
	return fmt.Errorf("an error was found")
}

func main() {
	am := &AuthorModel{}
	as := AuthorService{}
	ProcessUpdate(am, as)
}
