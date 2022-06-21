# prombackup

A webserver to trigger and backup Prometheus snapshots and expose backup status for prometheus
scrapping.

The two endpoints are:

- `/metrics`: for prometheus scrapping
- `/snap`: trigger a pipeline that will:
  1. start a Prometheus snapshot
  2. tarball the snapshot
  3. upload it to S3

