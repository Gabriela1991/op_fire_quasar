package env

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

var PORT_SERVER string
var KEY_AES_SODIMAC string

var (
	// Get current file full path from runtime
	_, b, _, _ = runtime.Caller(0)

	// Root folder of this project
	ProjectRootPath = filepath.Join(filepath.Dir(b), "../../")
)

func init() {
	environment := os.Getenv("ENV")
	if environment == "PROD" {
		_ = godotenv.Load(ProjectRootPath + "/environment_prod.env")
	} else {
		_ = godotenv.Load(ProjectRootPath + "/environment_test.env")
	}

	PORT_SERVER = os.Getenv("PORT_SERVER")

}
