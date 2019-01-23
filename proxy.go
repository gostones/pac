package pac

import (
	"fmt"
	"runtime"
)

// SetupProxy sets up system proxy
func SetupProxy(proxyURL string, remove bool) error {
	fmt.Println("Current detected system: " + runtime.GOOS)

	// proxyURL := config.GetProxyPacURL()
	system := runtime.GOOS
	configurator := GetConfigurator(system)

	if configurator == nil {
		return fmt.Errorf("auto config not supported for OS: %v", system)
	}

	var err error
	if remove {
		err = configurator.SetDown()
	} else {
		err = configurator.SetUp(proxyURL)
	}

	if err != nil {
		return fmt.Errorf("Setup failed cause %s", err)
	}

	return nil
}
