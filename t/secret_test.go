package secret_test

import (
	"testing"

	"github.com/days365/gcpsm/secret"
)

func TestNewSecret(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		c := secret.Config{
			Version:   "1",
			ProjectID: "test",
			IsFile:    true,
			File:      "test.json",
			Name:      "test",
		}

		type Sec struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		}

		v := Sec{}
		if err := secret.NewSecret(c, &v); err != nil {
			t.Error(err)
		}

		if g, w := v.Key, "some key"; g != w {
			t.Errorf("want %s, but got %s", w, g)
		}
		if g, w := v.Value, "some value"; g != w {
			t.Errorf("want %s, but got %s", w, g)
		}
	})
}
