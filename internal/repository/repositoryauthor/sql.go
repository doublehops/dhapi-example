package repositoryauthor

var cc = `INSERT INTO author (
	user_id,
	name,
    created_by,
    updated_by,
    created_at,
    updated_at
	  ) VALUES (
	$1,
	$2,
	$3,
	$4
	)
`

var insertRecordSQL = `INSERT INTO author (
	user_id,
	name,
    created_by,
    updated_by,
    created_at,
    updated_at
	  ) VALUES (
	$1,
	$2,
	$3,
	$4,
	$5,
	$6
	)
`

var updateRecordSQL = `UPDATE author SET 
	name=$1,
    updated_by=$2,
    updated_at=$3
	WHERE id=$4
`

var deleteRecordSQL = `UPDATE author SET 
    updated_by=$1,
    updated_at=$2,
    deleted_at=$3
	WHERE id=$4
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
    WHERE id=$1
    AND deleted_at IS NULL`

var selectCollectionQuery = `SELECT 
    id,
    user_id,
    name,
    created_by,
    updated_by,
    created_at,
    updated_at
    FROM author
`

var selectCollectionCountQuery = `SELECT 
    COUNT(*) count
    FROM author
`
