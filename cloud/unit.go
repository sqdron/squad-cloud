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
	Token string
	Name   string
	Region string
	Size   string
	Image  string
}

type ICloudUnit interface {
	Create(r *UnitCreateRequest) (*CloudUnit, error)
	List(*TokenSource) ([]*CloudUnit, error)
	Delete(*CloudUnit) (error)
}

