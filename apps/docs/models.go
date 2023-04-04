package writer

import (
	"github.com/demake-io/demake/sql"
)

type Action struct {
	sql.DBModel `db:"table:action"`
	sql.JSONModel
	sql.ConfigModel
	Name        string
	Type        string
	Description string
	//	ProjectID   int64
	//	Project     *Project `db:"fk:ProjectID"`
}

func Actions(db func() sql.DB, filters map[string]interface{}) ([]*Action, error) {
	action := &Action{}
	sql.Init(action, db)
	if actions, err := sql.LoadMany(action, filters); err != nil {
		return nil, err
	} else {
		sqlActions := make([]*Action, len(actions))
		for i, action := range actions {
			sqlActions[i] = action.(*Action)
		}
		return sqlActions, nil
	}
}

func (c *Action) Save() error {
	return sql.Save(c)
}

func (c *Action) ByExtID(id []byte) error {
	return sql.LoadOne(c, map[string]interface{}{"ext_id": id})
}

func (c *Action) ByID(id int64) error {
	return sql.LoadOne(c, map[string]interface{}{"id": id})
}
