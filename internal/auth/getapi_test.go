package auth

import (
	"net/http"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	type testCase struct {
		name        string
		headers     http.Header
		expected    string
		expectedErr bool
	}

	testCases := []testCase{
		{
			name:        "valid API key",
			headers:     http.Header{"Authorization": []string{"ApiKey ligma"}},
			expected:    "ligm",
			expectedErr: false,
		},
		{
			name:        "missing api key",
			headers:     http.Header{"Authorization": []string{""}},
			expected:    "",
			expectedErr: true,
		},
		{
			name:        "no authorization header",
			headers:     http.Header{},
			expected:    "",
			expectedErr: true,
		},
		{
			name:        "malformed authorization header",
			headers:     http.Header{"Authorization": []string{"malformedHeader malformedApiKey"}},
			expected:    "",
			expectedErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := GetAPIKey(tc.headers)
			if (err != nil) != tc.expectedErr {
				t.Fatalf("error expected mismatch for case %s: got err %v expected err %v", tc.name, err, tc.expectedErr)
			}

			if result != tc.expected {
				t.Fatalf("got %v, expected %v for case %s", result, tc.expected, tc.name)
			}
		})
	}
}
