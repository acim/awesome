# Databases

- [DB Fiddle - PostgreSQL, MySQL and SQLite playground](https://www.db-fiddle.com/)

## MySQL

- [17 Key MySQL Config File Settings](http://www.speedemy.com/17-key-mysql-config-file-settings-mysql-5-7-proof/)
- [Example configuration 1](my1.cnf)
- [Example configuration 2](my2.cnf)
- [Example configuration 2](my3.cnf)
- [Semi-sync example master](my-master.cnf)
- [Semi-sync example slave](my-slave.cnf)
- [SSL Configuration Generator](https://ssl-config.mozilla.org/#server=mysql&version=8.0.19&config=modern&openssl=1.1.1k&guideline=5.6)

## PostgreSQL

- [Awesome Postgres](https://github.com/dhamaniasad/awesome-postgres)
- [Lesser Known PostgreSQL Features](https://hakibenita.com/postgresql-unknown-features)
- [SSL Configuration Generator](https://ssl-config.mozilla.org/#server=postgresql&version=12.1&config=modern&openssl=1.1.1k&guideline=5.6)
- [UUID, serial or identity columns for PostgreSQL auto-generated primary keys](https://www.cybertec-postgresql.com/en/uuid-serial-or-identity-columns-for-postgresql-auto-generated-primary-keys/)
- [Everything I've seen on optimizing Postgres on ZFS](https://vadosware.io/post/everything-ive-seen-on-optimizing-postgres-on-zfs-on-linux/#dead-end-ulimit)
- [Keep Storage and Backups when Dropping Cluster with PGO](https://blog.crunchydata.com/blog/keep-storage-and-backups-when-dropping-cluster-with-pgo)
- [Five Easy to Miss PostgreSQL Query Performance Bottlenecks](https://pawelurbanek.com/postgresql-query-bottleneck)
- [PostgreSQL execution plan visualizer](https://explain.dalibo.com/)
- [Blazing-fast cloning of PostgreSQL databases](https://github.com/postgres-ai/database-lab-engine)
- [Postgres indexes for newbies](https://blog.crunchydata.com/blog/postgres-indexes-for-newbies)
- [Read-only mode for PostgreSQL](https://www.jkatz05.com/post/postgres/postgres-read-only/)
- [Reshape - easy-to-use, zero-downtime schema migration tool for Postgres](https://github.com/fabianlindfors/reshape)
- [pgroll - PostgreSQL zero-downtime migrations made easy](https://github.com/xataio/pgroll)
- [Vacuuming update-heavy tables](https://dataegret.com/2022/02/vacuuming-update-heavy-tables/)
- [PostgreSQL's Powerful New Join Type: LATERAL](https://heap.io/blog/postgresqls-powerful-new-join-type-lateral)
- [Hubert Lubaczewski's Blog](https://www.depesz.com/)
- [OrioleDB – modern cloud-native storage engine](https://github.com/orioledb/orioledb)
- [How to build a job queue with Rust and PostgreSQL](https://kerkour.com/rust-job-queue-with-postgresql)
- [How we optimized PostgreSQL queries 100x](https://towardsdatascience.com/how-we-optimized-postgresql-queries-100x-ff52555eabe)
- [Optimize PostgreSQL Server Performance Through Configuration](https://blog.crunchydata.com/blog/optimize-postgresql-server-performance)
- [Why isn't Postgres using my index?](https://www.pgmustard.com/blog/why-isnt-postgres-using-my-index)
- [Our Experience with PostgreSQL on ZFS](https://lackofimagination.org/2022/04/our-experience-with-postgresql-on-zfs/)
- [Generate PostgreSQL credentials with Hashicorp Vault and Go](https://splitmind.dev/posts/generate-creds-postgres-vault-with-golang/)
- [Use PostgreSQL SSL connection in rust with self-signed certificates](https://dev.to/yugabyte/use-postgresql-ssl-connection-in-rust-with-self-signed-certificates-4k9g)
- [Find and fix a missing PostgreSQL index](https://www.cybertec-postgresql.com/en/find-and-fix-a-missing-postgresql-index/)
- [8 Fascinating Things You Probably Didn't Know PostgreSQL Can Do!](https://postgresweekly.com/link/122726/web)
- [Storing Network Addresses using PostgreSQL](https://www.compose.com/articles/storing-network-addresses-using-postgresql/)
- [How to Find and Stop Running Queries on PostgreSQL](https://adamj.eu/tech/2022/06/20/how-to-find-and-stop-running-queries-on-postgresql/)
- [Enabling and Enforcing SSL/TLS for PostgreSQL Connections](https://www.percona.com/blog/enabling-and-enforcing-ssl-tls-for-postgresql-connections/)
- [Choosing a Postgres primary key](https://supabase.com/blog/choosing-a-postgres-primary-key)
- [How to JSON in PostgreSQ](https://ftisiot.net/postgresqljson/main/)
- [Nine ways to shoot yourself in the foot with PostgreSQL](https://philbooth.me/blog/nine-ways-to-shoot-yourself-in-the-foot-with-postgresql)
- [The Unexpected Find That Freed 20GB of Unused Index Space](https://hakibenita.com/postgresql-unused-index-size)
- [How to Get the Most out of Postgres Memory Settings](https://tembo.io/blog/optimizing-memory-usage)
- [Postgres as a search engine](https://anyblockers.com/posts/postgres-as-a-search-engine)
- [Optimizing Postgres table layout for maximum efficiency](https://r.ena.to/blog/optimizing-postgres-table-layout-for-maximum-efficiency/)

### Tools

- [pgwatch2 - flexible self-contained PostgreSQL metrics monitoring solution](https://github.com/cybertec-postgresql/pgwatch2)
- [explain.depesz.com - PostgreSQL's explain analyze made readable](https://explain.depesz.com/)
- [pg_back - dump tool for PostgreSQL](https://github.com/orgrim/pg_back)
- [PGTune - tool for postgresql.conf generation with optimal parameters](https://pgtune.leopard.in.ua/#/)
- [pgbadger - fast PostgreSQL log analyzer](https://github.com/darold/pgbadger)
- [pgcopydb - pg_dump | pg_restore on steroids](https://github.com/dimitri/pgcopydb)
- [PGLoader - data loading tool for PostgreSQL using the COPY command](https://github.com/dimitri/pgloader)
- [Transporter - database transformations from one store to another (supports ElasticSearch, MongoDB, PostgreSQL, RabbitMQ, MySQL)](https://github.com/compose/transporter)
- [pg_rman - online backup and restore tool for PostgreSQL](https://github.com/ossc-db/pg_rman)
- [Barman - Backup and recovery manager for PostgreSQL](https://github.com/EnterpriseDB/barman)
- [EverSQL - online SQL query optimization tool](https://www.eversql.com/sql-query-optimizer/)
- [PGSync - PostgreSQL to Elasticsearch sync](https://github.com/toluaina/pgsync)
- [pg_netstat - PostgreSQL database network traffic monitor](https://github.com/supabase/pg_netstat)
- [Pigsty - battery-included free RDS alternative](https://github.com/Vonng/pigsty)
- [PgHero - performance dashboard for Postgres](https://github.com/ankane/pghero)

### AI

- [An early look at HNSW performance with pgvector](https://jkatz05.com/post/postgres/pgvector-hnsw-performance/)
- [pgvector 0.5.0 feature highlights and howtos](https://jkatz05.com/post/postgres/pgvector-overview-0.5.0/)
- [pgai - allows you to develop RAG, semantic search, and other AI applications directly in PostgreSQL](https://github.com/timescale/pgai)
- [pgvectorscale - complement to pgvector for high performance, cost efficient vector search on large workloads](https://github.com/timescale/pgvectorscale)
- [Timescale - AI and pgvector](https://www.youtube.com/playlist?list=PLsceB9ac9MHR7IL2kSiHN8NUCmXoEEAf8)
- [Timescale Blog](https://www.timescale.com/blog/tag/ai/)

### Video tutorials

- [Citus Con: An Event for Postgres 2022](https://www.youtube.com/playlist?list=PLlrxD0HtieHjSzUZYCMvqffEU5jykfPTd)

### Tips & tricks

- ALTER DATABASE db SET statement_timeout = '60s';
- CREATE EXTENSION pg_stat_statements;
- ALTER DATABASE db SET log_min_duration_statement = '100ms';

### Install from apt

```sh
wget -qO - https://www.postgresql.org/media/keys/ACCC4CF8.asc | gpg --dearmor | sudo dd of=/usr/share/keyrings/postgresql-archive-keyring.gpg

echo 'deb [ arch=amd64 signed-by=/usr/share/keyrings/postgresql-archive-keyring.gpg ] http://apt.postgresql.org/pub/repos/apt jammy-pgdg main' | sudo tee /etc/apt/sources.list.d/postgresql.list

sudo apt-get update
sudo apt-get install postgresql-client
```

## Elasticsearch alternatives

- [zinc - lightweight alternative to elasticsearch that requires minimal resources](https://github.com/zinclabs/zinc)
- [sonic - fast, lightweight and schema-less search backend](https://github.com/valeriansaliou/sonic)
- [pg_bm25 - Elastic quality full text search inside Postgres](https://github.com/paradedb/paradedb)
- [bleve - modern text indexing library](https://github.com/blevesearch/bleve)

## MongoDB

- [Managing schema changes with MongoDB](https://derickrethans.nl/managing-schema-changes.html)
- [mongodump, mongorestore, mongoimport, mongoexport and other tools](https://github.com/mongodb/mongo-tools/)

## Redis

- [SSL Configuration Generator](https://ssl-config.mozilla.org/#server=redis&version=6.0&config=modern&openssl=1.1.1k&guideline=5.6)
- [Dragonfly - faster Redis and Memcached compatible store](https://github.com/dragonflydb/dragonfly)

## Graph databases

- [Nebula](https://github.com/vesoft-inc/nebula)

## Key-Value stores

- [TiKV - Distributed transactional key-value database](https://github.com/tikv/tikv)
