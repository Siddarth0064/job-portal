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
	"go.uber.org/mock/gomock"
)

func Test_handler_companyCreation(t *testing.T) {

	tests := []struct {
		name               string
		setup              func() (*gin.Context, *httptest.ResponseRecorder, services.CompanyService)
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name: "missing trace id",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.CompanyService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
				c.Request = httpRequest

				return c, rr, nil
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"msg":"Internal Server Error"}`,
		},
		{name: "input validation",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.CompanyService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", strings.NewReader(`{"name":"",
			"email":    "name@gmail.com",
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
		{name: "input validation failure",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.CompanyService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", strings.NewReader(`{internal server error}`))
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middlewear.TraceIdKey, "123")
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest

				return c, rr, nil
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"msg":"Internal Server Error"}`,
		}, {
			name: "success",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.CompanyService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", strings.NewReader(`{"company_name":"names",
				"company_adress":    "name@gmail.com",
				"domain": "hfhhfhfh"}`))
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middlewear.TraceIdKey, "123")
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest

				mc := gomock.NewController(t)
				ms := services.NewMockCompanyService(mc)

				ms.EXPECT().CompanyCreate(gomock.Any()).Return(model.Company{}, nil).AnyTimes()

				return c, rr, ms
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   `{"ID":0,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"company_name":"","company_adress":"","domain":""}`,
		},
		{
			name: "failure case",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.CompanyService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", strings.NewReader(`{"company_name":"names",
				"company_adress":    "name@gmail.com",
				"domain": "hfhhfhfh"}`))
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middlewear.TraceIdKey, "123")
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest

				mc := gomock.NewController(t)
				ms := services.NewMockCompanyService(mc)

				ms.EXPECT().CompanyCreate(gomock.Any()).Return(model.Company{}, errors.New("errors")).AnyTimes()

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
				cs: ms,
			}
			h.companyCreation(c)
			assert.Equal(t, tt.expectedStatusCode, rr.Code)
			assert.Equal(t, tt.expectedResponse, rr.Body.String())
		})
	}
}

func Test_handler_getAllCompany(t *testing.T) {

	tests := []struct {
		name               string
		h                  *handler
		setup              func() (*gin.Context, *httptest.ResponseRecorder, services.CompanyService)
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name: "missing trace id",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.CompanyService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
				c.Request = httpRequest

				return c, rr, nil
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"msg":"Internal Server Error"}`,
		}, {
			name: "success",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.CompanyService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", nil)
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middlewear.TraceIdKey, "123")
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest

				mc := gomock.NewController(t)
				ms := services.NewMockCompanyService(mc)

				ms.EXPECT().GetAllCompanies().Return([]model.Company{}, nil).AnyTimes()

				return c, rr, ms
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   "[]",
		},
		{
			name: "failure to get all company",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.CompanyService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", nil)
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middlewear.TraceIdKey, "123")
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest

				mc := gomock.NewController(t)
				ms := services.NewMockCompanyService(mc)

				ms.EXPECT().GetAllCompanies().Return(nil, errors.New("errors")).AnyTimes()

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
				cs: ms,
			}
			h.getAllCompany(c)
			assert.Equal(t, tt.expectedStatusCode, rr.Code)
			assert.Equal(t, tt.expectedResponse, rr.Body.String())
		})
	}
}

