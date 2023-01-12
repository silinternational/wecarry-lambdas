package domain

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func Test_RunTask(t *testing.T) {
	tests := []struct {
		name string
		task string
	}{
		{
			name: "files",
			task: "file_cleanup",
		},
		{
			name: "tokens",
			task: "token_cleanup",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			}))
			defer server.Close()

			err := RunTask(server.URL, tt.task, 2*time.Second)
			if err != nil {
				t.Error(err)
			}
		})
	}
}
