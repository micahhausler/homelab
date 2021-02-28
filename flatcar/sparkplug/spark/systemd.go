package spark

import (
	"github.com/kinvolk/ignition/config/v2_2/types"
)

func DisableUnitsConfigOpt(unitNames ...string) ConfigOpt {
	return ConfigOpt(func(config *types.Config) error {
		// Find unit if it exists
		for _, unitName := range unitNames {
			foundUnit := false
			for i, unit := range config.Systemd.Units {
				if unit.Name == unitName {
					config.Systemd.Units[i].Enabled = boolPtr(false)
					foundUnit = true
					break
				}
			}
			if foundUnit {
				continue
			}
			// Add it and disable it
			units := config.Systemd.Units
			config.Systemd.Units = append(units, types.Unit{
				Name:    unitName,
				Enabled: boolPtr(false),
			})
		}
		return nil
	})
}
