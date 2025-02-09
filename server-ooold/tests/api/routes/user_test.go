package tests

import "github.com/stretchr/testify/mock"

type MockUserInterface struct {
	mock.Mock
}

func (m *MockUserInterface) GetUserByID(id int) (string, error) {
    args := m.Called(id)
    return args.String(0), args.Error(1)
}

func Test_GetUs90210_StatusCodeShouldEqual200() {
	response, err :=
		BadExpr
}
