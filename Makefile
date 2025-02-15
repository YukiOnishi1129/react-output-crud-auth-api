include .env

empty:
	echo "empty"

# enter the container
backend-ssh:
	docker exec -it $(BACKEND_CONTAINER_NAME) sh

db-sh:
	docker exec -it $(POSTGRES_HOST) sh

# psql -h 127.0.0.1 -p 5432 -U user REACT_OUTPUT_CRUD_AUTH_API_DB