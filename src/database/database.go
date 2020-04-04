package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"os"
	"strings"
)

var DB *gorm.DB

func Args() string {
	environmentVariables := []string{"HOST", "PORT", "DBNAME", "USER", "PASSWORD", "SSL_MODE"}
	var stringBuilder strings.Builder
	for _, key := range environmentVariables {
		envVar, hasEnvVar := os.LookupEnv(fmt.Sprintf("DATABASE_%s", key))
		if !hasEnvVar || len(envVar) == 0 {
			missingEnv(key)
		} else {
			stringBuilder.WriteString(
				fmt.Sprintf("%s=%s ", strings.ReplaceAll(strings.ToLower(key), "_", ""), envVar))
		}
	}
	return strings.TrimSpace(stringBuilder.String())
}

func Dialect() string {
	dialectKey := "DATABASE_DIALECT"
	dialectVar, hasDialectVar := os.LookupEnv(dialectKey)
	if !hasDialectVar || len(dialectVar) == 0 {
		missingEnv(dialectKey)
	}
	return dialectVar
}

func LogMode() bool {
	logModeKey := "DATABASE_LOG_MODE"
	logModeVar, hasLogModeVar := os.LookupEnv(logModeKey)
	return hasLogModeVar && logModeVar == "true"
}

func missingEnv(variable string) {
	log.Fatalln(fmt.Sprintf("Environment variable %s is missing", variable))
}
