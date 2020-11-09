# 1. Containers

This section contains a simple Go lang webserver that accepts a query parameter, `name`, and responds back `Hello, <name>`.

## Building the docker container
To build the container, simply run:
```
$ docker build -t helloworld:v1 .
```

Next you can run the container and expose the webservers port 8080 on your machines `localhost`:
```
$ docker run -d -p 8080:8080 helloworld:v1
```

You can now test that everything works: 

```
$ curl localhost:8080/hello?name=world
```

The response should be; `Hello, world`