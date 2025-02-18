package service

type ConfigurationInterface interface {
}

type configurationImpl struct{}

func Configuration() ConfigurationInterface {
	return &configurationImpl{}
}
