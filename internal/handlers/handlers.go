package handlers

import (
	"job-portal-api/internal/auth"
	"job-portal-api/internal/middleware"
	"job-portal-api/internal/services"

	"github.com/gin-gonic/gin"
)

// =================api func contains all methods and ends points of the api=================
func Api(a *auth.Auth, s *services.Service) *gin.Engine {
	router := gin.New()
	h, _ := NewHandler(a, s, s)

	m, _ := middlewear.NewMiddleWear(a)
	router.Use(m.Log(), gin.Recovery())
	router.POST("/signin", (h.userSignin))
	router.POST("/login", h.userLoginin)
	router.POST("/createCompany", m.Auth(h.companyCreation))
	router.GET("/getAllCompany", m.Auth(h.getAllCompany))
	router.GET("/getCompany/:id", m.Auth(h.getCompany))
	router.POST("/api/companies/:company_id/jobs", m.Auth(h.postJobByCompany))
	router.GET("/companies/:company_id/viewjobs", m.Auth(h.getJob))
	router.GET("/viewAllJobs", m.Auth(h.getAllJob))
	router.POST("/api/applications", h.ProcessApplication)
	router.POST("/forget/login/password", h.ForgetPassword)
	router.POST("/changePassword/otp", h.ChangePassword)

	return router
}
