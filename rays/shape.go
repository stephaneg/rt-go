package rays

type Shape interface {
	ID() int64
	Name() string
}


type VoidShape struct {
	Id    int64
	Label string
}


func NewVoidShape() *VoidShape {
	return &VoidShape{
		Id:    0,
		Label: "void",
	}
}

func (shape VoidShape) ID() int64 {
	return shape.Id
}

func (shape VoidShape) Name() string {
	return shape.Label
}
