#!/bin/sh


DBHOST=localhost
DBUSER=postgres
DBPASSWD=123456
DBPORT=5434
DBNAME=core

echo "[CORE] Creating db..."

PGPASSWORD=$DBPASSWD psql -h $DBHOST -p $DBPORT -U $DBUSER -d $DBNAME -f database.sql > /dev/null