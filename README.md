# heroku-go-embed-template

## Start a new project:

Run the command `make new-project name=PROJECT-NAME-HERE`. This command replaces the project name and all references.

## Run locally

- Run the db with the command `make env-up`.
- then run `make run`, it will build the `ui` and run the main.go file, repeat this step everytime you change something in the backend project.

## Run backend tests

- Run the command `make tests`.

## Run frontend tests [TODO]

## Deployment

Uncomment the deploy step in .github/workflows/ci.yml and configure `heroku_api_key` and `heroku_app_name`.

# Set HEROKU_API_KEY and project name [TODO]
