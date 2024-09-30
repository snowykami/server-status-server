package dao

type Report struct {
	// 鉴权字段
	Auth struct {
		Token string `json:"token"` // 令牌，用于鉴权，防止恶意请求
	} `json:"auth"`

	Meta struct {
		Name string `json:"name"`
		OS   struct {
			Name    string `json:"name"`
			Version string `json:"version"`
		}
		Labels   []string `json:"labels"`   // 服务器标签
		Location string   `json:"location"` // Chongqing, China
		Duration int64    `json:"duration"` // uptime in seconds
	}

	Hardware struct {
		Mem struct {
			Total int64 `json:"total"`
			Used  int64 `json:"used"`
		} `json:"mem"`
		Swap struct {
			Total int64 `json:"total"`
			Used  int64 `json:"used"`
		} `json:"swap"`
		Cpu struct {
			Cores  int `json:"cores"`
			Logics int `json:"logics"`
		} `json:"cpu"`
		Disks []struct {
			Name  string `json:"name"`
			Total int64  `json:"total"`
		} `json:"disks"`
		Net []struct {
			Up   int64  `json:"up"`
			Down int64  `json:"down"`
			Type string `json:"type"` // IPv4 or IPv6 or IPv4/6
		} `json:"net"`
	} `json:"hardware"`
}
