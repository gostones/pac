package pac

import (
	"bufio"
	"strings"
)

// OSXConfigurator implements Configurator for windows
type OSXConfigurator struct{}

// SetUp is responsible for setting up the new proxy
func (c *OSXConfigurator) SetUp(proxyURL string) error {
	_, err := RunnerDefault.Run(`networksetup`, `-setautoproxyurl "Wi-Fi" "`+proxyURL+`"`)
	return err
}

// SetDown is responsible for remove the proxy
func (c *OSXConfigurator) SetDown() error {
	_, err := RunnerDefault.Run(`networksetup`, `-setautoproxyurl "Wi-Fi" ""`)
	return err
}

// Get returns set proxy URL if any
func (c *OSXConfigurator) Get() string {
	out, err := RunnerDefault.Run(`networksetup`, `-getautoproxyurl "Wi-Fi"`)
	if err != nil {
		return ""
	}
	scanner := bufio.NewScanner(strings.NewReader(out))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		s := scanner.Text()
		if strings.HasPrefix(s, "URL:") {
			sa := strings.SplitN(s, ":", 2)
			return strings.TrimSpace(sa[1])
		}
	}
	return ""
}
