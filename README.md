# To Do List - REST API

Hello, astronauts!

In this repository you will find an example of a REST API made in Go with the To Do List theme. This application was
built to improve my skills in Go and REST APIs, and now I would like to share the knowledge acquired with you.
This repository is still closely monitored by the creator, so feel free to contribute to the repository by reporting
problems or suggesting new features. If this project is useful to you, consider giving the repository a star.

## The Beginning of the Journey

### Step 1 - Way

You will have two ways to take advantage of all the features of the repository. The first way is the simplest, you just
need to download this repository and have [Docker](https://www.docker.com/get-started/) installed on your machine.
Having met all the previous requirements, just enter the project folder and run the following command and the entire API
will be running without needing any further configuration:

```shell
$ docker-compose up --build
```

However, if you are a developer and want to make some modification to the project, I advise you to use another container
that is also available in the project and that will only have the database. So, to make the API work on your machine you
will first need to have [Go](https://go.dev/dl/) installed, in addition to this repository and Docker already mentioned
above. With all the requirements met and inside the project folder, to run the API just run the following commands:

```shell
$ cd res/docker
$ docker-compose up --build --detach
$ cd ../..
$ go run main.go
```

### Step 2 - Settings

You can easily modify some API settings, such as where your server can be located or which Postgres database it will
access, for example. For that, just modify the values present in the [environment configuration file](.env) and
everything will be as you defined.

### Step 3 - Map

You will be able to see how it works and test the execution of all the API routes using only your browser. If you are
using the project's default values for `SERVER_ADDRESS` and `SERVER_PORT`, just type the following address in your
browser's search bar (Note that if you change these values you will need to update this search address putting the
chosen values for address and port):

```
localhost:8000/api/documentation/index.html
```
