// Package twofer can share things with people
package twofer

import "fmt"

// ShareWith returns a string of the form "One for {input}, one for me.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
