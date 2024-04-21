package test

import (
	"encoding/json"
	_ "github.com/miguoliang/arch-go/configs"
	"github.com/miguoliang/arch-go/internal/dto"
	"github.com/miguoliang/arch-go/pkg/str"
	"github.com/miguoliang/keycloakadminclient"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

type GroupTestSuite struct {
	Suite
}

func (s *GroupTestSuite) TestCreateGroupSucceed() {

	group := &keycloakadminclient.GroupRepresentation{
		Name: str.Ptr(s.T().Name()),
	}
	w := s.Post("/api/v1/groups", group)
	s.Equal(http.StatusCreated, w.Code)
	var created dto.CreatedResponse
	err := json.Unmarshal(w.Body.Bytes(), &created)
	s.NoError(err)
	s.NotEmpty(created.Id)
}

func (s *GroupTestSuite) TestCreateGroupConflict() {

	group := &keycloakadminclient.GroupRepresentation{
		Name: str.Ptr(s.T().Name()),
	}
	w := s.Post("/api/v1/groups", group)
	s.Equal(http.StatusCreated, w.Code)

	w = s.Post("/api/v1/groups", group)
	s.Equal(http.StatusConflict, w.Code)
}

func (s *GroupTestSuite) TestCreateGroupBadRequestWhenNameIsEmpty() {

	group := &keycloakadminclient.GroupRepresentation{}
	w := s.Post("/api/v1/groups", group)
	s.Equal(http.StatusBadRequest, w.Code)
}

func (s *GroupTestSuite) TestDeleteGroupSucceedIfExists() {

	group := &keycloakadminclient.GroupRepresentation{
		Name: str.Ptr(s.T().Name()),
	}
	w := s.Post("/api/v1/groups", group)
	s.Equal(http.StatusCreated, w.Code)
	err := json.Unmarshal(w.Body.Bytes(), &group)
	s.NoError(err)
	s.NotEmpty(group.Id)

	w = s.Delete("/api/v1/groups/" + *group.Id)
	s.Equal(http.StatusNoContent, w.Code)
}

func (s *GroupTestSuite) TestDeleteGroupNotFoundIfNotExists() {

	w := s.Delete("/api/v1/groups/123")
	s.Equal(http.StatusNotFound, w.Code)
}

func (s *GroupTestSuite) TestGetGroupSucceedIfExists() {

	group := &keycloakadminclient.GroupRepresentation{
		Name: str.Ptr(s.T().Name()),
	}
	w := s.Post("/api/v1/groups", group)
	s.Equal(http.StatusCreated, w.Code)
	var created dto.CreatedResponse
	err := json.Unmarshal(w.Body.Bytes(), &created)
	s.NoError(err)
	s.NotEmpty(created.Id)

	w = s.Get("/api/v1/groups/" + created.Id)
	s.Equal(http.StatusOK, w.Code)
	var got keycloakadminclient.GroupRepresentation
	err = json.Unmarshal(w.Body.Bytes(), &got)
	s.NoError(err)
	s.Equal(created.Id, got.GetId())
	s.Equal(group.GetName(), got.GetName())
}

func TestGroupTestSuite(t *testing.T) {
	suite.Run(t, new(GroupTestSuite))
}
