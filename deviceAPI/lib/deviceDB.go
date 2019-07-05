package lib

type DeviceDB interface {
	GetAllDevices() []Device
	AddDevice(d Device)
	GetDevice(sn string) (Device, error)
	UpdateDevice(string, Device) (Device, error)
	DeleteDevice(sn string) error
	GetUnHealthyDevices(map[string]struct{}) []Device
	GetCO2LimitDevice(limit float64) []Device
}
