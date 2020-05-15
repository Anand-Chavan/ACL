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
INSERT INTO users (uName,userId,password,userType) VALUES ('sagar','sagar','123','s');
--  INSERT INTO users (uName,userId,password,userType) VALUES ('chopade','chopade','".md5('123')."','s');
DROP TABLE IF EXISTS usersKey;
CREATE TABLE IF NOT EXISTS usersKey (
    userId  VARCHAR(30) NOT NULL,
    sessionKey VARCHAR(100) NOT NULL,
    logDate TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ,
    FOREIGN KEY (userId) REFERENCES users(userId),
    PRIMARY KEY (userId,sessionKey)
 );
    INSERT INTO usersKey (userId,sessionKey) VALUES ('sagar',replace(uuid(),'-',''));

DROP TABLE IF EXISTS groups;
CREATE TABLE IF NOT EXISTS groups(
    -- groupId  INT NOT NULL AUTO_INCREMENT,
    groupName  VARCHAR(30) NOT NULL,
    groupCreationDate  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ,
    groupDescription VARCHAR(20),
    PRIMARY KEY (groupName)
);
    INSERT INTO groups (groupName,groupDescription) VALUES ('abcgroup','groupDescription');

DROP TABLE IF EXISTS whoGroupCreated;
CREATE TABLE IF NOT EXISTS whoGroupCreated(
    userId VARCHAR(30),
    groupName VARCHAR(30),
    FOREIGN KEY (userId) REFERENCES users(userId),
    FOREIGN KEY (groupName) REFERENCES groups(groupName),
    PRIMARY KEY(groupName,userId)
);
INSERT INTO whoGroupCreated (userId,groupName) VALUES ('sagar','abcgroup');

DROP TABLE IF EXISTS userGroupMap;
CREATE TABLE IF NOT EXISTS  userGroupMap(
    groupName VARCHAR(30) NOT NULL,
    userId  VARCHAR(30) NOT NULL,
    FOREIGN KEY (groupName) REFERENCES groups(groupName),
    FOREIGN KEY (userId) REFERENCES users(userId),
    PRIMARY KEY(groupId,userId)
);
DROP TABLE IF EXISTS userGroupMap;
CREATE TABLE IF NOT EXISTS userGroupMap
(
    groupName VARCHAR(30) NOT NULL,
    userId  VARCHAR(30) NOT NULL,
    FOREIGN KEY (groupName) REFERENCES groups(groupName),
    FOREIGN KEY (userId) REFERENCES users(userId),
    PRIMARY KEY(groupName,userId)
);
INSERT INTO whoGroupCreated (groupName,userId) VALUES ('abcgroup','sagar');

DROP TABLE IF EXISTS filesFolderType;
CREATE TABLE IF NOT EXISTS  filesFolderType(
    filesOrFolderId VARCHAR(2) NOT NULL,
    filesFolderTypeInfo VARCHAR(30) NOT NULL,
    PRIMARY KEY (filesOrFolderId)
);
INSERT INTO filesFolderType (filesOrFolderId,filesFolderTypeInfo) VALUES ('f','represent File');
INSERT INTO filesFolderType (filesOrFolderId,filesFolderTypeInfo) VALUES ('d','represent Directory');

DROP TABLE IF EXISTS filesfolder;
CREATE TABLE IF NOT EXISTS  filesfolder(
    -- filesfolderId INT NOT NULL AUTO_INCREMENT,
    filefolderPath VARCHAR(500) NOT NULL,
    filefolderName VARCHAR(30) NOT NULL,
    filesOrFolderId VARCHAR(2) NOT NULL,
    FOREIGN KEY (filesOrFolderId) REFERENCES filesFolderType(filesOrFolderId),
    PRIMARY KEY (filefolderPath,filefolderName,filesOrFolderId)
);
-- Write service to create and delete Files and Folders
INSERT INTO filesfolder (filefolderPath,filefolderName,filesOrFolderId) VALUES ('/','university','d');
INSERT INTO filesfolder (filefolderPath,filefolderName,filesOrFolderId) VALUES ('/','unifile.txt','f');
INSERT INTO filesfolder (filefolderPath,filefolderName,filesOrFolderId) VALUES ('/university/','compsci','d');
INSERT INTO filesfolder (filefolderPath,filefolderName,filesOrFolderId) VALUES ('/university/compsci/','compfile.txt','f');
INSERT INTO filesfolder (filefolderPath,filefolderName,filesOrFolderId) VALUES ('/university/','math','d');
INSERT INTO filesfolder (filefolderPath,filefolderName,filesOrFolderId) VALUES ('/university/math/','mathfile.txt','f');
INSERT INTO filesfolder (filefolderPath,filefolderName,filesOrFolderId) VALUES ('/university/','underuni.txt','f');
INSERT INTO filesfolder (filefolderPath,filefolderName,filesOrFolderId) VALUES ('/university/compsci/','pp','d');

