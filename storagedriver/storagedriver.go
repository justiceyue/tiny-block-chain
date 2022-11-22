package storagedriver

type Setter interface {
	SetBlock(key []byte, value []byte) error
}

type Getter interface {
	GetBlock(value []byte) ([]byte, error)
}

type StorageDriver interface {
	Setter
	Getter
	Name() string
	Close()
}
