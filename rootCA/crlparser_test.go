package rootCA

import (
	"io/ioutil"
	"net"
	"testing"
)

func TestParse(t *testing.T) {
	email := "a@info.com.cn"
	commonName := "namedCSRRequest"
	nameList := []string{"a.info.com.cn", "b.info.com.cn", "localhost"}
	addrList := []net.IP{net.ParseIP("127.0.0.1")}
	buffer, keyBytes := makeCSR(email, commonName, nameList, addrList)
	// fmt.Println(string(x))
	parseCSRContent(buffer)
	ioutil.WriteFile(commonName+".key", keyBytes, 0600)
	t.Error(".")
}
