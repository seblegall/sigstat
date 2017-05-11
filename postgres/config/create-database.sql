CREATE TABLE command_group (
id  SERIAL PRIMARY KEY,
name VARCHAR(100) NOT NULL
);

CREATE TABLE command (
group_id int4 REFERENCES command_group(id),
cmd VARCHAR(300) NOT NULL,
exit_code INT,
cmd_path VARCHAR(300) NOT NULL,
status INT
);
