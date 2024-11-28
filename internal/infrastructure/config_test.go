package infrastructure

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildApiUri(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		version int
		path    string
		want    string
	}{
		{1, "endpoint", "https://makers-challenge.altscore.ai/v1/endpoint"},
		{2, "test", "https://makers-challenge.altscore.ai/v2/test"},
		{3, "api/test", "https://makers-challenge.altscore.ai/v3/api/test"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("version=%d,path=%s", tt.version, tt.path), func(t *testing.T) {
			got := buildApiUri(tt.version, tt.path)
			assert.Equal(tt.want, got, "buildApiUri(%d, %s) = %v, want %v", tt.version, tt.path, got, tt.want)
		})
	}
}
