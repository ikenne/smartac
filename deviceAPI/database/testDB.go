package database

import (
	"deviceAPI/lib"
	"sort"
	"sync"
	"time"
)

type TestDB struct {
	devices []lib.Device
}

var (
	mutex sync.Mutex
)

func NewTestDB() *TestDB {
	db := TestDB{}
	db.devices = populateDevices()
	return &db
}

func (t *TestDB) GetAllDevices() []lib.Device {
	return t.devices
}

func (t *TestDB) GetUnHealthyDevices(status map[string]struct{}) []lib.Device {
	var ud []lib.Device

	for _, v := range t.devices {
		if _, ok := status[v.Status]; ok {
			ud = append(ud, v)
		}
	}

	return ud
}

func (t *TestDB) GetCO2LimitDevice(limit float64) []lib.Device {
	var ud []lib.Device

	for _, v := range t.devices {
		if len(v.CO2) > 0 {
			var co2 = make([]lib.Sensor, len(v.CO2))
			copy(co2, v.CO2)

			sort.Slice(co2, func(i, j int) bool {
				return co2[i].Date.After(co2[j].Date)
			})

			first := co2[0]
			if first.Value > limit {
				ud = append(ud, v)
			}
		}
	}

	return ud
}

func (t *TestDB) GetDevice(sn string) (lib.Device, error) {
	for _, a := range t.devices {
		if a.SerialNumber == sn {
			return a, nil
		}
	}

	return lib.Device{}, lib.ErrorKeyNotFound{Key: sn, Message: "Device Not found"}
}

func (t *TestDB) AddDevice(d lib.Device) {
	mutex.Lock()
	defer mutex.Unlock()

	t.devices = append(t.devices, d)
}

func (t *TestDB) UpdateDevice(sn string, d lib.Device) (lib.Device, error) {
	mutex.Lock()
	defer mutex.Unlock()

	for i, v := range t.devices {
		if v.SerialNumber == sn {
			t.devices[i].Temperature = append(t.devices[i].Temperature, d.Temperature...)
			t.devices[i].Humidity = append(t.devices[i].Humidity, d.Humidity...)
			t.devices[i].CO2 = append(t.devices[i].CO2, d.CO2...)

			if d.Status != "" {
				t.devices[i].Status = d.Status
			}

			return t.devices[i], nil
		}
	}

	return lib.Device{}, lib.ErrorKeyNotFound{Key: sn, Message: "Device Not found"}
}

func (t *TestDB) DeleteDevice(sn string) error {
	mutex.Lock()
	defer mutex.Unlock()

	for i, a := range t.devices {
		if a.SerialNumber == sn {
			t.devices = append(t.devices[:i], t.devices[i+1:]...)
			return nil
		}
	}

	return lib.ErrorKeyNotFound{Key: sn, Message: "Device Not found"}
}

func populateDevices() []lib.Device {
	// const shortForm = "2006-Jan-02"
	t1, _ := time.Parse(time.RFC3339, "2019-05-10T00:00:00Z")
	t2, _ := time.Parse(time.RFC3339, "2019-05-10T00:00:00Z")

	t3, _ := time.Parse(time.RFC3339, "2019-06-10T15:04:05Z")
	t4, _ := time.Parse(time.RFC3339, "2019-06-10T15:05:05Z")

	devices := []lib.Device{
		lib.Device{SerialNumber: "1", RegistrationDate: t1, FirmwareVersion: "1.2", Status: "healthy",
			Temperature: []lib.Sensor{
				lib.Sensor{Value: 32, Date: t3},
				lib.Sensor{Value: 31, Date: t4},
			},
		},
		lib.Device{SerialNumber: "2", RegistrationDate: t2, FirmwareVersion: "1.3", Status: "good",
			Temperature: []lib.Sensor{
				lib.Sensor{Value: 32, Date: t3},
				lib.Sensor{Value: 31, Date: t4},
			},
			CO2: []lib.Sensor{
				lib.Sensor{Value: 8, Date: t3},
				lib.Sensor{Value: 9.5, Date: t4},
			},
		},
		lib.Device{SerialNumber: "3", RegistrationDate: t1, FirmwareVersion: "1.2", Status: "gas_leak",
			Temperature: []lib.Sensor{
				lib.Sensor{Value: 10, Date: t3},
				lib.Sensor{Value: 12, Date: t4},
			},			
		},
	}

	return devices
}
