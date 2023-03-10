package database

import (
	"sync"
	"testing"
)

func TestDatabase_Contains(t *testing.T) {
	type fields struct {
		mu   sync.Mutex
		Data map[string]string
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "Contains",
			fields: fields{Data: map[string]string{"a": "AAA"}},
			args:   args{"a"},
			want:   true,
		},
		{
			name:   "Contains not",
			fields: fields{Data: map[string]string{"a": "AAA"}},
			args:   args{"b"},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &Database{
				mu:   tt.fields.mu,
				Data: tt.fields.Data,
			}
			if got := db.Contains(tt.args.key); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
