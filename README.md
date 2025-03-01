# fraud-detection
Detects fraudulent or unusual transactions

## Current state

- Data generated in Kotlin
- Sent to Kafka
- Picked up by Go, connects to Postgres and inserts records
- Metrics picked up by Prometheus, then displayed visually on Grafana


## Todo

- MiniKube orchestration 
- Prometheus metrics on latency of requests
- Think of a way to allow a service to perform clustering/other ML on this to spot anomaly
- Mass loading of data into system (ETL, if it's missing values etc)