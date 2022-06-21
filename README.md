# progo-backup

This is a tool to help you backup Prometheus data/snapshots. If you run into a failure in your 
Prometheus server, you may end up losing your time series data. This is a go tool you can run 
to:

  - Trigger prometheus [snapshots](https://prometheus.io/docs/prometheus/latest/querying/api/#snapshot) 
    snapshots.
  - Dump them into S3/Backblaze
  - Expose prometheus metrics to monitor this process


