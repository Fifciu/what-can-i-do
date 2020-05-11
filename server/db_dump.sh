#!/bin/bash
if [ -f .env ]; then
    # Load Environment Variables
    export $(cat .env | grep -v '#' | awk '/=/ {print $1}')
    # For instance, will be example_kaggle_key
    echo $KAGGLE_KEY
fi
docker exec server_mariadb_1 /usr/bin/mysqldump -u $db_user --password=$db_pass $db_name > backup.sql;