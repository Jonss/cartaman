# Cartaman

## How to run the project?

There's two commands to run the project, one using your environment and another using only docker.
You can run the commands `make env-up` to start the dependencies (database only ), then run `make run` to start the application locally. Runs on port 8082.
You can also use the command `make run-docker` which starts the database and the app into a docker container. Runs on port 9000.

There's a docs directory which contains the collection of requests and environments to be used on docker.

## How to run the tests?

You can use the command `make test` to check the result on terminal and `make cover` to check the test coverage as a html file.

### Comments

I decided to use clean architecture, the `usecase` package doesn't have access to implementations, only abstractions. The `ports` contains the inputs to the application, a REST API on this project - but could easily adapted for graphql or gRPC for instance.
The `adapters` package contains the "output", I used a postgresql but could be a file or another type of database.

I didn't tested the usecase layer, because the `adapters/repository` was extensively tested using test-containers. Same for `ports/httprest`.
This project is based on this [article](https://threedots.tech/post/introducing-clean-architecture).

I had no time to add logs, but if I had I would add a set of functions wrapping [zap](https://github.com/uber-go/zap).


