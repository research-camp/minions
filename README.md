# Minions

Minions are simple cache storages that are being controlled by a
router for caching MinIO cluster's objects. By using Minions you can
speed up objects fetching process in MinIO cluster.

## How do they work?

Each Minion caches  a requested object in its local storage. Router manages each request's response
by executing a query on Minions. If it hits cache miss, it will get the
requested file from MinIO cluster and clones it into at least one of the
active cache instances.

### labeling

In order to increase our cache hit ratio, we use a labeling algorithm to map the object
name to a constant cache storage. We take the object name, and we count the number of letters.
After that we get it's mod by the number of cache storages. The result is the cache index
that the object should be stored in.

For example, if we upload a file named ```"file.txt"```, and we have five cache storages available,
our labeling algorithm goes like this:

```python
a = len("file.txt") # 8
index = a % 5 # 3
return f"cache-{index}"
```

Our labeling algorithm is stateless. Therefore, we don't need to use any storage space
in our router.
