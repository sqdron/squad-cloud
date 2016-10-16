package cloud

type CloudPlan struct{
	Name         string   `json:"slug,omitempty"`
	Memory       int      `json:"memory,omitempty"`
	Vcpus        int      `json:"vcpus,omitempty"`
	Disk         int      `json:"disk,omitempty"`
	PriceMonthly float64  `json:"price_monthly,omitempty"`
	PriceHourly  float64  `json:"price_hourly,omitempty"`
	Regions      []string `json:"regions,omitempty"`
}

type ICloudPlan interface {
	List(*TokenSource) ([]*CloudPlan, error)
}