CREATE TABLE command_group (
id  SERIAL PRIMARY KEY UNIQUE,
name VARCHAR(100) NOT NULL UNIQUE
);

CREATE TABLE command (
id SERIAL PRIMARY KEY UNIQUE,
group_id int4 REFERENCES command_group(id),
cmd VARCHAR(300) NOT NULL,
exit_code INT,
cmd_path VARCHAR(300) NOT NULL,
status INT
);
