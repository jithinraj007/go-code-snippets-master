
- SERVER_ADDRESS    `[IP Address of the machine]`
- SERVER_PORT       `[Port of the machine]`
- DB_USER           `[Database username]`
- DB_PASSWD         `[Database password]`
- DB_ADDR           `[IP address of the database]`
- DB_PORT           `[Port of the database]`
- DB_NAME           `[Name of the database]`

 
1. `resources/database.sql` this contains the SQL for generating the tables. In case you dont want to use the docker-compose file you can use this file to generate tables and insert the default data

# mocks generator
`./generate-mocks.sh`

# run unit tests
  `./run-tests.sh`
