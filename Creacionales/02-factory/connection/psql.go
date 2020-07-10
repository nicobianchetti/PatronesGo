package connection

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type Postgres struct {
	db *sql.DB
}

func (p *Postgres) Connect() error {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		"postgres",
		"postgres",
		"127.0.0.1",
		"5432",
		"postgres",
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}

	p.db = db
	return nil
}

func (p *Postgres) GetNow() (*time.Time, error) {
	t := &time.Time{}
	err := p.db.QueryRow("select CURRENT_DATE").Scan(t) //curdate funcion en mysql para obtener la fecha actual. Scan es para mandar el resultado de la query a t
	if err != nil {
		fmt.Printf("error al leer la fecha del servidor: %v", err)
		return nil, err
	}

	return t, nil
}

func (p *Postgres) Close() error {
	return p.db.Close()
}
