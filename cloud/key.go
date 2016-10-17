package cloud

type CloudKey struct {
	ID          int    `json:"id,float64,omitempty"`
	Name        string `json:"name,omitempty"`
	Fingerprint string `json:"fingerprint,omitempty"`
	PublicKey   string `json:"public_key,omitempty"`
}

type KeyGetRequest struct {
	ID int
	TokenSource
}

type KeyCreateRequest struct {
	TokenSource
	Name      string `json:"name,omitempty"`
	PublicKey string `json:"public_key,omitempty"`
}

type ICloudKey interface {
	Get(*KeyGetRequest) (*CloudKey, error)
	List(*TokenSource) ([]*CloudKey, error)
	Create(*KeyCreateRequest) (*CloudKey, error)
}


