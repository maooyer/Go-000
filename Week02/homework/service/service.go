package service

import (
	"github.com/maooyer/Go-000/Week02/homework/dao"
	"github.com/pkg/errors"
)

type Service struct {
}

func (s *Service) GetObject() {
	err := dao.Dao()
	if errors.Is(err, dao.ErrNotFound) {
		// do something
	}

}
