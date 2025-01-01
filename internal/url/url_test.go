package url

import "testing"

func TestShorten(t *testing.T) {
	type args struct {
		originalURL string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Shorten(tt.args.originalURL); got != tt.want {
				t.Errorf("Shorten() = %v, want %v", got, tt.want)
			}
		})
	}
}
