package mynewtablerepository

var insertRecordSQL = `INSERT INTO my_new_table (
	currency_id,
	name,
	created_at,
	updated_at,
	deleted_at
	  ) VALUES (
?,
?,
?,
?,
?
	)
`

var updateRecordSQL = `UPDATE my_new_table SET
	currency_id=?,
	name=?,
	created_at=?,
	updated_at=?,
	deleted_at=?
	WHERE id=?
`

var deleteRecordSQL = `UPDATE my_new_table SET
    updated_at=?,
    deleted_at=?
	WHERE id=?
`

var selectByIDQuery = `SELECT 
	id,
	currency_id,
	name,
	created_at,
	updated_at,
	deleted_at
    FROM my_new_table
    WHERE id=?
    AND deleted_at IS NULL`

var selectCollectionQuery = `SELECT 
	id,
	currency_id,
	name,
	created_at,
	updated_at,
	deleted_at
    FROM my_new_table
`

var selectCollectionCountQuery = `SELECT 
    COUNT(*) count
    FROM my_new_table
`
