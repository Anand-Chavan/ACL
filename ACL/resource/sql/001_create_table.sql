USE restapi;

CREATE TABLE IF NOT EXISTS user_detail (
    id                  INT         AUTO_INCREMENT      PRIMARY KEY,
    first_name          CHAR(25)    NOT NULL,
    last_name           CHAR(25)    NOT NULL,
    email               CHAR(64)    NOT NULL UNIQUE,
    password            VARBINARY(128)    NOT NULL,
    contact_number      CHAR(15)    NOT NULL,
    updated_by          INT         NOT NULL DEFAULT 0,
    deleted             TINYINT(1)  NOT NULL DEFAULT 0,
    creation_date       DATETIME    DEFAULT CURRENT_TIMESTAMP,
    last_update         DATETIME    DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
)ENGINE = INNODB CHARACTER SET=utf8;

DROP TABLE IF EXISTS books;
CREATE TABLE IF NOT EXISTS books (
    id   INT(11) AUTO_INCREMENT,
    title   VARCHAR(100)  NOT NULL,
    content       LONGTEXT  NOT NULL,
    PRIMARY KEY(id)
 );

DROP TABLE IF EXISTS users;
CREATE TABLE IF NOT EXISTS users (
    -- id   INT(11) AUTO_INCREMENT,
    uName VARCHAR(30) NOT NULL,
    userId  VARCHAR(30) NOT NULL,
    password  VARCHAR(20) NOT NULL,
    dateCreation TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ,
    userType VARCHAR(30) NOT NULL,
    PRIMARY KEY (userId)
 );
 INSERT INTO users (uName,userId,password,userType) VALUES ('sagar','sagar','".md5('123')."','s');
 INSERT INTO users (uName,userId,password,userType) VALUES ('chopade','chopade','".md5('123')."','s');
DROP TABLE IF EXISTS usersKey;
CREATE TABLE IF NOT EXISTS usersKey (
    userId  VARCHAR(30) NOT NULL,
    sessionKey VARCHAR(100) NOT NULL,
    logDate TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ,
    FOREIGN KEY (userId) REFERENCES users(userId),
    PRIMARY KEY (userId,sessionKey)
 );
    INSERT INTO usersKey (userId,sessionKey) VALUES ('sagar',replace(uuid(),'-',''));
