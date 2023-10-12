package repositoryauthor

var sql = `INSERT INTO author
	name,
    created_at,
    updated_at
		VALUES
	$1,
	$2,
	$3
`
