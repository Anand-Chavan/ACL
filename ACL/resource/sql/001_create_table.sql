use acl;


CREATE TABLE  users
(
    Name VARCHAR(30) NOT NULL,
    Id  INT NOT NULL AUTO_INCREMENT,
    Password  VARCHAR(20) NOT NULL,
    DateCreation VARCHAR(30) NOT NULL,
    UserType VARCHAR(30) NOT NULL,
    PRIMARY KEY (Id)
);

CREATE TABLE  sessionInfo
(
    Id  INT NOT NULL AUTO_INCREMENT,
    sessionId VARCHAR(20) NOT NULL,
    logintime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ,
    PRIMARY KEY (Id,sessionId)
);

CREATE TABLE  groups
(
    groupId  INT NOT NULL AUTO_INCREMENT,
    groupName  VARCHAR(30) NOT NULL,
    groupCreationDate  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ,
    groupDescription VARCHAR(20),
    PRIMARY KEY (groupId)
);

CREATE TABLE  whoGroupCreated
(
    Id  INT NOT NULL AUTO_INCREMENT,
    groupId INT,
    FOREIGN KEY (Id) REFERENCES users(Id),
    FOREIGN KEY (groupId) REFERENCES groups(groupId),
    PRIMARY KEY(groupId,Id)
);

CREATE TABLE  userGroupMap
(
    groupId  INT,
    Id  INT NOT NULL AUTO_INCREMENT,
    FOREIGN KEY (groupId) REFERENCES groups(groupId),
    FOREIGN KEY (Id) REFERENCES users(Id),
    PRIMARY KEY(groupId,Id)
);





CREATE TABLE  content(
    contentId int,
    contentName varchar(20),
    contentInfo varchar(2),
    PRIMARY KEY (contentId)
);

CREATE TABLE permission(
permissionValue varchar(20),
PRIMARY KEY(permissionValue)
);


CREATE TABLE  userPermission(
    Id  INT NOT NULL AUTO_INCREMENT,
    contentId int,
    permissionValue varchar(20),
    FOREIGN KEY (Id) REFERENCES users(Id),
    FOREIGN KEY (contentId) REFERENCES content(contentId),
    FOREIGN KEY (permissionValue) REFERENCES permission(permissionValue),
    PRIMARY KEY(Id,contentId,permissionValue)
);

CREATE TABLE  groupPermission(
    groupId  INT,
    contentId int,
    permissionValue varchar(20),
    FOREIGN KEY (groupId) REFERENCES groups(groupId),
    FOREIGN KEY (contentId) REFERENCES content(contentId),
    FOREIGN KEY (permissionValue) REFERENCES permission(permissionValue),
    PRIMARY KEY(groupId,contentId,permissionValue)
);
