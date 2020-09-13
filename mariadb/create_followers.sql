create table Followers (
    `userid` INT NOt NULL,
    `follows` INT not null,
    FOREIGN KEY(`userid`) REFERENCES Users(userid),  
    FOREIGN KEY(`follows`) REFERENCES Users(userid),
    INDEX(`userid`)
);