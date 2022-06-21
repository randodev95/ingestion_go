package utils

type UserDeviceInfo struct {
	IP    string `json:"ip"`
	Agent struct {
		Browser struct {
			Name    string
			Version string
		}
		Device struct {
			Name string
		}
		Os struct {
			Name    string
			Version string
		}
	}
	IsMobile  bool
	IsTablet  bool
	IsDesktop bool
	IsBot     bool
	IsIOS     bool
}
