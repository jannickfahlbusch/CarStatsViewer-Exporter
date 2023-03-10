# CarStatsViewer-Exporter

# Running

To locally run the exporter you need to run the following:
```bash
go run cmd/carstatsviewer-exporter/main.go
```

By default, the exporter will listen on `:8080`.
You can customize this by specifying `-listen-address`.

Alternatively you can run the exporter using Docker via:
```bash
docker build -t carstatsviewer-exporter .
docker run -p 8080:8080 carstatsviewer-exporter
```

# Endpoints

The metrics are available under the `/metrics` endpoint.

Live data needs to be POSTed to `/live`. There is currently no security in place. Use this exporter only for experiments to play around with the HTTP live data PoC (See [Ixam97/CarStatsViewer#139](https://github.com/Ixam97/CarStatsViewer/pull/139)) of the [CarStatsViewer](https://github.com/Ixam97/CarStatsViewer).