DROP TABLE IF EXISTS permission;
CREATE TABLE IF NOT EXISTS  permission (
permissionValue varchar(20),
permissionValueDesc VARCHAR(50),
PRIMARY KEY(permissionValue)
);
INSERT INTO permission(permissionValue,permissionValueDesc) VALUES ('r','only Read Access');
INSERT INTO permission(permissionValue,permissionValueDesc) VALUES ('w','Read Write  Access');

-- Create/Delete service to read specific user/group permission 
DROP TABLE IF EXISTS groupPermission;
CREATE TABLE IF NOT EXISTS  groupPermission(
   filefolderPath VARCHAR(500) NOT NULL,
   filefolderName VARCHAR(30) NOT NULL,
   filesOrFolderId VARCHAR(2) NOT NULL,
   groupName VARCHAR(30),
   permissionValue varchar(20),
   FOREIGN KEY (filefolderPath,filefolderName,filesOrFolderId) REFERENCES filesfolder(filefolderPath,filefolderName,filesOrFolderId),
   FOREIGN KEY (groupName) REFERENCES groups(groupName),
   FOREIGN KEY (permissionValue) REFERENCES permission(permissionValue),
   PRIMARY KEY(filefolderPath,filefolderName,filesOrFolderId,groupName,permissionValue)
);

INSERT INTO groupPermission(groupName,filefolderPath,filefolderName,filesOrFolderId,permissionValue) VALUES ('g1','/','university','d','w');
INSERT INTO groupPermission(groupName,filefolderPath,filefolderName,filesOrFolderId,permissionValue) VALUES ('g2','/','unifile.txt','f','r');
INSERT INTO groupPermission(groupName,filefolderPath,filefolderName,filesOrFolderId,permissionValue) VALUES ('g3','/university/','compsci','d','w');
INSERT INTO groupPermission(groupName,filefolderPath,filefolderName,filesOrFolderId,permissionValue) VALUES ('g2','/university/compsci/','compfile.txt','f','w');
-- INSERT INTO groupPermission(groupName,filefolderPathName,filesOrFolderId,permissionValue) VALUES ('abc2','/unifile.txt','f','w');

DROP TABLE IF EXISTS userPermission;
CREATE TABLE IF NOT EXISTS  userPermission (
    filefolderPath VARCHAR(500) NOT NULL,
    filefolderName VARCHAR(30) NOT NULL,
    filesOrFolderId VARCHAR(2) NOT NULL,
    userId  VARCHAR(30) NOT NULL,
    permissionValue VARCHAR(20) NOT NULL DEFAULT 'w',
    FOREIGN KEY (filefolderPath,filefolderName,filesOrFolderId) REFERENCES filesfolder(filefolderPath,filefolderName,filesOrFolderId),
    FOREIGN KEY (userId) REFERENCES users(userId),
    FOREIGN KEY (permissionValue) REFERENCES permission(permissionValue),
    PRIMARY KEY(filefolderPath,filefolderName,filesOrFolderId,userId,permissionValue)
);
-- take user of abc2 which give access of '/unifile.txt' as write access
INSERT INTO userPermission(userId,filefolderPath,filefolderName,filesOrFolderId,permissionValue) VALUES ('u2','/','unifile.txt','f','w');
