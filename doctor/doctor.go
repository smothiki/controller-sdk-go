package doctor

import (
	"io/ioutil"
	"strconv"

	deis "github.com/deis/controller-sdk-go"
)

// Send doctor info to deis servers.
func Send(dc *deis.Client) (string, error) {
	path := "/doctor"
	b := []byte(nil)
	b = strconv.AppendBool(b, dc.VerifySSL)
	res, err := dc.Request("POST", path, b)
	if err == nil {
		resbody, err1 := ioutil.ReadAll(res.Body)
		return string(resbody), err1
	}
	return "", err

}
