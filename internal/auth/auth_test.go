package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	type test struct {
		input http.Header
		want  string
	}

	tests := []test{
		{
			input: http.Header{"Authorization": []string{"ApiKey thisisauthkey"}},
			want:  "thisisauthkey",
		},
		{
			input: http.Header{"Authorization": []string{"ApiKey token"}},
			want:  "token",
		},
	}

	for _, tc := range tests {
		got, err := GetAPIKey(tc.input)
		if err != nil {
			t.Fatalf("Error should not rise: %s", err)
		}
		if tc.want != got {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}

	}
}
