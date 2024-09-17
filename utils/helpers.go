package utils

// IsAlphanumeric checks if a rune is an alphanumeric character
func IsAlphanumeric(char rune) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9')
}
