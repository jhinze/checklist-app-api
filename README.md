# checklist-app-api
API for Checklist webapp

## Required Environment Variables
`DATABASE_HOST=<database URL>`  
`DATABASE_PORT=<port>`  
`DATABASE_DBNAME=<database name>`  
`DATABASE_USER=<user name>`  
`DATABASE_PASSWORD=<password>`  
`DATABASE_SSL_MODE=<enable|disable>`  
`DATABASE_DIALECT=<postgres|mysql|etc..>` (if other than postgres it must be imported in main.go)  

# Optional Environment Variables
`DATABASE_LOGMODE=true` (enable more verbose query logging)

## Docker
```
docker run \
  --network=<where the database is accessible> \
  -p <port>:80 \
  -e DATABASE_HOST=<database URL> \ 
  -e DATABASE_PORT=<port> \
  -e DATABASE_DBNAME=<database name> \
  -e DATABASE_USER=<user name> \
  -e DATABASE_PASSWORD=<password> \
  -e DATABASE_SSL_MODE=<enable|disable> \ 
  -e DATABASE_DIALECT=<postgres|mysql|etc..>  \
  jhinze/checklist-app-api
```

