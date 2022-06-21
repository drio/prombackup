# prombackup

A webserver to trigger and backup Prometheus snapshots and expose backup status for prometheus
scrapping.

The two endpoints are:

- `/metrics`: for prometheus scrapping
- `/snap`: trigger the snap pipeline.

When you hit the `/snap` endpoint via a GET request, the server will start a snapshot pipeline that
will:

  1. start a Prometheus snapshot
  2. tarball the snapshot
  3. upload it to S3

The `prom_backup_size` gauge it is zero most of the time. When the server finishes a successfull 
snapshot pipeline run, it sets the gauge to the size (in bytes) of the snapshot tarball. After a 
few seconds (configurable option) it sets the gauge back to zero. 

We can then use the gauge to determine when we did a successful snapshot backup.



