package handlers

import (
	"github.com/gin-gonic/gin"
	"job-portal-api/internal/auth"
	"job-portal-api/internal/middleware"
	"job-portal-api/internal/services"
)

func Api(a *auth.Auth, s *services.Service) *gin.Engine {
	r := gin.New()
	h, _ := NewHandler(a, s, s)
	//h := handler{a: a, us: &s}
	m, _ := middlewear.NewMiddleWear(a)
	r.Use(m.Log(), gin.Recovery())
	r.POST("/signup", (h.userSignin))
	r.POST("/login", h.userLoginin)
	r.POST("/createCompany", m.Auth(h.companyCreation))
	r.GET("/getAllCompany", m.Auth(h.getAllCompany))
	r.GET("/getCompany/:id", m.Auth(h.getCompany))
	r.POST("/companies/:company_id/jobs", m.Auth(h.postJob))
	r.GET("/companies/:company_id/jobs", m.Auth(h.getJob))
	r.GET("/jobs", m.Auth(h.getAllJob))
	return r
}
