build-image:
	docker build . -t kong_example_app

create-container:
	docker run -p 3000:3000 --name kong_example_app -dit kong_example_app
