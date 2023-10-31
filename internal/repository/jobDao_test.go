package repository

import (
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	model "job-portal-api/internal/models"
	"testing"
)

type TestRepo struct {
	companyStore map[string]model.Company
}

func (tr *TestRepo) CreateCompany(company model.Company) (model.Company, error) {
	tr.companyStore[company.CompanyName] = company
	return company, nil
}

func TestCreateCompany(t *testing.T) {
	dsn := "host=localhost user=postgres password=admin dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error().Msg("error in connecting to the database")
		t.Fatal(err) // Terminate the test if the database connection fails
	}

	testRepo := &TestRepo{
		companyStore: make(map[string]model.Company),
	}

	company := model.Company{
		CompanyName: "company1",
		Adress:      "address1",
		Domain:      "domain1",
	}

	createdCompany, err := testRepo.CreateCompany(company)

	assert.NoError(t, err, "Expected no error")

	assert.Equal(t, company.CompanyName, createdCompany.CompanyName, "Company names should match")
	assert.Equal(t, company.Adress, createdCompany.Adress, "Addresses should match")
	assert.Equal(t, company.Domain, createdCompany.Domain, "Domains should match")

	retrievedCompany, found := testRepo.companyStore[company.CompanyName]
	if !found {
		t.Fatal("Company not found in the test repository")
	}
	assert.Equal(t, company.CompanyName, retrievedCompany.CompanyName, "Company names should match")
	assert.Equal(t, company.Adress, retrievedCompany.Adress, "Addresses should match")
	assert.Equal(t, company.Domain, retrievedCompany.Domain, "Domains should match")
}
