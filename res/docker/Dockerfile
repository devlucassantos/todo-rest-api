FROM postgres

COPY sql/create_service_user.sql /docker-entrypoint-initdb.d
COPY sql/init_db.sql /docker-entrypoint-initdb.d
COPY sql/load_fixtures.sql /docker-entrypoint-initdb.d
COPY fixtures /fixtures

ARG POSTGRES_DB=$POSTGRES_DB
ARG POSTGRES_SERVICE_USER=$POSTGRES_SERVICE_USER
ARG POSTGRES_SERVICE_USER_PASSWORD=$POSTGRES_SERVICE_USER_PASSWORD

RUN sed -i -e 's/$POSTGRES_DB/'"$POSTGRES_DB"'/g' \
    -e 's/$POSTGRES_SERVICE_USER_PASSWORD/'"$POSTGRES_SERVICE_USER_PASSWORD"'/g' \
    -e 's/$POSTGRES_SERVICE_USER/'"$POSTGRES_SERVICE_USER"'/g' /docker-entrypoint-initdb.d/create_service_user.sql


EXPOSE 5432
