package database

import (
	"errors"

	"github.com/Alish26/bank-system-golang/adapter/repository"
)

var (
	errInvalidNoSQLDatabaseInstance = errors.New("invalid nosql db instance")
)

const (
	InstanceMongoDB int = iota
)

func NewDatabaseNoSQLFactory(instance int) (repository.NoSQL, error) {
	switch instance {
	case InstanceMongoDB:
		return NewMongoHandler(newConfigMongoDB())
	default:
		return nil, errInvalidNoSQLDatabaseInstance
	}
}
