package vehicle

// Define a new type VehicleType
type Type int

// Enum values using iota
// This ensures type safety and restricts values to predefined constants
const (
	Small Type = iota
	Medium
	Large
)

// String method to convert enum to readable string
func (v Type) String() string {
	switch v {
	case Small:
		return "Small"
	case Medium:
		return "Medium"
	case Large:
		return "Large"
	default:
		return "Unknown"
	}
}
