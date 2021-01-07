package service

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"algogrit.com/emp-server/employee/entities"
	"algogrit.com/emp-server/employee/repository"
	"github.com/golang/mock/gomock"
)

func TestIndex(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockEmployeeRepository(ctrl)

	sut := NewJSON(mockRepo)

	mockRepo.EXPECT().RetrieveAll().Return([]entities.Employee{})

	req := httptest.NewRequest("GET", "/employees", nil)
	w := httptest.NewRecorder()

	sut.Index(w, req)

	resp := w.Result()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
