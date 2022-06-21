# prombackup

This is a webserver to trigger Prometheus snapshots and expose the status via metrics.

The two endpoints are:

- `/metrics`: for prometheus scrapping
- `/snap`: trigger a pipeline that will:
  1. start a Prometheus snapshot
  2. tarball the snapshot
  3. upload it to S3

