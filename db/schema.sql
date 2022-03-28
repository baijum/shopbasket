CREATE TABLE inventory(
id BIGSERIAL PRIMARY KEY ,
name VARCHAR(50),
description VARCHAR(200),
price INTEGER,
status BOOLEAN -- true indicates available
)