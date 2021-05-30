CREATE DATABASE `amazon`;

USE `amazon`;

CREATE TABLE `amazon`.`users`(
    `id` INT NOT NULL AUTO_INCREMENT,
    `username` VARCHAR(45) NOT NULL,
    `password` VARCHAR(45) NOT NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE `amazon`.`products`(
    `id` INT NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(45) NOT NULL,
    `price` INT NOT NULL,
    `user_number` INT NOT NULL,
    PRIMARY KEY (`id`),
    INDEX `fk_product_user1_idx` (`user_number` ASC),
    CONSTRAINT `fk_product_user1`
        FOREIGN KEY (`user_number`)
            REFERENCES `amazon`.`users` (`id`)
            ON DELETE NO ACTION
            ON UPDATE NO ACTION
);

/*
to copy and upload csv files:

docker cp localpath labmvc_db_1:.
SET GLOBAL local_infile=1;
quit
mysql --local-infile=1 -u root -p

LOAD DATA LOCAL INFILE 'user.csv'
INTO TABLE `amazon`.`users`
FIELDS TERMINATED BY ','
LINES TERMINATED BY '\n';

LOAD DATA LOCAL INFILE 'product.csv'
INTO TABLE `amazon`.`products`
FIELDS TERMINATED BY ','
LINES TERMINATED BY '\n';
*/