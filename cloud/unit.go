package cloud

type Unit struct {

}

type IUnit interface {
	Create() (*Unit, error)
	List() ([]*Unit, error)
	Delete(*Unit) (error)
}