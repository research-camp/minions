# Minions

A distributed caching system for __MinIO__ cluster. Build with __Golang__. This system uses a main router
which monitors cache systems and balances the input load between them. It also uses an __ETCD__ storage in order
to keep the hit ratio as high as possible.
By using this system, we can speed our object storage cluster operations up to 20% faster as the average.

## Collaborators

- Amirhossein Najafizadeh (Amirkabir University of Technology)
- Dr. Niloofar Charmchi (Université de Rennes I)

## How it works?

We group our clients based on regions. After that we create a Minion for our regions. Some regions may have more than
one Minion based on the traffic that we receive from there. For every input request, we pass them to the router and it
will proxy the request to one of the Minions. Through this proxy, it gets a factor number from each Minion and it balances
the load between them by selecting the right Minion each time.

Each Minion updates its factor in an interval and sends it to the router. The router keeps this metadata in an __ETCD__ storage
to use them for future requests.

![](.github/asserts/diagram.svg)

## How to run?

First, you need to setup a __MinIO__ cluster. You can do this by docker:

```
docker run \
  -p 9000:9000 \
  -p 9001:9001 \
  -e "MINIO_ROOT_USER=AKIAIOSFODNN7EXAMPLE" \
  -e "MINIO_ROOT_PASSWORD=wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY" \
  quay.io/minio/minio server /data --console-address ":9001"
```

After that you need to install two components, the minions and the router.

### Router

Router is the brain behind all the system. It monitors Minions, controls the input load by
defining a load factor for each Minion, and tracks the cluster objects by storing them in an ETCD database.

You can start the router by using the following command:

```
go run main.go router
```

### Minion

Each Minion is a file manager which uses an interface in order to make calls to MinIO cluster.
For each Minion, a load factor is defined by $count(files) * sum(files.size)$. This factor allows the router to
handle the input traffic and keep the hitting ratio high.

You can start a single Minion by using the following command:

```
go run main.go minion
```
