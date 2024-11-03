package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	header := http.Header{
		"Authorization": []string{"ApiKey thisisauthkey"},
	}
	want := "thisisauthkey"

	got, err := GetAPIKey(header)
	if err != nil {
		t.Fatalf("Error should not rise: %s", err)
	}

	if want != got {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}
