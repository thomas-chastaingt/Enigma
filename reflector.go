package enigma

type Reflector struct {
	ID       string
	Sequence [26]int
}

func NewReflector(mapping string, id string) *Reflector {
	var seq [26]int
	for i, value := range mapping {
		seq[i] = CharToIndex(byte(value))
	}
	return &Reflector{id, seq}
}

type Reflectors []Reflector

func (refs *Reflectors) GetByID(id string) *Reflector {
	for _, ref := range *refs {
		if ref.ID == id {
			return &ref
		}
	}
	return nil
}
