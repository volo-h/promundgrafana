
### start without docker compose
  #### set up and start Prometheus
  ```sh
    ./prometheus
  ```
  http://localhost:9090/targets
  
  http://127.0.0.1:9090/graph?g0.expr=client_balance_response_duration_histogram_bucket&g0.tab=1&g0.stacked=0&g0.show_exemplars=0&g0.range_input=1h

  
  #### set up and start Grafana
  ```sh
    ./grafana/bin/grafana-server -homepath /Users/lviv/workmy/prometheus/nats/pubsubnats/v2/grafana/grafana web
  ```
  Should launched oa port :3000
  
  admin\prom-operator
  
  #### run once
  ```sh
    for ((i=1;i<=100;i++)); do curl "http://localhost:8080/api/v1/balances"; done  
  ```
  
  http://localhost:8080/metrics

### start over docker compose
  ```sh
    make docker-up
    make generate-metrics

    http://localhost:9090/targets - should be green (with labels UP)
    http://localhost:3000/d/QwQN1JdGz/client - should see activity
  ```

### start over minikube
