package repository

type repository interface {
	Create(interface{}) interface{}
}
