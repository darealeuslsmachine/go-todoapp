Patterns:
-
1. DDD
2. Middleware
3. HTTP server graceful shutdown
4. DB connection Pool

Start application:
-
1. `make env-up` - init env
2. `make migrate-up` - setup db/migrations 
3. `make env-port-forward` - optional, make db visible from outside
4. `make todoapp-run` - run the application

PS. only for local deploying, host is hardcoded in Makefile