package daos

import (
	"database/sql"
	"errors"
	"github.com/chandrababu1609/bheem/test989/pkg/rest/server/daos/clients/sqls"
	"github.com/chandrababu1609/bheem/test989/pkg/rest/server/models"
	log "github.com/sirupsen/logrus"
)

type TestfiledsDao struct {
	sqlClient *sqls.SQLiteClient
}

func migrateTestfileds(r *sqls.SQLiteClient) error {
	query := `
	CREATE TABLE IF NOT EXISTS testfileds(
		Id INTEGER PRIMARY KEY AUTOINCREMENT,
        
		Password TEXT NOT NULL,
        CONSTRAINT id_unique_key UNIQUE (Id)
	)
	`
	_, err1 := r.DB.Exec(query)
	return err1
}

func NewTestfiledsDao() (*TestfiledsDao, error) {
	sqlClient, err := sqls.InitSqliteDB()
	if err != nil {
		return nil, err
	}
	err = migrateTestfileds(sqlClient)
	if err != nil {
		return nil, err
	}
	return &TestfiledsDao{
		sqlClient,
	}, nil
}

func (testfiledsDao *TestfiledsDao) CreateTestfileds(m *models.Testfileds) (*models.Testfileds, error) {
	insertQuery := "INSERT INTO testfileds(Password)values(?)"
	res, err := testfiledsDao.sqlClient.DB.Exec(insertQuery, m.Password)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	m.Id = id

	log.Debugf("testfileds created")
	return m, nil
}

func (testfiledsDao *TestfiledsDao) UpdateTestfileds(id int64, m *models.Testfileds) (*models.Testfileds, error) {
	if id == 0 {
		return nil, errors.New("invalid updated ID")
	}
	if id != m.Id {
		return nil, errors.New("id and payload don't match")
	}

	testfileds, err := testfiledsDao.GetTestfileds(id)
	if err != nil {
		return nil, err
	}
	if testfileds == nil {
		return nil, sql.ErrNoRows
	}

	updateQuery := "UPDATE testfileds SET Password = ? WHERE Id = ?"
	res, err := testfiledsDao.sqlClient.DB.Exec(updateQuery, m.Password, id)
	if err != nil {
		return nil, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, sqls.ErrUpdateFailed
	}

	log.Debugf("testfileds updated")
	return m, nil
}

func (testfiledsDao *TestfiledsDao) DeleteTestfileds(id int64) error {
	deleteQuery := "DELETE FROM testfileds WHERE Id = ?"
	res, err := testfiledsDao.sqlClient.DB.Exec(deleteQuery, id)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sqls.ErrDeleteFailed
	}

	log.Debugf("testfileds deleted")
	return nil
}

func (testfiledsDao *TestfiledsDao) ListTestfileds() ([]*models.Testfileds, error) {
	selectQuery := "SELECT * FROM testfileds"
	rows, err := testfiledsDao.sqlClient.DB.Query(selectQuery)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)
	var testfileds []*models.Testfileds
	for rows.Next() {
		m := models.Testfileds{}
		if err = rows.Scan(&m.Id, &m.Password); err != nil {
			return nil, err
		}
		testfileds = append(testfileds, &m)
	}
	if testfileds == nil {
		testfileds = []*models.Testfileds{}
	}

	log.Debugf("testfileds listed")
	return testfileds, nil
}

func (testfiledsDao *TestfiledsDao) GetTestfileds(id int64) (*models.Testfileds, error) {
	selectQuery := "SELECT * FROM testfileds WHERE Id = ?"
	row := testfiledsDao.sqlClient.DB.QueryRow(selectQuery, id)
	m := models.Testfileds{}
	if err := row.Scan(&m.Id, &m.Password); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sqls.ErrNotExists
		}
		return nil, err
	}

	log.Debugf("testfileds retrieved")
	return &m, nil
}
