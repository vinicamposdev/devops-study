# observability
Created: 2022-12-11 13:52

## Concepts
### Basic Concepts
 - proactively collecting visualizing and applying intelligence to all right/meaningful metrics, events, logs, traces [1]
 - understand behavior of complex system
 - how well the understand of the system from the work it does
 - Control theory
	 - how engineers can infer the system's internal states
	 - from a knowledge of that system's external output
 - Cloud Computing
	 - how engineers can understand the system's current state
	 - from data it generates
### Observability vs Monitoring
 - Monitoring show errors
 - Monitoring is based on metrics and limits to monitor
 - Observability shows the reasons
### Pilars
 - Metrics
	 - Numbers, quantitative, precise
	 - helps guide infra and business
 - Logs
	 - helps understand complex errors
	 - compliance and auditory
 - Tracing
	 - sequence of events
	 - track down all steps
## Elastic Stack
### ELK Stack
 - Elasticsearch
	 - Search engine and analytics
	 - Apache Lucene
	 - 2010
	 - Fast
	 - Scalable
	 - Rest API
	 - Analytics of geospatial
	 - Application, website and enterprise search
	 - Logging and analytics
	 - Work distributed way from shards that has data resilience
	 - Scale with thousands of server with pentabytes of data
 - Logstash
	 - Data processor with pipelines, transform and send data
	 - Manipulate logs
	 - Work with pipelines
	 - Normalize and transform data
	 - Send/Receive data to multiple sources
	 - Plugins
	 - Nowadays is used less than before
 - Kibana
	 - Dashboard that vizualize and explore elasticsearch data
	 - See logs, series analysis, monitoring, operational inteligence
	 - Data aggregation and filters
	 - Maps and charts
### Beats 
 - 2015
 - Lightweight data shipper
 - "Replace" Logstash
 - Data Collector Agent
 - Logs, metrics, network data, audit data, uptime monitoring, custom beat


## References
1. [NewRelic](https://newrelic.com/blog/best-practices/what-is-observability)