package server

import (
	mock_repository "homework-8/internal/app/user/mocks"
	"testing"

	"github.com/golang/mock/gomock"
)

type userRepoFixture struct {
	ctrl *gomock.Controller
	repo *mock_repository.MockUserRepo
	serv server
}

func setUp(t *testing.T) userRepoFixture {
	ctrl := gomock.NewController(t)
	repo := mock_repository.NewMockUserRepo(ctrl)
	serv := server{userRepo: repo}

	return userRepoFixture{ctrl: ctrl, repo: repo, serv: serv}
}

func (u *userRepoFixture) tearDown() {
	u.ctrl.Finish()
}
