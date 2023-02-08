# GoResearch
## How to run api
+ Run go mod tidy to install neccessary packages
+ Install Make, Docker(simple database management) or database server
+ Execute command "make setup" to copy example configuration file
+ With docker user: 
    * Comment the services that are not desired (only one is actives)
    * Execute command "make start-docker" to start database
    * Execute command "make stop-docker" to stop database
+ Execute command "make migrate-db" to migrate table and seed data
+ Execute command "make start-server" to start the server
    * Add option cfg=yml or pkg=fiber to use desired config or pkg
+ Execute command "make clear-log" to clear log file

## Todo
### Routers
+ Mux: Done
+ Echo: Done
+ Fiber: Done

### Database
+ PostgreSQL: Done
+ MySQL: Done
+ MongoDB: In-progress

### Logging
+ Zap: Done

### Authentication
+ JWT: In-progress
+ Oauth: In-progress

### Message Queue
+ RabbitMQ: In-progress