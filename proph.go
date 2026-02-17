package primitive

type ProphId struct{}

func NewProph() ProphId {
	return ProphId{}
}

func (p ProphId) ResolveBool(b bool)  {}
func (p ProphId) ResolveU64(i uint64) {}
