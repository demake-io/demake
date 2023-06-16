module github.com/demake-io/demake

go 1.20

require (
	github.com/gospel-dev/gospel v0.0.0-20230404113406-1de6b80cc11c
	github.com/jackc/pgx/v5 v5.3.1
)

replace github.com/gospel-dev/gospel => ./modules/gospel
