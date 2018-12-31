package storage

// import

// TODO: where do I define types/interfaces?

// Storage interfaces facilitate writing to local or remote data stores
type Storage interface {
	get()
	set()
	update()
	delete()
}

type LocalStorage struct {
}
