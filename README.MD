## A basic message queue inside of GoLang

# Basic Request Examples
>curl -X POST http://localhost:8080/enqueue \
  -H "Content-Type: application/json" \
  -d '{"queueName": "TaskQueue", "processId": 1, "processStatus": "Pending"}'

>curl -X GET http://localhost:8080/health

# TODO:
1. Implement listener
2. Cleanup routes
3. Implement docker file
4. Work more on deserializer