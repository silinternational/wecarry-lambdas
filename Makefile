build:
	docker-compose run app ./codeship/build.sh

shell:
	docker-compose run app bash

deploy:
	docker-compose run app bash -c "cd cron/maintenance && sls deploy -v"

remove:
	docker-compose run app sls remove

clean:
	docker-compose kill
	docker-compose rm -f
