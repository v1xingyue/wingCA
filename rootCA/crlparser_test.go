package rootCA

import (
	"net"
	"testing"
)

func TestParse(t *testing.T) {
	email := "a@info.com.cn"
	commonName := "namedCSRRequest"
	nameList := []string{"a.info.com.cn", "b.info.com.cn", "localhost"}
	addrList := []net.IP{net.ParseIP("127.0.0.1")}
	buffer := makeCSR(email, commonName, nameList, addrList)
	// fmt.Println(string(x))
	parseCSRContent(buffer)
	t.Error(".")
}
