package test

import (
	"encoding/json"
	"github.com/miguoliang/arch-go/internal/dto"
	"github.com/miguoliang/arch-go/pkg/str"
	"github.com/miguoliang/keycloakadminclient"
	"github.com/stretchr/testify/suite"
	"testing"
)

type RoleTestSuite struct {
	Suite
}

func (s *RoleTestSuite) TestRoleCreateSucceed() {

	role := &keycloakadminclient.RoleRepresentation{
		Name: str.Ptr(s.T().Name()),
	}
	w := s.Post("/api/v1/roles", role)
	s.Equal(201, w.Code)
	var created dto.CreatedResponse
	err := json.Unmarshal(w.Body.Bytes(), &created)
	s.NoError(err)
	s.NotEmpty(created.Id)
}

func (s *RoleTestSuite) TestRoleCreateConflict() {

	role := &keycloakadminclient.RoleRepresentation{
		Name: str.Ptr(s.T().Name()),
	}
	w := s.Post("/api/v1/roles", role)
	s.Equal(201, w.Code)
	w = s.Post("/api/v1/roles", role)
	s.Equal(409, w.Code)
}

func (s *RoleTestSuite) TestRoleDeleteSucceed() {

	role := &keycloakadminclient.RoleRepresentation{
		Name: str.Ptr(s.T().Name()),
	}
	w := s.Post("/api/v1/roles", role)
	s.Equal(201, w.Code)
	var created dto.CreatedResponse
	err := json.Unmarshal(w.Body.Bytes(), &created)
	s.NoError(err)
	s.NotEmpty(created.Id)

	w = s.Delete("/api/v1/roles/" + created.Id)
	s.Equal(204, w.Code)
}

func (s *RoleTestSuite) TestRoleDeleteNotFound() {

	w := s.Delete("/api/v1/roles/" + "not-exist")
	s.Equal(404, w.Code)
}

func (s *RoleTestSuite) TestRoleGetSucceed() {

	role := &keycloakadminclient.RoleRepresentation{
		Name: str.Ptr(s.T().Name()),
	}
	w := s.Post("/api/v1/roles", role)
	s.Equal(201, w.Code)
	var created dto.CreatedResponse
	err := json.Unmarshal(w.Body.Bytes(), &created)
	s.NoError(err)
	s.NotEmpty(created.Id)

	w = s.Get("/api/v1/roles/" + created.Id)
	s.Equal(200, w.Code)
	var roleResp keycloakadminclient.RoleRepresentation
	err = json.Unmarshal(w.Body.Bytes(), &roleResp)
	s.NoError(err)
	s.Equal(role.GetName(), roleResp.GetName())
}

func (s *RoleTestSuite) TestRoleGetNotFound() {

	w := s.Get("/api/v1/roles/" + "not-exist")
	s.Equal(404, w.Code)
}

func (s *RoleTestSuite) TestRoleUpdateSucceed() {

	role := &keycloakadminclient.RoleRepresentation{
		Name: str.Ptr(s.T().Name()),
	}
	w := s.Post("/api/v1/roles", role)
	s.Equal(201, w.Code)
	var created dto.CreatedResponse
	err := json.Unmarshal(w.Body.Bytes(), &created)
	s.NoError(err)
	s.NotEmpty(created.Id)

	role.Name = str.Ptr("new name")
	w = s.Put("/api/v1/roles/"+created.Id, role)
	s.Equal(204, w.Code)
}

func (s *RoleTestSuite) TestRoleUpdateNotFound() {

	role := &keycloakadminclient.RoleRepresentation{
		Name: str.Ptr(s.T().Name()),
	}
	w := s.Put("/api/v1/roles/"+"not-exist", role)
	s.Equal(404, w.Code)
}

func (s *RoleTestSuite) TestRoleUpdateConflict() {

	role := &keycloakadminclient.RoleRepresentation{
		Name: str.Ptr(s.T().Name()),
	}
	w := s.Post("/api/v1/roles", role)
	s.Equal(201, w.Code)
	var created dto.CreatedResponse
	err := json.Unmarshal(w.Body.Bytes(), &created)
	s.NoError(err)
	s.NotEmpty(created.Id)

	w = s.Post("/api/v1/roles", role)
	s.Equal(409, w.Code)
}

func (s *RoleTestSuite) TestRoleListSucceed() {

	role := &keycloakadminclient.RoleRepresentation{
		Name: str.Ptr(s.T().Name()),
	}
	w := s.Post("/api/v1/roles", role)
	s.Equal(201, w.Code)

	w = s.Get("/api/v1/roles")
	s.Equal(200, w.Code)
	var roles []keycloakadminclient.RoleRepresentation
	err := json.Unmarshal(w.Body.Bytes(), &roles)
	s.NoError(err)
	s.NotEmpty(roles)
}

func (s *RoleTestSuite) TestCheckRoleNameSuccess() {

	role := &keycloakadminclient.RoleRepresentation{
		Name: str.Ptr(s.T().Name()),
	}
	w := s.Post("/api/v1/roles", role)
	s.Equal(201, w.Code)

	w = s.Head("/api/v1/roles?roleName=" + s.T().Name())
	s.Equal(204, w.Code)
}

func (s *RoleTestSuite) TestCheckRoleNameNotFound() {

	w := s.Head("/api/v1/roles?roleName=" + "not-exist")
	s.Equal(404, w.Code)
}

func TestRoleTestSuite(t *testing.T) {
	suite.Run(t, new(RoleTestSuite))
}
