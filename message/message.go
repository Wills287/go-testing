package message

import "fmt"

/*
Greet is exported and can be used for blackbox testing.
*/
func Greet(name string) string {
	return fmt.Sprintf("Hello, %v!\n", name)
}

// Not exported, must use whitebox testing.
func depart(name string) string {
	return fmt.Sprintf("Goodbye, %v\n", name)
}
