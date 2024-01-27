test:
	docker build -t testing_my_image:latest .
	docker run testing_my_image:latest "Shaxx"
