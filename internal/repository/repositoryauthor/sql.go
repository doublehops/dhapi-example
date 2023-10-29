package repositoryauthor

var insertRecordSQL = `INSERT INTO author (
	user_id,
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

var deleteRecordSQL = `UPDATE author SET 
    updated_by=?,
    updated_at=?,
    deleted_at=?
	WHERE id=?
`

var selectByIDQuery = `SELECT 
    id,
    user_id,
    name,
    created_by,
    updated_by,
    created_at,
    updated_at
    FROM author
    WHERE id=?
    AND deleted_at IS NULL`

var selectCollectionCountQuery = `SELECT 
    count(*) count
    FROM author
    WHERE deleted_at IS NULL
`

var selectCollectionQuery = `SELECT 
    __COUNT__
    id,
    user_id,
    name,
    created_by,
    updated_by,
    created_at,
    updated_at
    FROM author
	__WHERE_CLAUSE__
	__PAGINATION__
`
var selectCollectionQueryORIGINAL = `SELECT 
    id,
    user_id,
    name,
    created_by,
    updated_by,
    created_at,
    updated_at
    FROM author
    WHERE deleted_at IS NULL
	LIMIT ?, ?
`
