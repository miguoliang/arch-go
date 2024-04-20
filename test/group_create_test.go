package test

import (
	"github.com/gin-gonic/gin"
	"github.com/miguoliang/arch-go/internal/resource"
	"github.com/miguoliang/arch-go/pkg/str"
	"github.com/miguoliang/keycloakadminclient"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateGroup_Succeed(t *testing.T) {

	r := gin.Default()
	r.POST("/groups", resource.CreateGroupHandler)

	group := &keycloakadminclient.GroupRepresentation{
		Name: str.Ptr(t.Name()),
	}
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/groups", str.StructToJsonReader(group))
	r.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("expected %d, got %d", http.StatusCreated, w.Code)
	}
}

func TestCreateGroup_Conflict(t *testing.T) {

	r := gin.Default()
	r.POST("/groups", resource.CreateGroupHandler)

	group := &keycloakadminclient.GroupRepresentation{
		Name: str.Ptr(t.Name()),
	}
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/groups", str.StructToJsonReader(group))
	r.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("expected %d, got %d", http.StatusCreated, w.Code)
	}

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/groups", str.StructToJsonReader(group))
	r.ServeHTTP(w, req)

	if w.Code != http.StatusConflict {
		t.Errorf("expected %d, got %d", http.StatusConflict, w.Code)
	}
}

func TestCreateGroup_BadRequest_WhenNameIsEmpty(t *testing.T) {

	r := gin.Default()
	r.POST("/groups", resource.CreateGroupHandler)

	group := &keycloakadminclient.GroupRepresentation{}
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/groups", str.StructToJsonReader(group))
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected %d, got %d", http.StatusBadRequest, w.Code)
	}
}
