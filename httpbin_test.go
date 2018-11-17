package httpbin

import "testing"

func TestHttpbinNew(t *testing.T) {
	got := "foo"
	want := "foo"

	if want != got {
		t.Errorf("Want %v, Got %v\n", want, got)
	}
}
