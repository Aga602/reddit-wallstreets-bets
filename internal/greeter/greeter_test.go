package greeter_test

import (
	"testing"

	"github.com/Aga602/reddit-wallstreets-bets/internal/greeter"
)

func TestGreeter_Greet(t *testing.T) {
	tests := []struct {
		name     string
		greetee  string
		expected string
	}{
		{
			name:     "greet with name",
			greetee:  "Alice",
			expected: "Hello, Alice!",
		},
		{
			name:     "greet with empty name",
			greetee:  "",
			expected: "Hello, World!",
		},
		{
			name:     "greet with different name",
			greetee:  "Bob",
			expected: "Hello, Bob!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := greeter.New(tt.greetee)
			result := g.Greet()
			if result != tt.expected {
				t.Errorf("Greet() = %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestGreeter_Name(t *testing.T) {
	tests := []struct {
		name     string
		greetee  string
		expected string
	}{
		{
			name:     "returns name",
			greetee:  "Alice",
			expected: "Alice",
		},
		{
			name:     "returns empty string",
			greetee:  "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := greeter.New(tt.greetee)
			result := g.Name()
			if result != tt.expected {
				t.Errorf("Name() = %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestNew(t *testing.T) {
	g := greeter.New("Test")
	if g == nil {
		t.Error("New() returned nil")
	}
}
