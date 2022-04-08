<h1 align="center"><img src="./Comet.png" width="100" /></h1>
<h1 align="center">Comet</h1>
<p align="center">A simple and fast web-powered fileserver written in Go.</p>

## Deployment
We recommend using Docker for deployment.
```
$ docker run -d -p 8080:8080 -v /path/to/comet/files:/files gcr.io/astralservices/comet:latest
```

### Traefik
We recommend using Traefik for reverse-proxying.

You can also use `rsync` to sync your files to multiple nodes.
```
$ rsync -avz /path/to/comet/files/ node1:/path/to/comet/files/
```