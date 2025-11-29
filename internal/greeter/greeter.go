// Package greeter provides greeting functionality.
package greeter

import "fmt"

// Greeter holds the name of the person to greet.
type Greeter struct {
	name string
}

// New creates a new Greeter with the specified name.
func New(name string) *Greeter {
	return &Greeter{name: name}
}

// Greet returns a greeting message.
func (g *Greeter) Greet() string {
	if g.name == "" {
		return "Hello, World!"
	}
	return fmt.Sprintf("Hello, %s!", g.name)
}

// Name returns the name of the person being greeted.
func (g *Greeter) Name() string {
	return g.name
}
