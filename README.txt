We are currently utilizing self-signed certificates for TLS/HTTPS. As such, you will receive warnings about an insecure connection when attempting to make requests to the server. 
If you see this in the browser, go to the Advanced options and proceed anyway. To test the server with the curl command, you must use the -k flag to skip checking the TLS certification or the request will not go through.

To build the docker image, run "make build". Then to start the container, run "make start".

The service is avaliable at: 
https://localhost:443/assets/all
https://localhost:443/assets/featured

