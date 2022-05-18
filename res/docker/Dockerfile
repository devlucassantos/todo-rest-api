FROM postgres:13.2

WORKDIR /mydata

COPY ./fixtures ./fixtures

EXPOSE 5432

ADD ./sql/init_db.sql /docker-entrypoint-initdb.d
ADD ./sql/load_fixtures.sql /docker-entrypoint-initdb.d
