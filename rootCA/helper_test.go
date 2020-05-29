package rootCA

import (
	"log"
	"os"
	"testing"
)

func InitEnv() {
	os.Chdir("../")
	d, err := os.Getwd()
	log.Println(d, err)
	//RootCAPath = "../ssl"
}

func TestCertExpired(t *testing.T) {
	InitEnv()
	log.Println(CertExpired("a.b.ssl.com.cn", CertTypeSite))
	t.Error("")
}
