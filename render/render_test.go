package render

import (
	"math"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestJSON(t *testing.T) {
	data := struct {
		Name string
	}{
		Name: "Foo",
	}

	w := httptest.NewRecorder()
	err := JSON(w, data)

	got := strings.TrimSpace(w.Body.String())
	want := `{"Name":"Foo"}`

	if got != want {
		t.Errorf("Want: %v, Got: %v\n", want, got)
	}

	if err != nil {
		t.Errorf("Want: %v, Got: %v\n", nil, err)
	}
}

func TestJSONError(t *testing.T) {
	w := httptest.NewRecorder()
	err := JSON(w, math.NaN())

	if err == nil {
		t.Errorf("Want: not %v, Got: %v\n", err, nil)
	}
}
