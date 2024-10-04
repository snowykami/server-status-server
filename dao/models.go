package dao

type Report struct {
	// 鉴权字段
	Meta struct {
		ID   string `json:"id"` // 服务器ID，用于标识服务器
		Name string `json:"name"`
		OS   struct {
			Name    string `json:"name"`
			Version string `json:"version"`
		} `json:"os"`
		Labels     []string `json:"labels"`      // 服务器标签
		Location   string   `json:"location"`    // Chongqing, China
		UpTime     int64    `json:"uptime"`      // uptime in seconds
		Link       string   `json:"link"`        // 链接或是nil
		ObservedAt int64    `json:"observed_at"` // unix timestamp
		StartTime  int64    `json:"start_time"`  // unix timestamp
	} `json:"meta"`

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
			Cores   int     `json:"cores"`
			Logics  int     `json:"logics"`
			Percent float32 `json:"percent"` // 0-100
		} `json:"cpu"`
		Disks map[string]map[string]int64 `json:"disks"`
		Net   struct {
			Up   int64  `json:"up"`
			Down int64  `json:"down"`
			Type string `json:"type"` // IPv4 or IPv6 or IPv4/6
		} `json:"net"`
	} `json:"hardware"`
}
