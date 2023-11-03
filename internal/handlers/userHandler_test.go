package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	//"github.com/stretchr/testify/assert"
	"github.com/go-playground/assert/v2"
)

func Test_handler_userSignin(t *testing.T) {
	// type args struct {
	// 	c *gin.Context
	// }
	tests := []struct {
		name string
		h    *handler
		// args args
		setup              func() (*gin.Context, *httptest.ResponseRecorder)
		expectedStatusCode int
		expectedResponse   string
	}{
		{name: "missing trace id",
			setup: func() (*gin.Context, *httptest.ResponseRecorder) {
				rr := httptest.NewRecorder()      // responseRecorder
				c, _ := gin.CreateTestContext(rr) // context
				return c, rr
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `"error":"Internal Servaer Error"`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, rr := tt.setup()
			h := handler.New()
			tt.h.userSignin(c)
			assert.Equal(t, tt.expectedStatusCode, rr.Code)
			assert.Equal(t, tt.expectedResponse, rr.Body.String())
		})
	}
}
