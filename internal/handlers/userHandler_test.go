package handlers

import (
	"context"
	"errors"

	middlewear "job-portal-api/internal/middleware"
	model "job-portal-api/internal/models"
	"job-portal-api/internal/services"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/mock/gomock"
)

func Test_handler_userSignin(t *testing.T) {

	tests := []struct {
		name               string
		h                  *handler
		setup              func() (*gin.Context, *httptest.ResponseRecorder, services.UsersService)
		expectedStatusCode int
		expectedResponse   string
	}{
		{name: "missing trace id",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.UsersService) {
				rr := httptest.NewRecorder()      // responseRecorder
				c, _ := gin.CreateTestContext(rr) // context
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
				c.Request = httpRequest
				return c, rr, nil
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"msg":"Internal Server Error"}`,
		},
		{
			name: "validate request body",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.UsersService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", strings.NewReader(`{"name":"",
				"email":    "siddarth@gmail.com",
				"password": "siddarth"}`))
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middlewear.TraceIdKey, "123")
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest

				return c, rr, nil
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"msg":"invalid input"}`,
		},
		{
			name: "validate request body failure",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.UsersService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", strings.NewReader(`{error}`))
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middlewear.TraceIdKey, "123")
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest

				return c, rr, nil
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"msg":"Internal Server Error"}`,
		},
		{
			name: "success",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.UsersService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", strings.NewReader(`{"name":"cece",
				"email":    "siddarth@gmail.com",
				"password": "siddarth"}`))
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middlewear.TraceIdKey, "123")
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest
				c.Params = append(c.Params, gin.Param{Key: "id", Value: "123"})
				mc := gomock.NewController(t)
				ms := services.NewMockUsersService(mc)

				ms.EXPECT().UserSignup(gomock.Any()).Return(model.User{}, nil).AnyTimes()

				return c, rr, ms
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   `{"ID":0,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"name":"","email":""}`,
		},
		{
			name: "failure case",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.UsersService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", strings.NewReader(`{"name":"cece",
				"email":    "siddarth@gmail.com",
				"password": "siddarth"}`))
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middlewear.TraceIdKey, "123")
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest
				c.Params = append(c.Params, gin.Param{Key: "id", Value: "123"})
				mc := gomock.NewController(t)
				ms := services.NewMockUsersService(mc)

				ms.EXPECT().UserSignup(gomock.Any()).Return(model.User{}, errors.New("error")).AnyTimes()

				return c, rr, ms
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"msg":"user signup failed"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			c, rr, ms := tt.setup()
			h := &handler{
				us: ms,
			}
			h.userSignin(c)
			assert.Equal(t, tt.expectedStatusCode, rr.Code)
			assert.Equal(t, tt.expectedResponse, rr.Body.String())
		})
	}
}

func Test_handler_userLoginin(t *testing.T) {

	tests := []struct {
		name               string
		setup              func() (*gin.Context, *httptest.ResponseRecorder, services.UsersService)
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name: "missing trace id",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.UsersService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
				c.Request = httpRequest

				return c, rr, nil
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"msg":"Internal Server Error"}`,
		},
		{
			name: "validate request body",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.UsersService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", strings.NewReader(`{
			"email":    "",
			"password": "hfhhfhfh"}`))
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middlewear.TraceIdKey, "123")
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest

				return c, rr, nil
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"msg":"invalid input"}`,
		},
		{
			name: "validate request body failure",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.UsersService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", strings.NewReader(`{"error}`))
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middlewear.TraceIdKey, "123")
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest

				return c, rr, nil
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"msg":"Internal Server Error"}`,
		},
		{
			name: "validate request body",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.UsersService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", strings.NewReader(`{
			"email":    "",
			"password": "hfhhfhfh"}`))
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middlewear.TraceIdKey, "123")
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest

				return c, rr, nil
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"msg":"invalid input"}`,
		},
		{
			name: "success case",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.UsersService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", strings.NewReader(`{
			"email":    "",
			"password": "hfhhfhfh"}`))
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middlewear.TraceIdKey, "123")
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest
				c.Params = append(c.Params, gin.Param{Key: "id", Value: "123"})
				mc := gomock.NewController(t)
				ms := services.NewMockUsersService(mc)

				ms.EXPECT().Userlogin(gomock.Any()).Return(jwt.RegisteredClaims{}, nil).AnyTimes()

				return c, rr, nil
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"msg":"invalid input"}`,
		},
		// {
		// 	name: "success",
		// 	setup: func() (*gin.Context, *httptest.ResponseRecorder, services.UsersService) {
		// 		rr := httptest.NewRecorder()
		// 		c, _ := gin.CreateTestContext(rr)
		// 		httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", strings.NewReader(`{
		// 		"email":    "siddarth@gmail.com",
		// 		"password": "siddarth"}`))
		// 		ctx := httpRequest.Context()
		// 		ctx = context.WithValue(ctx, middlewear.TraceIdKey, "123")
		// 		httpRequest = httpRequest.WithContext(ctx)
		// 		c.Request = httpRequest
		// 		c.Params = append(c.Params, gin.Param{Key: "id", Value: "123"})
		// 		mc := gomock.NewController(t)
		// 		ms := services.NewMockUsersService(mc)

		// 		ms.EXPECT().Userlogin(gomock.Any()).Return(jwt.RegisteredClaims{}, nil).AnyTimes()

		// 		return c, rr, ms
		// 	},
		// 	expectedStatusCode: http.StatusOK,
		// 	expectedResponse:   `jwt.RegisteredClaims{Issuer: "service project", Subject: "0", Audience: jwt.ClaimStrings{"users"}, ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)), IssuedAt: jwt.NewNumericDate(time.Now())}`,
		// },
		// {
		// 	name: "failure case",
		// 	setup: func() (*gin.Context, *httptest.ResponseRecorder, services.UsersService) {
		// 		rr := httptest.NewRecorder()
		// 		c, _ := gin.CreateTestContext(rr)
		// 		httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", strings.NewReader(`{
		// 	"email":    "",
		// 	"password": "hfhhfhfh"}`))
		// 		ctx := httpRequest.Context()
		// 		ctx = context.WithValue(ctx, middlewear.TraceIdKey, "123")
		// 		httpRequest = httpRequest.WithContext(ctx)
		// 		c.Request = httpRequest
		// 		c.Params = append(c.Params, gin.Param{Key: "id", Value: "123"})
		// 		mc := gomock.NewController(t)
		// 		ms := services.NewMockUsersService(mc)

		// 		ms.EXPECT().Userlogin(gomock.Any()).Return(nil, errors.New("error")).AnyTimes()

		// 		return c, rr, nil
		// 	},
		// 	expectedStatusCode: http.StatusBadRequest,
		// 	expectedResponse:   `{"msg": "user signup failed"}`,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			c, rr, ms := tt.setup()
			h := &handler{
				us: ms,
			}
			h.userLoginin(c)
			assert.Equal(t, tt.expectedStatusCode, rr.Code)
			assert.Equal(t, tt.expectedResponse, rr.Body.String())
		})
	}
}
