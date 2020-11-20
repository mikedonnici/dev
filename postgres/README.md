# PostgreSQL


Install postgresql and postgis:

```shell script
$ sudo apt-get install postgresql postgresql-contrib postgresql-postgis 
```

Connect as postgres user:

```shell script
$ sudo -u postgres psql 
```

Create db, role, and add postgis extension:

```shell script
# create database foo;
CREATE DATABASE
# create role foouser with password 'foopass';
CREATE ROLE
# alter role foo with LOGIN;
ALTER ROLE
# alter database foo owner to foouser;
ALTER DATABASE  
```

change password:
```
# alter role foo with password 'newpass';
```

### Client Authentication

ref: https://www.postgresql.org/docs/12/client-authentication.html

By default, local connections use `peer` auth, ie, the user needs to exist on 
the local operating system. To allow a 'role' to connect local postgresql server
without having an OS user account, edit `pg_hba.conf`:

```shell script
# Note version here is 12
$ sudo vi /etc/postgresql/12/main/pg_hba.conf
```

...and make this change:

```shell script
# "local" is for Unix domain socket connections only
local   all  all md5 # <- change this from peer to md5
```

Restart server:

```
$ /etc/init.d/postgresql restart
```
