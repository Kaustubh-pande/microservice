package proto

import (
	"fmt"

	"github.com/jinzhu/gorm"

	uuid "github.com/satori/go.uuid"
)

func (model *User) BeforeCreate(scope *gorm.Scope) error {
	uuid, err := uuid.NewV4()
	if err != nil {
		fmt.Println(err)
	}
	return scope.SetColumn("Id", uuid.String())
}
