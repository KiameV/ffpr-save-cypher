package padder

type (
	Padder interface {
		Encode(source []byte) (result []byte)
		Decode(source []byte) (result []byte)
	}
	padder struct {
		blockSize int
	}
)

func New() Padder {

	return &padder{blockSize: 32}
}

func (p padder) Encode(source []byte) []byte {
	count := p.blockSize - ((len(source)+p.blockSize-1)%p.blockSize + 1)
	for i := 0; i < count; i++ {
		source = append(source, 0)
	}
	return source
}

func (p padder) Decode(source []byte) []byte {
	offset := len(source)
	if offset == 0 {
		return []byte{}
	}
	end := offset - p.blockSize + 1
	for offset > end {
		offset--
		if source[offset] != 0 {
			return source[:offset+1]
		}
	}
	return source[:end]
}
