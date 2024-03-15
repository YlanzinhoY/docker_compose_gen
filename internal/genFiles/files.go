package genfiles

func Postgres() string {

	dockerfile := `version: '3'
	services:
		postgres_prod:
			image: postgres:latest
			environment:
				POSTGRES_USER: postgres
				POSTGRES_PASSWORD: postgres
				POSTGRES_DB: postgres
			ports:
				- "5432:5432"
		postgres_test:
			image: postgres:latest
			environment:
				POSTGRES_USER: postgres_test
				POSTGRES_PASSWORD: postgres_test
				POSTGRES_DB: postgres_test
			ports:
				- "5433:5432"
	`

	return dockerfile

}

func Mysql() string {
	dockerfile := `version: '3'
	services:
		mysql_prod:
			image: mysql:latest
			environment:
				MYSQL_ROOT_PASSWORD: rootpassword
				MYSQL_DATABASE: mysqldb
				MYSQL_USER: mysqluser
				MYSQL_PASSWORD: mysqlpassword
			ports:
				- "3306:3306"
		mysql_test:
			image: mysql:latest
			environment:
				MYSQL_ROOT_PASSWORD: rootpassword
				MYSQL_DATABASE: mysqldb_test
				MYSQL_USER: mysqluser_test
				MYSQL_PASSWORD: mysqlpassword_test
			ports:
				- "3307:3306"
	`

	return dockerfile
}

func MongoDB() string {
	dockerfile := `version: '3'
	services:
		mongodb_prod:
			image: mongo:latest
			environment:
				MONGO_INITDB_ROOT_USERNAME: root
				MONGO_INITDB_ROOT_PASSWORD: example
			ports:
				- "27017:27017"
		mongodb_test:
			image: mongo:latest
			environment:
				MONGO_INITDB_ROOT_USERNAME: root_test
				MONGO_INITDB_ROOT_PASSWORD: example_test
			ports:
				- "27018:27017"
	`

	return dockerfile
}
