package builder

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestResourcePoolConfigBuilderBuild(t *testing.T) {
	tests := []struct {
		name    string
		builder *ResourcePoolConfigBuilder
		want    *ResourcePoolConfig
		wantErr bool
	}{
		{
			name: "name empry",
			builder: &ResourcePoolConfigBuilder{
				name:     "",
				maxTotal: 0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "maxIdle < minIdle",
			builder: &ResourcePoolConfigBuilder{
				name:     "test",
				maxTotal: 0,
				maxIdle:  10,
				minIdle:  20,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "sucess",
			builder: &ResourcePoolConfigBuilder{
				name: "test",
			},
			want: &ResourcePoolConfig{
				name:     "test",
				maxTotal: defaultMaxTotal,
				maxIdle:  defaultMaxIdle,
				minIdle:  defaultMinIdle,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.builder.Build()
			require.Equalf(t, tt.wantErr, err != nil, "Build() error = %v, wantErr %v", err, tt.wantErr)
			assert.Equal(t, tt.want, got)
		})
	}
}
