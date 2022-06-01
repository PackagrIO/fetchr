package config

func Create() (Interface, error) {
	config := new(BaseConfiguration)
	if err := config.Init(); err != nil {
		return nil, err
	}
	return config, nil
}
