# Prometheus-manager

> prometheus unified rules and alarms management platform

<h1 style="display: flex; align-items: center; justify-content: center; gap: 10px; width: 100%; text-align: center;">
    <img alt="Prometheus" src="doc/img/logo.svg">
    <img alt="Prometheus" src="doc/img/prometheus-logo.svg">
</h1>

## Architecture overview

![Architecture overview](doc/img/Prometheus-manager.png)

## server 开发

```bash
make server
```

## agent 开发

```bash
make agent
```

## Docker 打包/运行

```bash
# build
docker build -t <your-docker-image-name> .

# run
docker run --rm -p 8000:8000 -p 9000:9000 -v </path/to/your/configs>:/data/conf <your-docker-image-name>
```