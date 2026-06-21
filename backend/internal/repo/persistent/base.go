package persistent

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
)

type BaseRepo struct {
	DB      *sql.DB
	Builder squirrel.StatementBuilderType
}

func NewBaseRepo(db *sql.DB) BaseRepo {
	return BaseRepo{
		DB:      db,
		Builder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).RunWith(db),
	}
}
