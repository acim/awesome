# Databases

## MySQL

- [17 Key MySQL Config File Settings](http://www.speedemy.com/17-key-mysql-config-file-settings-mysql-5-7-proof/)
- [Example configuration 1](my1.cnf)
- [Example configuration 2](my2.cnf)
- [Example configuration 2](my3.cnf)
- [Semi-sync example master](my-master.cnf)
- [Semi-sync example slave](my-slave.cnf)
- [SSL Configuration Generator](https://ssl-config.mozilla.org/#server=mysql&version=8.0.19&config=modern&openssl=1.1.1k&guideline=5.6)

## PostgreSQL

- [Lesser Known PostgreSQL Features](https://hakibenita.com/postgresql-unknown-features)
- [SSL Configuration Generator](https://ssl-config.mozilla.org/#server=postgresql&version=12.1&config=modern&openssl=1.1.1k&guideline=5.6)
- [UUID, serial or identity columns for PostgreSQL auto-generated primary keys](https://www.cybertec-postgresql.com/en/uuid-serial-or-identity-columns-for-postgresql-auto-generated-primary-keys/)
- [Everything I've seen on optimizing Postgres on ZFS](https://vadosware.io/post/everything-ive-seen-on-optimizing-postgres-on-zfs-on-linux/#dead-end-ulimit)
- [Keep Storage and Backups when Dropping Cluster with PGO](https://blog.crunchydata.com/blog/keep-storage-and-backups-when-dropping-cluster-with-pgo)
- [Five Easy to Miss PostgreSQL Query Performance Bottlenecks](https://pawelurbanek.com/postgresql-query-bottleneck)

### Tips & tricks

- ALTER DATABASE db SET statement_timeout = '60s';
- CREATE EXTENSION pg_stat_statements;
- ALTER DATABASE db SET log_min_duration_statement = '100ms';

## MongoDB

- [Managing schema changes with MongoDB](https://derickrethans.nl/managing-schema-changes.html)
- [mongodump, mongorestore, mongoimport, mongoexport and other tools](https://github.com/mongodb/mongo-tools/)

## Redis

- [SSL Configuration Generator](https://ssl-config.mozilla.org/#server=redis&version=6.0&config=modern&openssl=1.1.1k&guideline=5.6)

## Graph databases

- [Nebula](https://github.com/vesoft-inc/nebula)

## Key-Value stores

- [TiKV - Distributed transactional key-value database](https://github.com/tikv/tikv)
