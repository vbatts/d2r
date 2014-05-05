package sum

import (
	"github.com/dotcloud/docker/utils"
	"io"
	"io/ioutil"
)

// if out is not nil, then the tar input is written there instead
func SumTarLayer(tar io.Reader, json io.Reader, out io.Writer) (string, error) {
	var writer io.Writer = ioutil.Discard
	if out != nil {
		writer = out
	}
	ts := &utils.TarSum{Reader: tar}
	_, err := io.Copy(writer, ts)
	if err != nil {
		return "", err
	}

	buf, err := ioutil.ReadAll(json)
	if err != nil {
		return "", err
	}

	return ts.Sum(buf), nil
}
