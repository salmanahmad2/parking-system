#!/bin/bash

DB_USER="carparking"
DB_NAME="carparking"
MIGRATIONS_DIR="./db/migrations"

# Update the psql installation path
PSQL_INSTALLATION_DIR="psql"

# Set the password
PASSWORD="carparking123"

# Sort migration files by version sort and execute them
for FILE in $(ls $MIGRATIONS_DIR/*.up.sql | sort -t'/' -k4 -V); do
    MIGRATION_NUMBER=$(basename $FILE | cut -d '_' -f 1)
    PGPASSWORD=$PASSWORD $PSQL_INSTALLATION_DIR -U $DB_USER -d $DB_NAME -f $FILE
done
