package assert

import (
	"net/http"
	"testing"
)

func StatusOK(t *testing.T, statusCode int) {
	if statusCode == http.StatusOK {
		t.Errorf("Want %v, Got %v", http.StatusOK, statusCode)
	}
}
