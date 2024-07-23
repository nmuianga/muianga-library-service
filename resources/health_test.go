package resources

import (
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"testing"
)

func TestHealth(t *testing.T) {
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "First test",
			args: args{c: ctx},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Health(tt.args.c)
		})
	}
}
