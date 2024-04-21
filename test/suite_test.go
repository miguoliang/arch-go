package test

import (
	"github.com/gin-gonic/gin"
	"github.com/miguoliang/arch-go/internal/keycloak"
	"github.com/miguoliang/arch-go/internal/resource"
	"github.com/miguoliang/arch-go/pkg/str"
	"github.com/stretchr/testify/suite"
	"golang.org/x/net/context"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
)

type Suite struct {
	suite.Suite
	r *gin.Engine
}

func (s *Suite) SetupSuite() {
	log.Println("Setup suite")
	s.deleteCustomRealm()
	s.createCustomRealm()
	s.r = resource.SetupRoutes()
}

func (s *Suite) createCustomRealm() {

	f, err := os.OpenFile("../configs/realm-export.json", os.O_RDONLY, 0644)
	if err != nil {
		panic("../configs/realm-export.json not found")
	}
	client := keycloak.GetAdminClient()
	response, err := client.RealmsAdminAPI.AdminRealmsPost(context.Background()).
		Body(f).
		Execute()
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
}

func (s *Suite) deleteCustomRealm() {
	client := keycloak.GetAdminClient()
	response, err := client.RealmsAdminAPI.AdminRealmsRealmDelete(context.Background(), resource.CustomRealmName).
		Execute()
	if err != nil {
		return
	}
	defer response.Body.Close()
}

func (s *GroupTestSuite) Get(url string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", url, nil)
	s.r.ServeHTTP(w, req)
	return w
}

func (s *Suite) Post(url string, body interface{}) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", url, str.StructToJsonReader(body))
	s.r.ServeHTTP(w, req)
	return w
}

func (s *Suite) Delete(url string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", url, nil)
	s.r.ServeHTTP(w, req)
	return w
}
