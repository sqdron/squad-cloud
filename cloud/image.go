package cloud

type Image struct {
	ID           int      `json:"id,float64,omitempty"`
	Name         string   `json:"name,omitempty"`
	Distribution string   `json:"distribution,omitempty"`
	Slug         string   `json:"slug,omitempty"`
	Regions      []string `json:"regions,omitempty"`
	MinDiskSize  int      `json:"min_disk_size,omitempty"`
	Created      string   `json:"created_at,omitempty"`
}

type IImage interface {
	List(*TokenSource) ([]*Image, error)
}


