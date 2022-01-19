build:
	docker build -t cahk/betaservice .    

start:
	docker run -d -p 5001:5004 cahk/betaservice