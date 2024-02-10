-- You need to separate multiple queries with this dotted line: ------------------

CREATE TABLE my_new_table (
    id INT(11) NOT NULL,
    currency_id INT(11) NOT NULL,
    name VARCHAR(100),
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME,
    PRIMARY KEY (id)
    );
	