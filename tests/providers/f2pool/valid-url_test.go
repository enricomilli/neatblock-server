package f2pool

import (
	"strings"
	"testing"

	"github.com/enricomilli/neat-server/api/v1/pools/providers/f2pool"
)

func TestF2Pool_ValidateURL(t *testing.T) {
	tests := []struct {
		name    string
		url     string
		wantErr bool
		errMsg  string
	}{
		{
			name:    "Valid F2Pool URL",
			url:     "https://www.f2pool.com/mining-user/6f2e4214d79688ab7697aca78f243d50?user_name=adamantiumnewvz",
			wantErr: false,
			errMsg:  "",
		},
		{
			name:    "Empty URL",
			url:     "",
			wantErr: true,
			errMsg:  "URL cannot be empty",
		},
		{
			name:    "Wrong domain",
			url:     "https://wrong-domain.com/mining-user/123?user_name=test",
			wantErr: true,
			errMsg:  "invalid domain",
		},
		{
			name:    "Wrong path",
			url:     "https://www.f2pool.com/wrong-path/123?user_name=test",
			wantErr: true,
			errMsg:  "invalid path",
		},
		// {
		// 	name:    "Missing user_name parameter",
		// 	url:     "https://www.f2pool.com/mining-user/123",
		// 	wantErr: true,
		// 	errMsg:  "missing required parameter: user_name",
		// },
		// {
		// 	name:    "Empty user_name parameter",
		// 	url:     "https://www.f2pool.com/mining-user/123?user_name=",
		// 	wantErr: true,
		// 	errMsg:  "user_name parameter cannot be empty",
		// },
		{
			name:    "Valid URL with different user_name",
			url:     "https://www.f2pool.com/mining-user/6f2e4214d79688ab7697aca78f243d50?user_name=adamantiumnuno",
			wantErr: false,
			errMsg:  "",
		},
	}

	provider := &f2pool.F2Pool{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := provider.ValidateURL(tt.url)

			// Check if error occurred when it shouldn't, or didn't occur when it should
			if (err != nil) != tt.wantErr {
				t.Errorf("F2Pool.ValidateURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// If we expect an error, check if the error message contains what we expect
			if tt.wantErr && err != nil && !strings.Contains(err.Error(), tt.errMsg) {
				t.Errorf("F2Pool.ValidateURL() error message = %v, want to contain %v", err, tt.errMsg)
			}
		})
	}
}
