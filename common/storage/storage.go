package storage

type IStorable interface {
	GetID() int64
	SetID(int64)
}

type IStorage[T IStorable] interface {
	Add(T)
	Delete(T)
	DeleteById(int64)
	Exists(T) bool
	ExistsKey(int64) bool
	Update(int64, T) bool
	GetAll() []T
}
