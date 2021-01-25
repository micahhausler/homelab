package spark

import (
	"github.com/kinvolk/ignition/config/v2_2/types"
)

type ConfigOpt func(*types.Config) error

func NewConfig(opts ...ConfigOpt) (*types.Config, error) {

	config := &types.Config{
		Ignition: types.Ignition{
			Version: "2.2.0",
		},
	}

	for _, opt := range opts {
		err := opt(config)
		if err != nil {
			return nil, err
		}
	}

	return config, nil
}
