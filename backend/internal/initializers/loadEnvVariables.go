package initializers

import (
	"github.com/joho/godotenv"
)

func LoadEnvVariables() error {
	// Load environment variables
	err := godotenv.Load("/Users/muh/Developer/projects/midwifery-ehr/.env")
	if err != nil {
		return err
	}
	return nil
}
