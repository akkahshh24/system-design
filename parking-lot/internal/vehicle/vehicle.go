package vehicle

// IVechile interface defines methods that any vehicle should implement
type IVechile interface {
	GetType() Type
	GetLicensePlate() string
}
