package checker

import (
	"io"
	"strings"
	"testing"
)

func TestJSON_Check(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		j       JSON
		args    args
		want    bool
		wantErr bool
	}{
		{"validjson", JSON{}, args{r: strings.NewReader("{}")}, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := JSON{}
			got, err := j.Check(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("JSON.Check() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("JSON.Check() = %v, want %v", got, tt.want)
			}
		})
	}
}
