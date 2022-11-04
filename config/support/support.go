package support

import (
	"os"
	"runtime"
	"strings"
)

var cryptoRoot string
var externalDir string

func init() {

	var goPath string
	// windows has ; separator vs linux has :
	if runtime.GOOS == "windows" {
		goPath = strings.Split(os.Getenv("GOPATH"), ";")[0]
	} else {
		goPath = strings.Split(os.Getenv("GOPATH"), ":")[0]
	}

	// Derive the app root directory
	cryptoRoot = goPath + "/src/crypto_price_tracker"
}

func GetCryptoRootDir() string {
	return cryptoRoot
}
