//go:build android
// +build android

package platform

type Platform struct{}

const name = "platform"

func (self *Platform) Name() string {
	return name
}

func (self *Platform) Collect() (result any, err error) {
	result, err = getPlatformInfo()
	return
}

func getPlatformInfo() (platformInfo map[string]any, err error) {

	return
}
