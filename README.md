# TransacAI Workload Manager Service

This project is the codebase for the Workload Manager Service (WMS) of the TransacAI project.

## TransacAI

TransacAI project is geared towards generation of enriched summaries and insights of transactional data in real-time or batch using Generative AI and Large Language Models (LLMs). It goes beyond visual and analytical processing of transactional data by generating context-aware and enriched insights in natural language using LLMs. It focuses on delivering human-centric analysis that is easier to understand and act upon, eliminating the need for multiple complex data processing steps to derive insights from raw data.

## Workload Manager Service (WMS)

Workload Manager Service (WMS) is one of the core services of the TransacAI project which serves at the forefront of the insights generation pipeline. It not handles manual requests for insights generation, it is also responsible for managing scheduled insights generation tasks. It is designed to be scalable and fault-tolerant, and can be deployed in a distributed environment.

### Connect RPC / gRPC

The WMS service primarily uses the [Connect](https://connectrpc.com/docs/introduction) libraries for defining the service and handling the RPC communication. The service is defined in the `proto/wms/v1/transac_ai_wms.proto` file.

Once primary advantage of using Connect is interoperability. Connect supports multiple languages and platforms, and the service can handle requests over HTTP/1.1, HTTP/2, and HTTP/3, i.e., it supports gRPC, gRPC-Web, and Connect Protocol right out of the box. This makes it much easier to integrate WMS with other micro-services in the TransacAI project (which use gRPC) and with the front-end applications (which use Connect Protocol for better performance using binary-encoded Protobuf for data transmission).

### Distributed Processing

The WMS service mainly interacts with the RSS (Requests Storage Service) and the IGS (Insights Generation Service) to generate insights. The RSS service is responsible for storing the request data and returning a trackable `request_id`, while the IGS service is responsible for generating the insights based on the requests and posting an update to a certain Kafka topic to inform active clients about the status of the request.

This asynchronous distributed processing is a key feature of the Transac AI project, as it allows each micro-service to scale independently and handle the requests in a fault-tolerant manner.

### Insights Generation

The WMS service is responsible for generating insights based on the requests received from the clients. The insights are generated using the IGS service, which is a separate micro-service in the TransacAI project. The IGS service uses the Generative AI and Large Language Models (LLMs) to generate the insights in natural language. The insights are then saved in a database accessible through the Insights Storage Service (ISS), and a message is posted to a Kafka topic to inform the clients about the status of the request with the `insights_id` and the `request_id` to let them know that the insights are ready to be fetched.

From WMS's perspective, workflow for a **manual request** for insights generation is as follows:

1. WMS receives a request for insights generation from the client (through the `GenerateInsights` RPC call).
2. WMS submits the request to the RSS service to store the request data and return a `request_id`.
3. WMS returns the `request_id` to the client so that they can track the status of the request through the RSS service.
4. WMS passes the `request_id` along with the request data to the IGS service to initiate generation of the insights.

From WMS's perspective, workflow for a **scheduled request** for insights generation is as follows:

1. WMS receives the request for scheduled insights generation from STM (Scheduler Task Manager) service.
2. WMS submits the request to the RSS service to store the request data and return a `request_id`.
3. WMS passes the `request_id` along with the request data to the IGS service to initiate generation of the insights.

From the perspective of the WMS service, the workflow for insights generation request ends once it has successfully submitted the request to the **IGS service**. The IGS service is responsible for generating the insights and updating the status of the request in the database.

## Local Development

This project is setup to run using Docker.

### Clone the repository

```bash
git clone
cd transacai-wms
```

### Setup Environment Variables

Create a `.env` file in the root of the project and add the following environment variables:

```bash
TRANSAC_AI_WMS_API_KEY=
```

You will need to send this API key in the `Authorization` header of the request to the WMS service.

### Build the Docker image

```bash
docker build -t transacai-wms .
```

### Run the Docker container

```bash
docker run --rm --env-file .env -p 8080:8080 transacai-wms
```

### Test the service

You can test the service by sending a POST request to the `/wms.v1.WMSService/GenerateInsights` endpoint. Here is an example request:

```bash
curl \
 --header "Content-Type: application/json" \
 --header "Authorization: Bearer 5657t6f0-e8sf67-gf83-fgcc9-70fggfgfge86" \
 --data '{"clientId":"client id","promptId":0,"recordsSourceId":"SUPABASE_PRIME_1","promptTemplatesSourceId":"AWS_D_DB_2","fromTime":"2023-12-29T06:39:22Z","toTime":"2023-12-29T23:49:22Z"}' \
 0.0.0.0:8080/wms.v1.WMSService/GenerateInsights
```

## Connect for Go

To learn more about writing RPC services in Go using Connect, check out the [Connect for Go](https://connectrpc.com/docs/go/getting-started) documentation.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Issues

If you encounter any issues or bugs while using this project, please report them by following these steps:

1. Check if the issue has already been reported by searching our [issue tracker](https://github.com/pranav-kural/transacai-workload-manager-service/issues).
2. If the issue hasn't been reported, create a new issue and provide a detailed description of the problem.
3. Include steps to reproduce the issue and any relevant error messages or screenshots.

[Open Issue](https://github.com/pranav-kural/transacai-workload-manager-service/issues/new)
