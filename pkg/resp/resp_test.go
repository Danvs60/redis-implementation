package resp

import "testing"

func TestEncodeSimpleString(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{
			name:  "valid",
			input: "OK",
			want:  "+OK\r\n",
		},
		{
			name:    "invalid",
			input:   "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EncodeSimpleString(tt.input)
			if (err != nil) != tt.wantErr {
				t.Error("wanted error, but didn't get one")
			}
			if got != tt.want {
				t.Errorf("want %q, got %q", tt.want, got)
			}
		})
	}
}

func TestDecodeSimpleString(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{
			name:    "valid",
			input:   "+OK\r\n",
			want:    "OK",
			wantErr: false,
		},
		{
			name:    "invalid",
			input:   "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DecodeSimpleString(tt.input)
			if (err != nil) != tt.wantErr {
				t.Error("wanted error, but didn't get one")
			}
			if got != tt.want {
				t.Errorf("want %q, got %q", tt.want, got)
			}
		})
	}

}

func TestEncodeErrorMessage(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{
			name:  "valid",
			input: "OK",
			want:  "-OK\r\n",
		},
		{
			name:    "invalid",
			input:   "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EncodeErrorMessage(tt.input)
			if (err != nil) != tt.wantErr {
				t.Error("wanted error, but didn't get one")
			}
			if got != tt.want {
				t.Errorf("want %q, got %q", tt.want, got)
			}
		})
	}
}

func TestDecodeErrorMessage(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{
			name:    "valid",
			input:   "-OK\r\n",
			want:    "OK",
			wantErr: false,
		},
		{
			name:    "invalid",
			input:   "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DecodeErrorMessage(tt.input)
			if (err != nil) != tt.wantErr {
				t.Error("wanted error, but didn't get one")
			}
			if got != tt.want {
				t.Errorf("want %q, got %q", tt.want, got)
			}
		})
	}

}
