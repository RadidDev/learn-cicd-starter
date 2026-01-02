package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		headers http.Header
		want    string
		wantErr bool
	}{
		"valid api key": {
			headers: http.Header{"Authorization": []string{"ApiKey secret-token-123"}},
			want:    "secret-token-123",
			wantErr: false,
		},
		"missing authorization header": {
			headers: http.Header{},
			want:    "",
			wantErr: true,
		},
		"malformed authorization header (missing ApiKey prefix)": {
			headers: http.Header{"Authorization": []string{"Bearer some-token"}},
			want:    "",
			wantErr: true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := GetAPIKey(tc.headers)
			if (err != nil) != tc.wantErr {
				t.Fatalf("GetAPIKey() error = %v, wantErr %v", err, tc.wantErr)
			}
			if got != tc.want {
				t.Errorf("GetAPIKey() got = %v, want %v", got, tc.want)
			}
		})
	}
}
