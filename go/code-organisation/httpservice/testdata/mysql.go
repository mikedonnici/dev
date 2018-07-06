package testdata

const CREATE_MYSQL_DB = `CREATE DATABASE IF NOT EXISTS %s`

const DROP_MYSQL_DB = `DROP DATABASE %s`

const CREATE_MYSQL_TABLE = `CREATE TABLE %s.people (
  id int(9) unsigned NOT NULL AUTO_INCREMENT,
  firstname varchar(100) NOT NULL,
  lastname varchar(100) NOT NULL,
  age tinyint(3) unsigned NOT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;`

const INSERT_MYSQL_DATA = `INSERT INTO %s.people VALUES 
  ('1','Broderick','Reynolds','68'),
  ('2','Osborne','Jast','77'),
  ('3','Dawn','Hayes','61'),
  ('4','Gabriella','McDermott','20'),
  ('5','Declan','Shields','46')`
