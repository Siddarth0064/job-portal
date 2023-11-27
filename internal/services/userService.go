package services

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"net/smtp"

	"job-portal-api/internal/cache"
	model "job-portal-api/internal/models"
	"job-portal-api/internal/repository"

	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	r   repository.Users
	c   repository.Company
	rdb cache.Caching
}

var otps string

// ===================== NEW SERVICE FUNC IS USED TO INITIALIZE TO SERVICE STRUCT =====================
func NewService(r repository.Users, c repository.Company, rdb cache.Caching) (*Service, error) {
	if r == nil {
		return nil, errors.New("db connection not given")
	}

	return &Service{
		r:   r,
		c:   c,
		rdb: rdb,
	}, nil

}

// ==================== USERS SERVICE INTERFACE ===========================================
//
//go:generate mockgen -source=userService.go -destination=userService_mock.go -package=services
type UsersService interface {
	UserSignup(nu model.UserSignup) (model.User, error)
	Userlogin(l model.UserLogin) (jwt.RegisteredClaims, error)
	ForgetPassword(ctx context.Context, f model.ForgetPass) (bool, string, error)
	ChangePass(ctx context.Context, c model.ChnagePass) (string, error)
}

// ======================== USER SIGNUP FUNC ==================================================
func (s *Service) UserSignup(nu model.UserSignup) (model.User, error) {

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(nu.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error().Msg("error occured in hashing password")
		return model.User{}, errors.New("hashing password failed")
	}

	user := model.User{UserName: nu.UserName, Email: nu.Email, PasswordHash: string(hashedPass), DateOfBorn: nu.DateOfBorn}
	// database.CreateTable()
	//fmt.Println(nu.DateOfBorn)
	cu, err := s.r.CreateUser(user)
	if err != nil {
		log.Error().Err(err).Msg("couldnot create user")
		return model.User{}, errors.New("user creation failed")
	}

	return cu, nil

}

// ================ USER LOG IN FUNC =============================================================
func (s *Service) Userlogin(l model.UserLogin) (jwt.RegisteredClaims, error) {
	fu, err := s.r.FetchUserByEmail(l.Email)
	if err != nil {
		log.Error().Err(err).Msg("couldnot find user")
		return jwt.RegisteredClaims{}, errors.New("user login failed")
	}

	err = bcrypt.CompareHashAndPassword([]byte(fu.PasswordHash), []byte(l.Password))
	if err != nil {
		log.Error().Err(err).Msg("password of user incorrect")
		return jwt.RegisteredClaims{}, errors.New("user login failed")
	}
	c := jwt.RegisteredClaims{
		Issuer:    "service project",
		Subject:   strconv.FormatUint(uint64(fu.ID), 10),
		Audience:  jwt.ClaimStrings{"users"},
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	return c, nil

}

//==========================================================================

func (s *Service) ForgetPassword(ctx context.Context, l model.ForgetPass) (bool, string, error) {
	fmt.Println(l.Email)
	fmt.Println(l.DateOfBorn)
	ue, err := s.r.FetchUserByEmail(l.Email)
	fmt.Println(ue, "user mail fatch by server func")
	if err != nil {
		log.Error().Err(err).Msg("couldnot find user emaiil in server")
		return false, "", errors.New("user email failed")
	}
	// err = s.r.FetchUserByDob(l.DateOfBorn)
	// if err != nil {
	// 	log.Error().Err(err).Msg("couldnot find user dob")
	// 	return errors.New("user DOB failed")
	// }
	otps = generateRandomOTP(5)

	from := "mpsiddarthgowda@gmail.com"
	password := "kkfp vxwc lsgo qzsf"
	// Recipient's email address
	to := ue
	// SMTP server details
	smtpServer := "smtp.gmail.com"
	smtpPort := 587
	// Message content
	message := []byte(fmt.Sprintf("Subject: Test Email\n\nThis is a test email body.", (otps)))
	// Authentication information
	auth := smtp.PlainAuth("", from, password, smtpServer)
	// SMTP connection
	smtpAddr := fmt.Sprintf("%s:%d", smtpServer, smtpPort)
	err = smtp.SendMail(smtpAddr, auth, from, []string{to.Email}, message)
	if err != nil {
		fmt.Println("Error sending email:", err)
		return false, "", errors.New("error in sending otp to email")
	}
	fmt.Println("Email sent successfully!")
	err = s.rdb.AddtocacheOTP(ctx, l.Email, otps)
	// if err != nil {
	// 	fmt.Println("there is an error in adding cache ========================================================")
	// 	return false, "", errors.New("otp doesnot add to cache")
	// }

	return true, otps, nil

}
func generateRandomOTP(length int) string {
	rand.Seed(time.Now().UnixNano())

	// Define the characters allowed in the OTP
	otpChars := "0123456789abcdefghijklmnopqrstABCDEFGHIHIJKLMNOPQRSTUVWXYZ"

	// Generate the OTP
	otp := make([]byte, length)
	for i := range otp {
		otp[i] = otpChars[rand.Intn(len(otpChars))]
	}

	return string(otp)
}

//============================== CHANGING THE PASSWORD ============================================

func (s *Service) ChangePass(ctx context.Context, cc model.ChnagePass) (string, error) {
	otpData, _ := s.rdb.GetCacheDataOTP(ctx, cc.Email)
	// if err != nil {
	// 	return "", errors.New("error in geting cache opt data in cache")
	// }

	if cc.Otp == otpData {
		if cc.NewPassword == cc.ComfirmPassword {
			newpassuser, err := s.r.FetchUserByEmail(cc.Email)
			if err != nil {
				return "", errors.New("error in fetching data by email in server layer")
			}
			hashedPass, err := bcrypt.GenerateFromPassword([]byte(cc.ComfirmPassword), bcrypt.DefaultCost)
			if err != nil {
				return "", errors.New("error in hassing password for new password")
			}
			user := model.User{UserName: newpassuser.UserName, DateOfBorn: newpassuser.DateOfBorn, Email: newpassuser.Email, PasswordHash: string(hashedPass)}

			_, err = s.r.UpdateUser(cc.Email, user)
			if err != nil {
				log.Error().Err(err).Msg("couldnot update user")
				return "", errors.New("user updating failed")
			}

		} else {
			return "", errors.New("password is mismatch")
		}
	} else {
		return "", errors.New("otp is not valid")
	}
	return "SUCCESSFULLY NEW PASSWORD IS SET", nil

}
