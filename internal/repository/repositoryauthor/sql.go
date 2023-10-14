package repositoryauthor

var insertRecordSQL = `INSERT INTO author (
	name,
    created_by,
    updated_by,
    created_at,
    updated_at
	  ) VALUES (
	?,
	?,
	?,
	?,
	?
	)
`

var updateRecordSQL = `UPDATE author SET 
	name=?,
    updated_by=?,
    updated_at=?
	WHERE id=?
`

var selectByIDQuery = `SELECT 
    id,
    name,
    created_by,
    updated_by,
    created_at,
    updated_at
    FROM author
    WHERE id=?`

var selectAllQuery = `SELECT 
    id,
    name,
    created_by,
    updated_by,
    created_at,
    updated_at
    FROM author
`
