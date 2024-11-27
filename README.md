# To Do List - REST API

🌍 *[English](README.md) ∙ [Português](README_pt.md)*

In this repository you will find an example of a REST API made in Go with the To Do List theme. This application was
built to improve my skills in Go and REST APIs, and now I would like to share the knowledge acquired with you.
This repository is still closely monitored by the creator, so feel free to contribute to the repository by reporting
problems or suggesting new features. If this project is useful to you, consider giving the repository a star.

## Important concepts applied in this repository

* REST API with simple authentication using JWT token
* CRUD using PostgreSQL and Go
* Hexagonal architecture
* Dockerization

## How to run

### Prerequisites

You will have two ways to take advantage of all the features of this repository. The first way is the simplest, you just
need to download this repository and have [Docker](https://www.docker.com/get-started/) installed on your machine.
Once you've done that, go to the project folder and run the command below, after which the API will be running without
needing any additional configuration:

```shell
$ docker-compose up
```

However, if you are a developer and want to make some modification to the project, I advise you to use another container
that is also available in the project and that will only have the database. So, to make the API work on your machine you
will first need to have [Go](https://go.dev/dl/) installed, in addition to this repository and Docker already mentioned
above. With all the previous requirements met and inside the project folder, to run the API just run the following
commands:

```shell
$ cd res/docker
$ docker-compose up --detach
$ cd ../..
$ go run main.go
```

### Documentation

You will be able to  test the execution of all the API routes using only your browser. If you are using the project's
default values for `HOST` and `PORT`, just type the following address in your browser's search bar (Note
that if you change these values you will need to update this search address putting the chosen values for the host and
port):

```
localhost:8000/api/documentation/index.html
```

Note that to use all the API routes you will need to be authenticated, which is a simple process to do. To do this, look
for the **Sign Up** and **Sign In** routes. In these routes, you'll find all the information you need to register and
log in to the API. After making one of the previous requests, you will receive your `access_token` in the body of the
response. The value of this access token must be copied and pasted into the field that appears when you click on the
**Authorize** button above all the routes on the left side, as you can see in the image below.

![todo-rest-api](https://user-images.githubusercontent.com/89457923/169172172-1c112bf0-14d0-43c2-89d9-ba52c8391ac2.png)

### Settings

You can easily modify some API settings, such as the address of your server or the PostgreSQL database it will access,
for example. For that, just modify the values present in the [environment configuration file](.env) and everything will
be as you defined.
