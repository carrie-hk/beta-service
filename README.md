### Building locally

To build locally, run `go run app/main.go -local` from the root directory.

You'll need to include a `.env` file in the root directory with the appropriate environment variables included; that file can be retrieved from Slack. For access to the Slack channel with the `.env` file, please request permission from Elliot or David.

### Building in a Docker container 

To build the Docker image, run `make build`. Then, to start the container, run `make start`.

### Deploying to Google App Engine

To deploy to Google Cloud/Google App Engine, you need to create a Google Cloud account and be given the appropriate set of permissions. For appropriate Google Cloud permissions, please speak to Carrie.

You'll need to include a `app.yaml` file in the root directory with the appropriate environment variables included; that file can be retrieved from Slack. For access to the Slack channel with the `app.yaml` file, please request permission from Elliot or David.

To initialize the deployment process, run `gcloud init`. From there, select the appropriate configuration you'd like to deploy to. Once you've configured your build, run `gcloud deploy`.



