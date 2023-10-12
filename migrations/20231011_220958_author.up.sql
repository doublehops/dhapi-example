CREATE TABLE author
(
    id         INT(11) NOT NULL,
    name       varchar(100) NOT NULL,
    created_by INT(11),
    updated_by INT(11),
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME,
    PRIMARY KEY (id)
);
	