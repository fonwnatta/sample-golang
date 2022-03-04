package mock

import (
	"errors"
	"golang-training/main/internal/repository"
	"github.com/stretchr/testify/mock"
)

type userDeatail struct {
	mock.Mock
}

func (m *userDeatail) Insert(entity repository.UserDetailEntity) error { //insert to db 
	return m.Called(entity).Error(0)
}

func (m *userDeatail) FindAllData() ([]repository.UserDetailEntity, error) {
	args := m.Called()
	if args.Get(0) != nil {
		return args.Get(0).([]repository.UserDetailEntity), nil
	}
	return nil, args.Error(1)
}
func (m *userDeatail) Update(entity repository.UserDetailEntity) error { //upsate to db
	return m.Called(entity).Error(0)
}


func UserDeatail(tc string) *userDeatail {
	m := new(userDeatail)
	switch tc {
	case "OK" : 
		m.On("Insert", mock.AnythingOfType("repository.UserDetailEntity")).Return(nil)
		m.On("FindAllData", mock.AnythingOfType("")).Return([]repository.UserDetailEntity{}, nil)
		m.On("Update", mock.AnythingOfType("repository.UserDetailEntity")).Return(nil)

	case "NOK" :
		m.On("Insert", mock.AnythingOfType("repository.UserDetailEntity")).Return(errors.New("Requrid error"))
		m.On("FindAllData", mock.AnythingOfType("")).Return(nil, errors.New("Database error"))
		m.On("Update", mock.AnythingOfType("repository.UserDetailEntity")).Return(errors.New("Requrid error"))
	}
	return m 
}