package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jackyrusly/jrgo/dto"
	"github.com/jackyrusly/jrgo/mocks"
	"github.com/jackyrusly/jrgo/services"
	"github.com/jackyrusly/jrgo/utils"
	"github.com/labstack/echo/v4"
)

func TestAuthController_ControllerRegister(t *testing.T) {
	type fields struct {
		as services.IAuthService
	}

	type routerFields struct {
		c   echo.Context
		rec *httptest.ResponseRecorder
	}

	createAuthRouter := func(json string) routerFields {
		e := utils.NewRouter()
		req := httptest.NewRequest(http.MethodPost, "/auth/register", strings.NewReader(json))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		return routerFields{
			c:   c,
			rec: rec,
		}
	}

	body := dto.RegisterRequestBody{
		Username: "jrusly",
		Name:     "Jacky Rusly",
		Password: "password",
	}

	bodyJson, _ := json.Marshal(body)

	tests := []struct {
		name         string
		createFields func() fields
		router       routerFields
		wantStatus   int
	}{
		{
			name: "SuccessRegister",
			createFields: func() fields {
				mockProfileService := mocks.NewIAuthService(t)
				mockProfileService.On("ServiceRegister", body).Return(nil)

				return fields{
					as: mockProfileService,
				}
			},
			router:     createAuthRouter(string(bodyJson)),
			wantStatus: http.StatusNoContent,
		},
		{
			name: "FailRegister",
			createFields: func() fields {
				mockProfileService := mocks.NewIAuthService(t)
				mockProfileService.On("ServiceRegister", body).Return(errors.New("Test error"))

				return fields{
					as: mockProfileService,
				}
			},
			router:     createAuthRouter(string(bodyJson)),
			wantStatus: http.StatusBadRequest,
		},
		{
			name: "InvalidJson",
			createFields: func() fields {
				mockProfileService := mocks.NewIAuthService(t)

				return fields{
					as: mockProfileService,
				}
			},
			router:     createAuthRouter(`{"invalid-json}`),
			wantStatus: http.StatusBadRequest,
		},
		{
			name: "InvalidRequestBody",
			createFields: func() fields {
				mockProfileService := mocks.NewIAuthService(t)

				return fields{
					as: mockProfileService,
				}
			},
			router:     createAuthRouter(`{"username":"jrusly"}`),
			wantStatus: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fields := tt.createFields()
			ac := NewAuthController(fields.as)

			ac.ControllerRegister(tt.router.c)

			if tt.router.rec.Result().StatusCode != tt.wantStatus {
				t.Errorf("AuthController.ControllerRegister() error = %v, wantErr %v", tt.router.rec.Result().StatusCode, tt.wantStatus)
			}
		})
	}
}
