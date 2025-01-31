package sqlstore

import (
	"database/sql"
	"github.com/Kravatka/somego/internal/app/store"

	_ "github.com/lib/pq" //анонимный импорт что бы методы сами не импортировались в код
)

type Store struct {
	db 				*sql.DB
	userRepository 	*UserRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}