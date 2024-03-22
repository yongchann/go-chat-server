package redis

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var rc *Client

func init() {
	tc, err := NewClient("localhost:6379", "12341234")
	if err != nil {
		panic(err)
	}

	rc = tc
}

func TestClient_SetAndGetKey(t *testing.T) {
	t.Run("GetKey", func(t *testing.T) {
		if err := rc.SetKey("yong", "chan", time.Second*10); err != nil {
			t.Error(err)
		}

		var value string
		if err := rc.GetKey("yong", &value); err != nil {
			t.Error(err)
		}

		assert.Equal(t, "chan", value, "not expected result")
	})
}
