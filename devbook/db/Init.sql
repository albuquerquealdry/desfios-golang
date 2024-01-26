CREATE TABLE usuarios(id int  auto_increment primary key, nome varchar(50) not null, email varchar(50) not null) ENGINE=INNODB;
CREATE USER 'golang'@'localhost' IDENTIFIED BY 'golang@123Apççs%#'
GRANT ALL PRIVILEGES ON devbook.* TO 'golang'@'localhost';