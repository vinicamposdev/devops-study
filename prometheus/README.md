# observability
Created: 2022-12-11 13:52

## Concepts
### Basic Concepts
 - proactively collecting visualizing and applying intelligence to all right/meaningful metrics, events, logs, traces [2]
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
![[Pasted image 20221211142430.png]]
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
### Installation
 - use examples [2] to get docker compose containers
### HeartBeat
 - get healthy of applications
### APM
 - can trace logs from frontend to backend using trace_id and span_id
	 - rum (Real User Monitoring)
	 - LRU cache is used to keep rate limite per IP
 - [Agents](https://www.elastic.co/guide/en/apm/agent/index.html): pluggins and libs to langs
### Filebeat
 - lightweight way to forward and centralize logs and files
 - create personalized dashboards
	 - knows where are the logs comming
	 - aggregate logs
	 - get size
## Prometheus
### Basic Concepts
 - monitoring solution
 - from metrics to insights
 - toolkit metrics and alerts
 - open-source
 - SoundClound
 - Cloud Native Computing Foundation
 - On Premise, independent, agnostic
 - Dimensional Datas
	 - not only raw data
	 - histograms
	 - aggregate data
 - Powerful querying
 - Easy to visualize data with Grafana
 - Efficiency storage
 - Easy
 - Smart alert
	 - chain alerts
	 - level of priorities
 - Exporters: integrations, libs
### Pull Logs
 - push: active -> elastic
 - pull: passive -> prometheus
 - pulling logs via http
 - need an endpoint /metrics
### Exporters
 - get metrics from
	 - MySQL
	 - Nginx / Apache
	 - Servers, etc
	 - Custom exporter
### Architecture
 - get
 - store
	 - TSDB - Time Series Database
		 - fast to query
		 - solve problems with timestamp
	 - Retrieval
		 - orchestrate requests
	 - HTTP server
		 - monitor himself and others
 - make available
### Metrics
 - Counter
	 - incremental value
	 - Prometheus can handle fault and start from the last number in case of eventual reset
	 - Eg.:
		 - Site visites
		 - sales
		 - errors
 - Gauge
	 - has delta on time
	 - can increase, reduce, stabilyze
	 - Eg.:
		 - users active
		 - servers active
 - Histogram
	 - Distribute data with frequency
	 - Measure based on samples
	 - Aggregate data
	 - Eg.: sales by age
 - Summary
	 - Aggregate values based on calculated values on prometheus server
	 - calculate with more accurate values
	 - Eg.: request duration data
	 - most common use histogram instead
### PromQL
 - Prometheus Query Language
 - Eg.:
	 - http_requests_total
	 - rate(http_requests_total[5m])
	 - http_requests_total{status!~"4.."}
## References
1. [NewRelic](https://newrelic.com/blog/best-practices/what-is-observability)
2. [Elastic Examples](https://github.com/elastic/examples)