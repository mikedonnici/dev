# PostgreSQL


Install postgresql and postgis:

```bash script
$ sudo apt-get install postgresql postgresql-contrib postgresql-postgis 
```

Connect as postgres user:

```bash script
$ sudo -u postgres psql 
```

Create db, role, and add postgis extension:

```bash script
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

```bash script
# Note version here is 12
$ sudo vi /etc/postgresql/12/main/pg_hba.conf
```

...and make this change:

```bash script
# "local" is for Unix domain socket connections only
local   all  all md5 # <- change this from peer to md5
```

Restart server:

```
$ /etc/init.d/postgresql restart
```


## Drop database when `database "foo" is being accessed by other users`

From [here](https://stackoverflow.com/questions/17449420/postgresql-unable-to-drop-database-because-of-some-auto-connections-to-db)

```sql
REVOKE CONNECT ON DATABASE foo FROM public;

-- May nbeed to run this a few times?
SELECT pg_terminate_backend(pg_stat_activity.pid)
FROM pg_stat_activity
WHERE pg_stat_activity.datname = 'foo';

DROP database foo;
```
