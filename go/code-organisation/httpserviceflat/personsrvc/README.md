# personsrvc

An http service as a single package.

Here I've avoided the temptation of splitting everything into sub packages
without good reason.

This package would provide a specific _service_ and could, therefore, be
utilised by other services.

I am trying this after looking at [gokit](https://github.com/go-kit/kit) - which looks awesome but I found hard to wrap my head around.

There are others, goa, micro, gizmo.

Here are the endpoints:

- `/` -
- `/arg` -
- `/person/{id}`
- `/person/oid/{oid}`
- `/people`
- `/mwtrue` - runs middleware
- `/mwfalse` - skips middleware

Set up MySQL:

```sql
CREATE DATABASE IF NOT EXISTS peoplesrvc;
CREATE TABLE peoplesrvc.people (
  id int(9) unsigned NOT NULL AUTO_INCREMENT,
  firstname varchar(100) NOT NULL,
  lastname varchar(100) NOT NULL,
  age tinyint(3) unsigned NOT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO peoplesrvc.people VALUES
  ('1','Broderick','Reynolds','68'),
  ('2','Osborne','Jast','77'),
  ('3','Dawn','Hayes','61'),
  ('4','Gabriella','McDermott','20'),
  ('5','Declan','Shields','46');
```

Set up MongoDB:

```sh
use peoplesrvc;
db.people.insert({"_id": ObjectId("5b3bcd72463cd6029e04de18"), "id" : 1, "firstname" : "Broderick", "lastname" : "Reynolds", "age" : 68});
db.people.insert({"_id": ObjectId("5b3bcd72463cd6029e04de1a"), "id" : 2, "firstname" : "Osborne",   "lastname" : "Jast", "age" : 77});
db.people.insert({"_id": ObjectId("5b3bcd72463cd6029e04de1c"), "id" : 3, "firstname" : "Dawn", "lastname" : "Hayes", "age" : 61});
db.people.insert({"_id": ObjectId("5b3bcd72463cd6029e04de1e"), "id" : 4, "firstname" : "Gabriella", "lastname" : "McDermott","age" : 20});
db.people.insert({"_id": ObjectId("5b3bcd72463cd6029e04de20"), "id" : 5, "firstname" : "Declan",    "lastname" : "Shields",  "age" : 46});
```

To run needs a `.env` file:

```env
MYSQL_DSN="root:password@tcp(localhost:3306)/"
MYSQL_DBNAME="peoplesrvc"
MONGO_DSN="mongodb://localhost"
MONGO_DBNAME="peoplesrvc"
```

go run main.go

Todo:

- [ ] Add logging
