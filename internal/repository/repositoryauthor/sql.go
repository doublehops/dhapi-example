package repositoryauthor

var InsertRecordSQL = `INSERT INTO author (
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
