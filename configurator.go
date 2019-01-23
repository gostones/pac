package pac

// Configurator is a interface of setup ergo configuration
type Configurator interface {
	SetUp(proxyURL string) error
	SetDown() error
	Get() string
}

// GetConfigurator gets the right configurator strategy for a given system
func GetConfigurator(system string) Configurator {
	switch system {
	case "windows":
		return &WindowsConfigurator{}
	case "darwin":
		return &OSXConfigurator{}
	case "linux":
		return &LinuxConfigurator{}
	default:
		return nil
	}
}
