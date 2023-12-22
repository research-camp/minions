# Minions

A distributed caching system for __MinIO__ cluster. Build with __Golang__. By using this system, we can
speed our object storage cluster operations up to 20% faster as the average.

## collaborators

- Amirhossein Najafizadeh (Amirkabir University of Technology)
- Dr. Niloofar Charmchi (Université de Rennes I)

## how to run?

You have to setup a __MinIO__ cluster with docker.

```
docker run \
  -p 9000:9000 \
  -p 9001:9001 \
  -e "MINIO_ROOT_USER=AKIAIOSFODNN7EXAMPLE" \
  -e "MINIO_ROOT_PASSWORD=wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY" \
  quay.io/minio/minio server /data --console-address ":9001"
```

After that you need to install two components, the minions and the router.

### router

```
go run main.go router
```

### minion

```
go run main.go minion
```
