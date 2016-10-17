package cloud

type CloudUnit struct {
	ID      int
	Name    string
	Memory  int
	Vcpus   int
	Disk    int
	Region  *UnitRegion
	Created string
	Tags    []string
}

type UnitRegion struct {
	Name string
}

type TokenSource struct {
	Token string
}

type UnitCreateRequest struct {
	Token  string        `json:"token"`
	Name   string        `json:"name"`
	Region string        `json:"region"`
	Size   string        `json:"size"`
	Image  string        `json:"image"`
	Key    int           `json:"key,string,omitempty"`
}

type ICloudUnit interface {
	Create(r *UnitCreateRequest) (*CloudUnit, error)
	List(*TokenSource) ([]*CloudUnit, error)
	Delete(*CloudUnit) (error)
}

