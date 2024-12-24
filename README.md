### Task
Takes in a JSON receipt (see example in the example directory) and returns a JSON object with an ID generated by your code.

The ID returned is the ID that should be passed into /receipts/{id}/points to get the number of points the receipt was awarded.

How many points should be earned are defined by the rules below.

Reminder: Data does not need to survive an application restart. This is to allow you to use in-memory solutions to track any data generated by this endpoint.

### Libraries
Docker, Gin, Swagger, Grafana

### Install dependencies
Please ensure you installed Docker on your machine.
```bash
make install
```

### Generate Swagger UI
If you don't see `docs` folder in your project, you can run the below command.
```bash
swag init
```

### Usage
Run the following command to start the application locally.

```bash
make run
```

### Endpoints
- http://localhost:5000/receipts/process
- http://localhost:5000/receipts/{id}/points

### Swagger UI
You can test the APIs through
http://localhost:5000/swagger/index.html
