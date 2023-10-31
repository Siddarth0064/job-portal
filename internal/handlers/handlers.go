package handlers

import (
	"github.com/gin-gonic/gin"
	"job-portal-api/internal/auth"
	"job-portal-api/internal/middleware"
	"job-portal-api/internal/services"
)

func Api(a *auth.Auth, s services.Service) *gin.Engine {
	r := gin.New()
	h := handler{a: a, us: s}
	m, _ := middlewear.NewMiddleWear(a)
	r.Use(m.Log(), gin.Recovery())
	r.POST("/signup", h.userSignin)
	r.POST("/login", h.userLoginin)
	r.POST("/createCompany", h.companyCreation)
	r.GET("/getAllCompany", h.getAllCompany)
	r.GET("/getCompany/:id", h.getCompany)
	r.POST("/companies/:company_id/jobs", h.postJob)
	r.GET("/companies/:company_id/jobs", h.getJob)
	r.GET("/jobs", h.getAllJob)
	return r
}
