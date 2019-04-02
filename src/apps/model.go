package apps

import (
	"libs/abstract"
)

type GenericModel interface {
	GetCollectionName() string
	GetAbstractModel() *abstract.AbstractModel
	Clean() error
}

type SaveFunction func(string, GenericModel) error

func Save(saveFunction SaveFunction, instance GenericModel) error {
	abstractInstance := instance.GetAbstractModel()
	abstractInstance.Clean()
	err := instance.Clean()

	err = saveFunction(instance.GetCollectionName(), instance)
	return err
}
