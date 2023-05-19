CREATE TABLE users(
    id INT NOT NULL AUTO_INCREMENT,
    user_uuid VARCHAR(250) NOT NULL,
    name VARCHAR(250) NOT NULL,
    email VARCHAR(250) NOT NULL,
    password VARCHAR(250) NOT NULL,
    PRIMARY KEY (id)
)ENGINE = InnoDB;
create index user_uuid on users(user_uuid);