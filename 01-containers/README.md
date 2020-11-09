# Containers

This section contains a simple Go lang webserver that accepts a query parameter, name, and responds back `Hello, <name>`.

## Building a container
To build a container, simply run:
```
docker build -t helloworld .
```

To run the container:
```
docker run -d -p 8080:8080 helloworld
```

This will run the container and expose port 8080 on `localhost`.

You can now test that everything works: 

```
curl localhost:8080/hello?name=world
```

The response should be; `Hello, world`