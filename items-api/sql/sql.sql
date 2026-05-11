CREATE DATABASE IF NOT EXISTS argentina;
CREATE DATABASE IF NOT EXISTS brasil;
CREATE DATABASE IF NOT EXISTS colombia;
CREATE DATABASE IF NOT EXISTS mexico;

USE brasil;
CREATE TABLE IF NOT EXISTS items(
    id int auto_increment primary key,
    Product varchar(10) not null,
    Name varchar(50) not null unique,
    Category varchar(50) not null unique,
    Token varchar(20) not null unique,
    CriadoEm timeStamp default current_timestamp()
) ENGINE=INNODB;

USE argentina;
CREATE TABLE IF NOT EXISTS items(
    id int auto_increment primary key,
    Product varchar(10) not null,
    Name varchar(50) not null unique,
    Category varchar(50) not null unique,
    Token varchar(20) not null unique,
    CriadoEm timeStamp default current_timestamp()
) ENGINE=INNODB;

USE colombia;
CREATE TABLE IF NOT EXISTS items(
    id int auto_increment primary key,
    Product varchar(10) not null,
    Name varchar(50) not null unique,
    Category varchar(50) not null unique,
    Token varchar(20) not null unique,
    CriadoEm timeStamp default current_timestamp()
) ENGINE=INNODB;

USE mexico;
CREATE TABLE IF NOT EXISTS items(
    id int auto_increment primary key, 
    Product varchar(10) not null, 
    Name varchar(50) not null unique,
    Category varchar(50) not null unique,
    Token varchar(20) not null unique,
    CriadoEm timeStamp default current_timestamp()
) ENGINE=INNODB;
