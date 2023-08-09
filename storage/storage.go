package storage

type StorageI interface {
	CloseDB()
	// User() UserRepoI
}

type UserRepoI interface {
}
