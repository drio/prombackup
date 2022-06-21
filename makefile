start_prometheus:
	prometheus --web.enable-admin-api

snap:
	@curl -s -X POST localhost:9090/api/v1/admin/tsdb/snapshot
