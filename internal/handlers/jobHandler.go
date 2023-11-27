package handlers

import (
	"encoding/json"

	//"errors"
	//"job-portal-api/internal/handlers"
	//	"job-portal-api/internal/handlers"
	middlewear "job-portal-api/internal/middleware"
	model "job-portal-api/internal/models"

	//"job-portal-api/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

// =============== COMPANY CREATION FUNC do's all the creations of company in the database ============
func (h *handler) companyCreation(c *gin.Context) {

	ctx := c.Request.Context()

	traceId, ok := ctx.Value(middlewear.TraceIdKey).(string)

	if !ok {
		log.Error().Str("traceId", traceId).Msg("trace id not found in userSignin handler")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return

	}

	var companyCreation model.CreateCompany
	body := c.Request.Body
	err := json.NewDecoder(body).Decode(&companyCreation)
	if err != nil {
		log.Error().Err(err).Msg("error in decoding")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}

	validate := validator.New()
	err = validate.Struct(&companyCreation)
	if err != nil {
		log.Error().Err(err).Msg("error in validating ")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": "invalid input"})
		return
	}

	us, err := h.cs.CompanyCreate(companyCreation)
	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId).Msg("user signup problem")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "user signup failed"})
		return
	}
	c.JSON(http.StatusOK, us)

}

// ======================GET ALL COMPANY FUNC will retrive all companies===============
func (h *handler) getAllCompany(c *gin.Context) {
	ctx := c.Request.Context()

	traceId, ok := ctx.Value(middlewear.TraceIdKey).(string)

	if !ok {
		log.Error().Str("traceId", traceId).Msg("trace id not found in userSignin handler")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}

	us, err := h.cs.GetAllCompanies()
	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId).Msg("user signup problem")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "user signup failed"})
		return
	}
	c.JSON(http.StatusOK, us)

}

// ====================GET COMPANY FUNC RETRIVE COMPANY BY ID IN THE DATABASE===============================
func (h *handler) getCompany(c *gin.Context) {

	ctx := c.Request.Context()
	traceId, ok := ctx.Value(middlewear.TraceIdKey).(string)
	id, _ := strconv.ParseInt(c.Param("id"), 10, 32)
	if !ok {
		log.Error().Str("traceId", traceId).Msg("trace id not found in userSignin handler")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}
	co, _ := h.cs.GetCompany(id)

	c.JSON(http.StatusOK, co)

}

// ========================== POST JOB By COMPANY FUNC DO'S CREATING JOB IN THE COMPANY =================================

func (h *handler) postJobByCompany(c *gin.Context) {
	ctx := c.Request.Context()
	traceId, ok := ctx.Value(middlewear.TraceIdKey).(string)
	if !ok {
		log.Error().Str("traceId", traceId).Msg("trace id not found in  handler")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	id, erro := strconv.ParseUint(c.Param("company_id"), 10, 32)
	if erro != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": http.StatusText(http.StatusBadRequest)})
		return
	}
	var jobCreation model.NewJobRequest
	body := c.Request.Body
	err := json.NewDecoder(body).Decode(&jobCreation)
	if err != nil {
		log.Error().Err(err).Msg("error in decoding")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	validate := validator.New()
	err = validate.Struct(&jobCreation)
	if err != nil {
		log.Error().Err(err).Msg("error in validating ")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}
	us, err := h.cs.JobCreate(jobCreation, id)

	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId).Msg("job creatuion problem in db")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}
	c.JSON(http.StatusOK, us)

}

// ========================== GET JOB FUNC RETRIVE ALL JOB IN THE DATABASE ===========================
func (h *handler) getJob(c *gin.Context) {
	ctx := c.Request.Context()
	traceId, ok := ctx.Value(middlewear.TraceIdKey).(string)
	id, erro := strconv.Atoi(c.Param("company_id"))
	if erro != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": http.StatusText(http.StatusBadRequest)})
		return
	}

	if !ok {
		log.Error().Str("traceId", traceId).Msg("trace id not found in userSignin handler")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}
	us, err := h.cs.GetJobs(id)
	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId).Msg("user signup problem")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "user signup failed"})
		return
	}
	c.JSON(http.StatusOK, us)

}

// ======================= GET ALL JOB FUNC RETRIVE ALL JOB IN THE ALL COMPANY ====================
func (h *handler) getAllJob(c *gin.Context) {
	ctx := c.Request.Context()
	traceId, ok := ctx.Value(middlewear.TraceIdKey).(string)
	if !ok {
		log.Error().Str("traceId", traceId).Msg("trace id not found in userSignin handler")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}

	us, err := h.cs.GetAllJobs()
	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId).Msg("user signup problem")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "user signup failed"})
		return
	}
	c.JSON(http.StatusOK, us)

}

func (h *handler) ProcessApplication(c *gin.Context) {
	ctx := c.Request.Context()
	traceid, ok := ctx.Value(middlewear.TraceIdKey).(string)
	if !ok {
		log.Error().Msg("traceid is not found")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	//_, ok = ctx.Value(auth.Key).(jwt.RegisteredClaims)
	// if !ok {
	// 	log.Error().Str("Trace Id", traceid).Msg("login first")
	// 	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
	// 	return
	// }
	var applicationData []model.NewUserApplication
	err := json.NewDecoder(c.Request.Body).Decode(&applicationData)
	if err != nil {
		log.Error().Err(err).Str("trace id", traceid)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "please provide valid details",
		})
		return
	}
	applicationData, err = h.cs.ProccessApplication(ctx, applicationData)
	//fmt.Println("[][][]", applicationData)
	if err != nil {
		log.Error().Err(err).Str("trace id", traceid)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, applicationData)

}
