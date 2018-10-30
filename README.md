### This demo uses open-source charts that have been slightly tweaked to work with each other.
### This is a great collection of tools to use if your are just getting started with EFK

To install run the following:

1. Install elasticsearch-operator (original chart https://github.com/upmc-enterprises/elasticsearch-operator/tree/master/charts/elasticsearch-operator):
`helm install --name elasticsearch-operator elasticsearch-operator --set rbac.enabled=True`

2. Install elasticsearch chart (original chart https://github.com/upmc-enterprises/elasticsearch-operator/tree/master/charts/elasticsearch)
`helm install --name elasticsearch elasticsearch --set rbac.enabled=True`

3. Install kibana chart (original chart https://github.com/helm/charts/tree/master/stable/kibana). This needs to be installed once Elasticsearch cluster is up as it will Kibana will fail if it cannot connect to Elasticsearch)
`helm install --name kibana kibana`

4. Install fluentd chart (original chart https://github.com/helm/charts/tree/master/stable/fluentd-elasticsearch)
`helm install --name fluentd fluentd-elasticsearch`

5. To create log generation run:
`kubectl create -f behaved-logs.yaml` - this wraps the `log-generator` app

6. To generate many logs exec in the new pod created by `behaved-logs.yaml` and run the following commands:
`curl localhost:4003/behave` - if this returns `true` running `curl localhost:4003/spam` - will start creating logs of a constant data mapping type. If `curl localhost:4003/behave` returns `false` running `curl localhost:4003/spam` - will generate a lot of mixed logs with fields of varying data types.

7. To stop log generation rerun `curl localhost:4003/spam` - if it returns false it is no longer generating logs.

Useful resources:
* https://www.elastic.co/blog/a-heap-of-trouble
* https://github.com/fluent/fluent-bit - lightweight version of fluentd
* https://github.com/helm/charts/tree/master/stable/fluent-bit - fluent-bit chart
* https://github.com/Yelp/elastalert - ElastAlert is a little python based opensource tool that pushes alerts from EFK to various alert receivers.
