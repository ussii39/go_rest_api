FROM --platform=linux/amd64 mysql:8.0
# FROM mysql:8.0

# ENV HOSTNAME [mysql]

COPY ./init/00_create_db.sql /docker-entrypoint-initdb.d
COPY ./init/test/00_create_db.sql /docker-entrypoint-initdb.d

# COPY ./init/01_create_table.sql /docker-entrypoint-initdb.d
# CMD [ "mysqld" "--character-set-server=utf8mb4" "--collation-server=utf8mb4_unicode_ci"]
# COPY ./data /var/lib/mysql