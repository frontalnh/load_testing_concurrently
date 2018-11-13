😃 Have a load testing easily and simply.

## Features

- No cache request
- Logging response
- Making an interval for each requests
- Requests concurrently!
- Check time elapsed
- Swarm

## No logging

```sh
➜ go run main.go
🔍   Request 0
🔍   Request 1
🔍   Request 2
🔍   Request 3
🔍   Request 4
🔍   Request 5
🔍   Request 6
🔍   Request 7
🔍   Request 8
🔍   Request 9
🔍   Request 10
🔍   Request 11
🔍   Request 12
🔍   Request 13
🔍   Request 14
🔍   Request 15
🔍   Request 16
🔍   Request 17
🔍   Request 18
🔍   Request 19
🔍   Request 20
🔍   Request 21
🔍   Request 22
🔍   Request 23
🔍   Request 24
🔍   Request 25
🔍   Request 26
🔍   Request 27
🔍   Request 28
🔍   Request 29
🔍   Request 30
🔍   Request 31
🔍   Request 32
🔍   Request 33
🔍   Request 34


⌛  >> 137.540897ms
```

## Log responses

```sh
➜ go run main.go
🔍   Request 0
🔍   Request 1
🔍   Request 2
🔍   Request 3
🔍   Request 4
🔍   Request 5
🔍   Request 6
🔍   Request 7
🔍   Request 8
🔍   Request 9
🔍   Request 10


✅   [Header] >>  &{200 OK 200 HTTP/1.1 1 1 map[Date:[Tue, 13 Nov 2018 02:31:39 GMT] X-Powered-By:[Express] Expires:[-1] Etag:[W/"1f-iPQ9/r7wjWGMEb/LV6fiufxwLS8"] Content-Length:[31] Vary:[Accept-Encoding] Connection:[keep-alive] Cache-Control:[private, no-cache, no-store, must-revalidate] Pragma:[no-cache] Content-Type:[application/json; charset=utf-8]] 0xc0001f0340 31 [] false false map[] 0xc00023c000 <nil>}

✅   [Header] >>  &{200 OK 200 HTTP/1.1 1 1 map[Pragma:[no-cache] Etag:[W/"1f-3Ldndn1jG8THIg1BYeyHobt0e9o"] Vary:[Accept-Encoding] Connection:[keep-alive] Cache-Control:[private, no-cache, no-store, must-revalidate] Expires:[-1] Content-Length:[31] Date:[Tue, 13 Nov 2018 02:31:39 GMT] X-Powered-By:[Express] Content-Type:[application/json; charset=utf-8]] 0xc0001f03c0 31 [] false false map[] 0xc000142100 <nil>}

✅   [Header] >>  &{200 OK 200 HTTP/1.1 1 1 map[Content-Length:[32] Date:[Tue, 13 Nov 2018 02:31:39 GMT] X-Powered-By:[Express] Expires:[-1] Pragma:[no-cache] Content-Type:[application/json; charset=utf-8] Cache-Control:[private, no-cache, no-store, must-revalidate] Etag:[W/"20-4IY1A6aDpP8QH+o0XH6Lr3zL40Q"] Vary:[Accept-Encoding] Connection:[keep-alive]] 0xc000022600 32 [] false false map[] 0xc00013e000 <nil>}

✅   [Header] >>  &{200 OK 200 HTTP/1.1 1 1 map[Vary:[Accept-Encoding] Connection:[keep-alive] X-Powered-By:[Express] Cache-Control:[private, no-cache, no-store, must-revalidate] Expires:[-1] Pragma:[no-cache] Etag:[W/"1f-f61wPjL4w9GmYdhlGgPi5MV9q0Q"] Content-Type:[application/json; charset=utf-8] Content-Length:[31] Date:[Tue, 13 Nov 2018 02:31:39 GMT]] 0xc0001f0440 31 [] false false map[] 0xc000140000 <nil>}

✅   [Header] >>  &{200 OK 200 HTTP/1.1 1 1 map[Expires:[-1] Pragma:[no-cache] Etag:[W/"1f-4r1buth43CsePII/3Z/mIXUZ07E"] Vary:[Accept-Encoding] X-Powered-By:[Express] Cache-Control:[private, no-cache, no-store, must-revalidate] Date:[Tue, 13 Nov 2018 02:31:39 GMT] Connection:[keep-alive] Content-Type:[application/json; charset=utf-8] Content-Length:[31]] 0xc000022680 31 [] false false map[] 0xc000142000 <nil>}

✅   [Header] >>  &{200 OK 200 HTTP/1.1 1 1 map[X-Powered-By:[Express] Cache-Control:[private, no-cache, no-store, must-revalidate] Content-Length:[32] Date:[Tue, 13 Nov 2018 02:31:39 GMT] Connection:[keep-alive] Expires:[-1] Pragma:[no-cache] Content-Type:[application/json; charset=utf-8] Etag:[W/"20-1KD14AML+y1UMwjXZl0o11M110U"] Vary:[Accept-Encoding]] 0xc000286380 32 [] false false map[] 0xc000162000 <nil>}

✅   [Header] >>  &{200 OK 200 HTTP/1.1 1 1 map[Cache-Control:[private, no-cache, no-store, must-revalidate] Pragma:[no-cache] Content-Type:[application/json; charset=utf-8] Content-Length:[31] Etag:[W/"1f-SC1sETev+6vXFNW9orGpLz86O5s"] Vary:[Accept-Encoding] Connection:[keep-alive] X-Powered-By:[Express] Expires:[-1] Date:[Tue, 13 Nov 2018 02:31:39 GMT]] 0xc000286400 31 [] false false map[] 0xc0001a4000 <nil>}

✅   [Header] >>  &{200 OK 200 HTTP/1.1 1 1 map[Etag:[W/"20-rVi9JhrQoNgoMg47p/EQS1K9c34"] X-Powered-By:[Express] Pragma:[no-cache] Content-Type:[application/json; charset=utf-8] Vary:[Accept-Encoding] Date:[Tue, 13 Nov 2018 02:31:39 GMT] Connection:[keep-alive] Cache-Control:[private, no-cache, no-store, must-revalidate] Expires:[-1] Content-Length:[32]] 0xc000022700 32 [] false false map[] 0xc000140100 <nil>}

✅   [Header] >>  &{200 OK 200 HTTP/1.1 1 1 map[X-Powered-By:[Express] Expires:[-1] Pragma:[no-cache] Date:[Tue, 13 Nov 2018 02:31:39 GMT] Connection:[keep-alive] Cache-Control:[private, no-cache, no-store, must-revalidate] Content-Type:[application/json; charset=utf-8] Content-Length:[31] Etag:[W/"1f-tiZK6uOt8t5L4QH0ukOrGiAks5E"] Vary:[Accept-Encoding]] 0xc0001f04c0 31 [] false false map[] 0xc000142200 <nil>}

✅   [Header] >>  &{200 OK 200 HTTP/1.1 1 1 map[Expires:[-1] Pragma:[no-cache] Vary:[Accept-Encoding] Date:[Tue, 13 Nov 2018 02:31:39 GMT] Connection:[keep-alive] X-Powered-By:[Express] Content-Type:[application/json; charset=utf-8] Content-Length:[31] Etag:[W/"1f-rC2c4FYgOjWdE93aLc6N4uomCQY"] Cache-Control:[private, no-cache, no-store, must-revalidate]] 0xc000286480 31 [] false false map[] 0xc0001d4000 <nil>}


⌛  >> 32.990381ms
```

```sh
➜ go run main.go
🔍   Request 0
🔍   Request 1
🔍   Request 2
🔍   Request 3
🔍   Request 4
🔍   Request 5
🔍   Request 6
🔍   Request 7
🔍   Request 8
🔍   Request 9
🔍   Request 10

✅   [BODY]  {"response":0.9728589881890994}

✅   [BODY]  {"response":0.2856301562614749}

✅   [BODY]  {"response":0.9409020976153937}

✅   [BODY]  {"response":0.29267852054349963}

✅   [BODY]  {"response":0.07673742899867952}

✅   [BODY]  {"response":0.9964633317966514}

✅   [BODY]  {"response":0.9275394677593376}

✅   [BODY]  {"response":0.6831850705111044}

✅   [BODY]  {"response":0.6979834369963405}

✅   [BODY]  {"response":0.6228743354707997}

✅   [BODY]  {"response":0.8211995797102749}

⌛  >> 124.162719ms
```
