build:
	docker build -t beta_service .    

start:
	 docker run -d -p 5001:5000 beta_service