func Test_handler_getCompany(t *testing.T) {

	tests := []struct {
		name               string
		h                  *handler
		setup              func() (*gin.Context, *httptest.ResponseRecorder, services.CompanyService)
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name: "missing trace id",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.CompanyService) {
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
			name: "error while fetching companies from service",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.CompanyService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", nil)
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middlewear.TraceIdKey, "123")
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest
				c.Params = append(c.Params, gin.Param{Key: "company_id", Value: "abc"})
				mc := gomock.NewController(t)
				ms := services.NewMockCompanyService(mc)

				ms.EXPECT().GetCompany(gomock.Any()).Return(model.Company{}, errors.New("test service error")).AnyTimes()

				return c, rr, ms
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   `{"ID":0,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"company_name":"","company_adress":"","domain":""}`,
		},
		{
			name: "success",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.CompanyService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", nil)
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middlewear.TraceIdKey, "123")

				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest
				c.Params = append(c.Params, gin.Param{Key: "company_id", Value: "2"})
				mc := gomock.NewController(t)
				ms := services.NewMockCompanyService(mc)

				ms.EXPECT().GetCompany(gomock.Any()).Return(model.Company{}, nil).AnyTimes()

				return c, rr, ms
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   `{"ID":0,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"company_name":"","company_adress":"","domain":""}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			c, rr, ms := tt.setup()
			h := &handler{
				cs: ms,
			}
			h.getCompany(c)
			assert.Equal(t, tt.expectedStatusCode, rr.Code)
			assert.Equal(t, tt.expectedResponse, rr.Body.String())

		})
	}
}

func Test_handler_getJob(t *testing.T) {

	tests := []struct {
		name               string
		setup              func() (*gin.Context, *httptest.ResponseRecorder, services.CompanyService)
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name: "missing trace id",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.CompanyService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", strings.NewReader(`"error"`))
				c.Request = httpRequest

				return c, rr, nil
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"msg":"Bad Request"}`,
		},
		{
			name: "success",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.CompanyService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", nil)
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middlewear.TraceIdKey, "123")

				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest
				c.Params = append(c.Params, gin.Param{Key: "company_id", Value: "2"})
				mc := gomock.NewController(t)
				ms := services.NewMockCompanyService(mc)

				ms.EXPECT().GetJobs(gomock.Any()).Return([]model.Job{}, nil).AnyTimes()

				return c, rr, ms
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   `[]`,
		},
		{
			name: "error in trace id",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.CompanyService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", strings.NewReader("error"))
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middlewear.TraceIdKey, nil)

				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest
				c.Params = append(c.Params, gin.Param{Key: "company_id", Value: "2"})
				mc := gomock.NewController(t)
				ms := services.NewMockCompanyService(mc)

				ms.EXPECT().GetJobs(gomock.Any()).Return(nil, errors.New("error in trace id")).AnyTimes()

				return c, rr, ms
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"msg":"Internal Server Error"}`,
		},
		{
			name: "failure case 1",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.CompanyService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", nil)
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middlewear.TraceIdKey, "123")

				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest
				c.Params = append(c.Params, gin.Param{Key: "company_id", Value: "2"})
				mc := gomock.NewController(t)
				ms := services.NewMockCompanyService(mc)

				ms.EXPECT().GetJobs(gomock.Any()).Return(nil, errors.New("error")).AnyTimes()

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
				cs: ms,
			}
			h.getJob(c)
			assert.Equal(t, tt.expectedStatusCode, rr.Code)
			assert.Equal(t, tt.expectedResponse, rr.Body.String())

		})
	}
}

func Test_handler_getAllJob(t *testing.T) {

	tests := []struct {
		name               string
		setup              func() (*gin.Context, *httptest.ResponseRecorder, services.CompanyService)
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name: "missing trace id",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.CompanyService) {
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
			name: "success",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.CompanyService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", nil)
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middlewear.TraceIdKey, "123")
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest

				mc := gomock.NewController(t)
				ms := services.NewMockCompanyService(mc)

				ms.EXPECT().GetAllJobs().Return([]model.Job{}, nil).AnyTimes()

				return c, rr, ms
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   "[]",
		},
		{
			name: "failure case",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.CompanyService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", nil)
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middlewear.TraceIdKey, "123")
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest

				mc := gomock.NewController(t)
				ms := services.NewMockCompanyService(mc)

				ms.EXPECT().GetAllJobs().Return(nil, errors.New("error")).AnyTimes()

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
				cs: ms,
			}
			h.getAllJob(c)
			assert.Equal(t, tt.expectedStatusCode, rr.Code)
			assert.Equal(t, tt.expectedResponse, rr.Body.String())
		})
	}
}

func Test_handler_postJobByCompany(t *testing.T) {

	tests := []struct {
		name               string
		setup              func() (*gin.Context, *httptest.ResponseRecorder, services.CompanyService)
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name: "missing trace id",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.CompanyService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
				c.Request = httpRequest

				return c, rr, nil
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"error":"Internal Server Error"}`,
		},
		{
			name: "missing param",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.CompanyService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middlewear.TraceIdKey, "123")
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest

				return c, rr, nil
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"error":"Bad Request"}`,
		},
		{name: "input id validation failure",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.CompanyService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", strings.NewReader(`{"job_title":"",
		"job_salary":    "",}`))
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middlewear.TraceIdKey, "123")
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest

				return c, rr, nil
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"error":"Bad Request"}`,
		},
		{name: "input validation failure",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.CompanyService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", strings.NewReader(`{"job_title":"",
		"job_salary":    "",}`))
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middlewear.TraceIdKey, nil)
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest

				return c, rr, nil
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"error":"Internal Server Error"}`,
		},
		// {name: "success",
		// 	setup: func() (*gin.Context, *httptest.ResponseRecorder, services.CompanyService) {
		// 		rr := httptest.NewRecorder()
		// 		c, _ := gin.CreateTestContext(rr)
		// 		httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
		// 		ctx := httpRequest.Context()
		// 		ctx = context.WithValue(ctx, middlewear.TraceIdKey, nil)
		// 		httpRequest = httpRequest.WithContext(ctx)
		// 		c.Request = httpRequest
		// 		c.Params = append(c.Params, gin.Param{Key: "company_id"})
		// 		mc := gomock.NewController(t)
		// 		ms := services.NewMockCompanyService(mc)

		// 		ms.EXPECT().JobCreate(gomock.Any(), gomock.Any()).Return(model.Response{}, nil).AnyTimes()

		// 		return c, rr, ms
		// 	},
		// 	expectedStatusCode: http.StatusOK,
		// 	expectedResponse:   `[]`,
		// },
		{
			name: "failure case",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.CompanyService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", strings.NewReader(`{"job_title":"rvrvrv",
					"job_salary":    "16786"}`))
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middlewear.TraceIdKey, "123")
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest

				c.Params = append(c.Params, gin.Param{Key: "company_id", Value: "2"})
				mc := gomock.NewController(t)
				ms := services.NewMockCompanyService(mc)

				ms.EXPECT().JobCreate(gomock.Any(), gomock.Any()).Return(model.Response{}, errors.New("error")).AnyTimes()

				return c, rr, ms
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"msg":"Internal Server Error"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			c, rr, ms := tt.setup()
			h := &handler{
				cs: ms,
			}
			h.postJobByCompany(c)
			assert.Equal(t, tt.expectedStatusCode, rr.Code)
			assert.Equal(t, tt.expectedResponse, rr.Body.String())
		})
	}
}

// func Test_handler_processApplications(t *testing.T) {

// 	tests := []struct {
// 		name               string
// 		setup              func() (*gin.Context, *httptest.ResponseRecorder, services.CompanyService)
// 		expectedStatusCode int
// 		expectedResponse   string
// 	}{
// 		{
// 			name: "missing trace id",
// 			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.CompanyService) {
// 				rr := httptest.NewRecorder()
// 				c, _ := gin.CreateTestContext(rr)
// 				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
// 				c.Request = httpRequest

// 				return c, rr, nil
// 			},
// 			expectedStatusCode: http.StatusInternalServerError,
// 			expectedResponse:   `{"msg":"Internal Server Error"}`,
// 		},
// 		{
// 			name: "failure case in json data",
// 			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.CompanyService) {
// 				rr := httptest.NewRecorder()
// 				c, _ := gin.CreateTestContext(rr)
// 				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", strings.NewReader(`{"job_title":"rvrvrv",
// 					"job_salary":    "16786"}`))
// 				ctx := httpRequest.Context()
// 				ctx = context.WithValue(ctx, middlewear.TraceIdKey, "123")
// 				httpRequest = httpRequest.WithContext(ctx)
// 				c.Request = httpRequest

// 				c.Params = append(c.Params, gin.Param{Key: "company_id", Value: "2"})
// 				mc := gomock.NewController(t)
// 				ms := services.NewMockCompanyService(mc)

// 				ms.EXPECT().ProcessJobApplications(gomock.Any()).Return(nil, errors.New("error")).AnyTimes()

// 				return c, rr, ms
// 			},
// 			expectedStatusCode: http.StatusBadRequest,
// 			expectedResponse:   `{"error":"please provide proper data"}`,
// 		},

// 		{
// 			name: "success",
// 			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.CompanyService) {
// 				rr := httptest.NewRecorder()
// 				c, _ := gin.CreateTestContext(rr)
// 				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", strings.NewReader(`{"job_title":"rvrvrv",
// 				"job_salary":    "16786"}`))
// 				ctx := httpRequest.Context()
// 				ctx = context.WithValue(ctx, middlewear.TraceIdKey, "123")
// 				httpRequest = httpRequest.WithContext(ctx)
// 				c.Request = httpRequest

// 				mc := gomock.NewController(t)
// 				ms := services.NewMockCompanyService(mc)

// 				ms.EXPECT().ProcessJobApplications(gomock.Any()).Return(nil, errors.New("errors")).AnyTimes()

// 				return c, rr, ms
// 			},
// 			expectedStatusCode: http.StatusBadRequest,
// 			expectedResponse:   `{"error":"please provide proper data"}`,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			gin.SetMode(gin.TestMode)
// 			c, rr, ms := tt.setup()
// 			h := &handler{
// 				cs: ms,
// 			}
// 			h.processApplications(c)
// 			assert.Equal(t, tt.expectedStatusCode, rr.Code)
// 			assert.Equal(t, tt.expectedResponse, rr.Body.String())
// 		})
// 	}
// }
