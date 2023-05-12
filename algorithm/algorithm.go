package algorithm

type Encryptable interface {
	~string | byte
}

type Encryptor[E Encryptable] interface {
	Encrypt(slice E) E
	Decrypt(slice E) E
}
