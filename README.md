# Minions

Minions are simple cache storages that are being controlled by a
router for caching MinIO cluster's objects. By using Minions you can
speed up objects fetching process in MinIO cluster.

## How do they work?

Each Minion caches  a requested object in its local storage. Router manages each request's response
by executing a query on Minions. If it hits cache miss, it will get the
requested file from MinIO cluster and clones it into at least one of the
active cache instances.
