package vo

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// PasswordVoTestSuite estrutura para os testes de PasswordVo
type PasswordVoTestSuite struct {
	suite.Suite
}

// TestValidPassword verifica se uma senha válida é tratada corretamente
func (suite *PasswordVoTestSuite) TestValidPassword() {
	password := "Valid1@Password"
	pv, err := NewPasswordVo(password)
	assert.Nil(suite.T(), err, "erro deve ser nil para uma senha válida")
	assert.Equal(suite.T(), password, pv.Value(), "o valor da senha deve ser igual ao valor fornecido")
}

// TestPasswordTooShort verifica se uma senha muito curta retorna erro
func (suite *PasswordVoTestSuite) TestPasswordTooShort() {
	password := "Short1@"
	pv, err := NewPasswordVo(password)
	assert.NotNil(suite.T(), err, "erro deve ser retornado para uma senha muito curta")
	assert.Equal(suite.T(), ErrPasswordIsTooShort, err)
	assert.Nil(suite.T(), pv, "PasswordVo deve ser nil para uma senha muito curta")
}

// TestPasswordNoUppercase verifica se uma senha sem letras maiúsculas retorna erro
func (suite *PasswordVoTestSuite) TestPasswordNoUppercase() {
	password := "lowercase1@"
	pv, err := NewPasswordVo(password)
	assert.NotNil(suite.T(), err, "erro deve ser retornado para uma senha sem letras maiúsculas")
	assert.Equal(suite.T(), ErrPasswordLowercaseLetter, err)
	assert.Nil(suite.T(), pv, "PasswordVo deve ser nil para uma senha sem letras maiúsculas")
}

// TestPasswordNoLowercase verifica se uma senha sem letras minúsculas retorna erro
func (suite *PasswordVoTestSuite) TestPasswordNoLowercase() {
	password := "UPPERCASE1@"
	pv, err := NewPasswordVo(password)
	assert.NotNil(suite.T(), err, "erro deve ser retornado para uma senha sem letras minúsculas")
	assert.Equal(suite.T(), ErrPasswordUppercaseLetter, err)
	assert.Nil(suite.T(), pv, "PasswordVo deve ser nil para uma senha sem letras minúsculas")
}

// TestPasswordNoDigit verifica se uma senha sem dígitos retorna erro
func (suite *PasswordVoTestSuite) TestPasswordNoDigit() {
	password := "NoDigit@Password"
	pv, err := NewPasswordVo(password)
	assert.NotNil(suite.T(), err, "erro deve ser retornado para uma senha sem dígitos")
	assert.Equal(suite.T(), ErrPasswordAtLeastOneDigit, err)
	assert.Nil(suite.T(), pv, "PasswordVo deve ser nil para uma senha sem dígitos")
}

// TestPasswordNoSpecialChar verifica se uma senha sem caracteres especiais retorna erro
func (suite *PasswordVoTestSuite) TestPasswordNoSpecialChar() {
	password := "NoSpecialChar1"
	pv, err := NewPasswordVo(password)
	assert.NotNil(suite.T(), err, "erro deve ser retornado para uma senha sem caracteres especiais")
	assert.Equal(suite.T(), ErrPasswordAtLeastSpecialChar, err)
	assert.Nil(suite.T(), pv, "PasswordVo deve ser nil para uma senha sem caracteres especiais")
}

// TestPasswordVoTestSuite inicia o suite de teste
func TestPasswordVoTestSuite(t *testing.T) {
	suite.Run(t, new(PasswordVoTestSuite))
}
