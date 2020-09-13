create table `Users` (
    `userid` INT AUTO_INCREMENT not null,
    `username` varchar(255) NOT NULL,
    `passwd` VARCHAR(500) NOT NULL,
    `first_name` varchar(255),
    `last_name` varchar(255),
    PRIMARY KEY (`userid`),
    UNIQUE KEY (`username`),
    INDEX(`userid`)
);
