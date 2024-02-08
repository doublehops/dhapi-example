-- You need to separate multiple queries with this dotted line: ------------------

CREATE TABLE user (
    id INT(11) NOT NULL,
    email_address VARCHAR(100) NOT NULL,
    password VARCHAR(255),
    active TINYINT(4),
    email_address_verified_token VARCHAR(255),
    email_address_verified_at DATETIME,
    password_reset_token VARCHAR(255),
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME,
    PRIMARY KEY (id)
 );

    ------------------

CREATE TABLE user_token (
    id INT(11) NOT NULL,
    user_id INT(11) NOT NULL,
    expires DATETIME,
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES user(id)
);
	