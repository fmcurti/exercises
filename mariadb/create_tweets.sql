create table `tweets` (
    `userid` INT not null,
    `text` varchar(280) not null,
    `publication_date` DATE not null,
    FOREIGN KEY(`userid`) REFERENCES Users(userid),
    INDEX(`userid`,`publication_date`)
);