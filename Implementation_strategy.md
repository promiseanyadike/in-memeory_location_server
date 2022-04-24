# Implementtaion Strategy

The project was implemented using the MVC ( Model, View, Controller) Architeccture. Every

- model: This is the domain data. A structured data of type location of type struct with Latitude and Logitude of time float64.

- service: Contains some business logic for the model. Got cache coutesy of the gocache package; which default expiration time of 5 minutes, and which purges expired items every 60 seconds (TTL)

- controllers: I had 3 routes (reportHandler.AddLocation, reportHandler.GetLocation, reportHandler.DeleteLocation). The routes(these actions[apis] are implemented via handler.go file) receive the request from the client as JSON, and in turn calls the service(location.go file) to perform an action for them on the CACHE(cache.go file). All communication between the client(browser) and server are via Golang net/http package.

- congiguration: Contains the connection to the Go server running at port 8080 via HTTP.

- logger: handles log events

All test can be carried out via the client.http file.

go build: in-memory_location_server file

