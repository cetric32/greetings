package greetings

import (
	"reflect"
	"testing"
)

func TestNewGreeting(t *testing.T) {
	type args struct {
		name    string
		message string
	}
	tests := []struct {
		name string
		args args
		want *Greeting
	}{
		// TODO: Add test cases.

		{
			name: "TestNewGreeting",
			args: args{
				name:    "TestNewGreeting",
				message: "TestNewGreeting",
			},
			want: &Greeting{
				Name:    "TestNewGreeting",
				Message: "TestNewGreeting",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGreeting(tt.args.name, tt.args.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGreeting() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGreeting_Greet(t *testing.T) {
	type fields struct {
		Name    string
		Message string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.

		{
			name: "TestGreeting_Greet",
			fields: fields{
				Name:    "TestGreeting_Greet",
				Message: "TestGreeting_Greet",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Greeting{
				Name:    tt.fields.Name,
				Message: tt.fields.Message,
			}
			g.Greet()
		})
	}
}

func TestGreet(t *testing.T) {
	type args struct {
		name    string
		message string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.

		{
			name: "TestGreet",
			args: args{
				name:    "TestGreet",
				message: "TestGreet",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Greet(tt.args.name, tt.args.message)
		})
	}
}
