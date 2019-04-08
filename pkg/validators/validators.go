package validators

// Validator exposes an interface
// to the validation algorithns
type Validator interface {
	IsValid(input string) bool
}
