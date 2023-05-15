package test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"new_project/big_hello"
	"new_project/mock_bighello"
)

func TestDoSomething(t *testing.T) {
	// Create a new instance of the controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock implementation of the Random interface
	mockedRandom := mock_bighello.NewMockRandom(ctrl)

	// Set the expected behavior of the mockRandom
	mockedRandom.EXPECT().RandomInt().Return(5)

	// Create a new BigHello instance using the mockRandom
	bigHello := bighello.New(mockedRandom)

	// Call the method being tested
	result := bigHello.DoSomething(2, 3)

	// Assert the result
	assert.Equal(t, 10, result)

}