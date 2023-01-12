build:
	docker-compose up -d app

bash:
	docker-compose run --rm app bash

deploy:
	docker-compose run --rm app bash -c "cd cron/maintenance && sls deploy --verbose --stage stg"

remove:
	docker-compose run --rm app sls remove

gosec:
	docker-compose run --rm gosec

test: gosec
	docker-compose run --rm app ./codeship/test.sh

clean:
	docker-compose kill
	docker-compose rm -f
