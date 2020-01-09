package user

import (
	"fmt"
	"testing"

	"github.com/yinliang-gl/golang_project/test_project/test_gomock/mock"

	"github.com/golang/mock/gomock"
)

func TestUser_GetUserInfo(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	var id int64 = 1
	mockMale := mock.NewMockMale(ctl)
	gomock.InOrder(
		mockMale.EXPECT().Get(id).Return(999.00),
	)

	user := NewUser(mockMale)
	val := user.GetUserInfo(id)
	fmt.Println(val)
}
