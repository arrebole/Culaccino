/var/lib/postgresql/data/pgdata

docker run 
	--name some-postgres 
	-e POSTGRES_PASSWORD=e8y28diq9hx
	-e POSTGRES_USER=postgres
	-v /var/lib/postgresql/data/pgdata:/var/lib/postgresql/data
	-p 5432:5432
	-d postgres