package algorithm

const (
	KEY_LENGTH_56 uint8 = 56
)

type DES[T Encryptable] struct {
	keyLength uint8
}

func NewDES[T Encryptable](keyLength uint8) DES[T] {
	return DES[T]{
		keyLength: KEY_LENGTH_56,
	}
}

func (a DES[T]) Encrypt(slice T) T {
	return slice
}

func (a DES[T]) Decrypt(slice T) T {
	return slice
}
