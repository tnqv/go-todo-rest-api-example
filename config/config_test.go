package config

import (
	"os"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/kelseyhightower/envconfig"
)

func TestConfig(t *testing.T) {
	var s DBConfig
	os.Clearenv()
	os.Setenv("DB_HOST", "HOST")
	os.Setenv("DB_USERNAME", "USERNAME")
	os.Setenv("DB_PASSWORD", "PASSWROD")
	os.Setenv("DB_NAME", "DBNAME")

	err := envconfig.Process("db", &s)

	if err != nil {
		t.Error(err.Error())
	}

	assert := assert.New(t)
	assert.Equal(s.Host, "HOST", "they should be equal")
	assert.Equal(s.Username, "USERNAME", "they should be equal")
	assert.Equal(s.Password, "PASSWROD", "they should be equal")
	assert.Equal(s.Name, "DBNAME", "they should be equal")
	assert.Equal(s.Port, 3306, "they should be equal")
}
