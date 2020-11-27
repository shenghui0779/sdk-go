package mp

import (
	"testing"

	"github.com/shenghui0779/gochat/helpers"
)

var postBody helpers.HTTPBody

func TestMain(m *testing.M) {
	postBody = helpers.NewPostBody(func() ([]byte, error) {
		return nil, nil
	})

	m.Run()
}
