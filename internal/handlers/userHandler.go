package handlers

import (
	"encoding/json"
	"errors"
	"job-portal-api/internal/auth"
	"job-portal-api/internal/middleware"
	"job-portal-api/internal/models"
	"job-portal-api/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

type handler struct {
	a  *auth.Auth
	us services.UsersService
	cs services.CompanyService
}

// =========================NEW HANDLERS IS INITIALIZE TO THE HANDLER STRUCT =========================
func NewHandler(a *auth.Auth, us services.UsersService, cs services.CompanyService) (*handler, error) {
	if us == nil {
		return nil, errors.New("service implementation not given")
	}

	return &handler{a: a, us: us, cs: cs}, nil

}

// ===================== USER SIGN IN FUNC IS USED TO SIGNIN TO THE DATABASE =============================
func (h *handler) userSignin(c *gin.Context) {
	ctx := c.Request.Context()

	traceId, ok := ctx.Value(middlewear.TraceIdKey).(string)
	if !ok {
		log.Error().Str("traceId", traceId).Msg("trace id not found in userSignin handler")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}

	var userCreate model.UserSignup
	body := c.Request.Body
	err := json.NewDecoder(body).Decode(&userCreate)
	if err != nil {
		log.Error().Err(err).Msg("error in decoding")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}

	validate := validator.New()
	err = validate.Struct(&userCreate)
	if err != nil {
		log.Error().Err(err).Msg("error in validating ")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": "invalid input"})
		return
	}
	us, err := h.us.UserSignup(userCreate)
	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId).Msg("user signup problem")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "user signup failed"})
		return
	}
	c.JSON(http.StatusOK, us)

}

// ================================== USER LOGIN IN FUNC IS USED TO LOGIN TO THE ACCOUNT ===========================
func (h *handler) userLoginin(c *gin.Context) {
	ctx := c.Request.Context()

	traceId, ok := ctx.Value(middlewear.TraceIdKey).(string)
	if !ok {
		log.Error().Str("traceId", traceId).Msg("trace id not found in userSignin handler")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}
	var userLogin model.UserLogin
	body := c.Request.Body
	err := json.NewDecoder(body).Decode(&userLogin)
	if err != nil {
		log.Error().Err(err).Msg("error in decoding")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}

	validate := validator.New()
	err = validate.Struct(&userLogin)
	if err != nil {
		log.Error().Err(err).Msg("error in validating ")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": "invalid input"})
		return
	}

	regClaims, err := h.us.Userlogin(userLogin)
	if err != nil {
		log.Error().Err(err).Msg("error in Loginin ")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": "invalid input"})
		return
	}

	token, err := h.a.GenerateToken(regClaims)
	if err != nil {
		log.Error().Err(err).Msg("error in Gneerating toek ")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return

	}

	c.JSON(http.StatusOK, token)

}

func (h *handler) ForgetPassword(c *gin.Context) {
	ctx := c.Request.Context()

	traceId, ok := ctx.Value(middlewear.TraceIdKey).(string)
	if !ok {
		log.Error().Str("traceId", traceId).Msg("trace id not found in forgetPassword handler")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}

	var forgetPassData model.ForgetPass
	body := c.Request.Body
	err := json.NewDecoder(body).Decode(&forgetPassData)
	if err != nil {
		log.Error().Err(err).Msg("error in decoding in forgetPassword func")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}

	validate := validator.New()
	err = validate.Struct(&forgetPassData)
	if err != nil {
		log.Error().Err(err).Msg("error in validating ")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": "invalid input pls provide valid struct"})
		return
	}

	_, otph, err := h.us.ForgetPassword(ctx, forgetPassData)
	if err != nil {
		log.Error().Err(err).Msg("error in forgetPassword func  ")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": "can't send password to mail,  error in server"})
		return
	}

	// token, err := h.a.GenerateToken(regClaims)
	// if err != nil {
	// 	log.Error().Err(err).Msg("error in Gneerating toek ")
	// 	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
	// 	return

	// }

	c.JSON(http.StatusOK, otph)

}
func (h *handler) ChangePassword(c *gin.Context) {
	ctx := c.Request.Context()

	traceId, ok := ctx.Value(middlewear.TraceIdKey).(string)
	if !ok {
		log.Error().Str("traceId", traceId).Msg("trace id not found in ChangePassword handler")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}
	var otpChange model.ChnagePass
	body := c.Request.Body
	err := json.NewDecoder(body).Decode(&otpChange)
	if err != nil {
		log.Error().Err(err).Msg("error in decoding")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}

	validate := validator.New()
	err = validate.Struct(&otpChange)
	if err != nil {
		log.Error().Err(err).Msg("error in validating ")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": "invalid input"})
		return
	}

	success, err := h.us.ChangePass(ctx, otpChange)
	if err != nil {
		log.Error().Err(err).Msg("error in Loginin ")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": "error in changePass func"})
		return
	}

	// token, err := h.a.GenerateToken(regClaims)
	// if err != nil {
	// 	log.Error().Err(err).Msg("error in Gneerating toek ")
	// 	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
	// 	return

	// }

	c.JSON(http.StatusOK, success)

}
