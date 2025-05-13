package auth

import (
	"errors"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetAPIKey(t *testing.T) {
	testCases := map[string]struct {
		input    http.Header
		expected string
		err      error
	}{
		"No Authorization header": {
			input:    http.Header{},
			expected: "",
			err:      ErrNoAuthHeaderIncluded,
		},
		"Empty Authorization header": {
			input:    http.Header{"Authorization": []string{""}},
			expected: "",
			err:      ErrNoAuthHeaderIncluded,
		},
		"Malformed header": {
			input:    http.Header{"Authorization": []string{"ApiK", "123"}},
			expected: "",
			err:      errors.New("malformed authorization header"),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got, gotErr := GetAPIKey(tc.input)
			if tc.err != nil && gotErr == nil {
				t.Fatal("Expected error error", tc.err)
			}
			if tc.err != nil && gotErr != nil {
				t.Fatal("Unexpected error", gotErr)
			}
			diff := cmp.Diff(got, tc.expected)
			if diff != "" {
				t.Fatal(diff)
			}
		})
	}
}
