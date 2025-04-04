Install go:
https://go.dev/doc/install

Check version:
go version

Clone Repository then go to inside task service:

cd task-service
go run main.go

cd ../api-gateway
go run main.go

Apis Documentation: http://localhost:8081/swagger/index.html#/


The project structure I shared is a microservices architecture, but it is organised in a single repository (called a monorepo) for development convenience.

Scalability: Explain how the service can be scaled horizontally :
We can:
* Split each service into its own Git repo
* Deploy each with Docker/Kubernetes
Hence, We can run more instances of task-service if needed

Inter-Service Communication: If you were to add another microservice (e.g., a User Service), how would the two services communicate? (Hint: Consider using REST, gRPC, or message queues.)

If we introduce another microservice, such as a user service, the way it communicates with the existing task service depends on the type of interaction and performance requirements.

* Communication Type  |    Task Service Role           |   When to Use
* REST API	        |   Provide	Easy to start with   |    public/internal access
* gRPC                |   Provide	High performance     |    internal only

