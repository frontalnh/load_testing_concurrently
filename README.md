😃 Have a load testing easily and simply.

## Features

- No cache request
- Logging response
- Making an interval for each requests
- Requests concurrently!
- Check time elapsed
- Swarm!

## No logging

```terminal
➜ go run main.go
==============================================

🔍  Request 1
🔍  Request 2
🔍  Request 3
🔍  Request 4
🔍  Request 5
🔍  Request 6
🔍  Request 7
🔍  Request 8
🔍  Request 9
🔍  Request 10

⌛  [Swarm #1]   Elapsed: 123.955546ms

==============================================

🔍  Request 11
🔍  Request 12
🔍  Request 13
🔍  Request 14
🔍  Request 15
🔍  Request 16
🔍  Request 17
🔍  Request 18
🔍  Request 19
🔍  Request 20

⌛  [Swarm #2]   Elapsed: 95.801712ms

==============================================

🔍  Request 21
🔍  Request 22
🔍  Request 23
🔍  Request 24
🔍  Request 25
🔍  Request 26
🔍  Request 27
🔍  Request 28
🔍  Request 29
🔍  Request 30

⌛  [Swarm #3]   Elapsed: 77.341668ms

==============================================

🔍  Request 31
🔍  Request 32
🔍  Request 33
🔍  Request 34
🔍  Request 35
🔍  Request 36
🔍  Request 37
🔍  Request 38
🔍  Request 39
🔍  Request 40

⌛  [Swarm #4]   Elapsed: 133.939943ms
```

## Log responses

```terminal
➜ go run main.go
==============================================

🔍  Request 1
🔍  Request 2
🔍  Request 3
🔍  Request 4
🔍  Request 5
🔍  Request 6
🔍  Request 7
🔍  Request 8
🔍  Request 9
🔍  Request 10

✅   [Header] >>  &{200 OK 200 HTTP/1.1 1 1 map[Cache-Control:[private, no-cache, no-store, must-revalidate] Content-Type:[application/json; charset=utf-8] Content-Length:[346] Vary:[Accept-Encoding] Date:[Tue, 13 Nov 2018 04:16:23 GMT] X-Powered-By:[Express] Pragma:[no-cache] Etag:[W/"15a-4gnsxqOsMCViv3zvIAVXxx1F88Y"] Connection:[keep-alive] Expires:[-1]] 0xc000276280 346 [] false false map[] 0xc0001a4000 <nil>}

✅   [Header] >>  &{200 OK 200 HTTP/1.1 1 1 map[Content-Type:[application/json; charset=utf-8] Etag:[W/"10f-Kf3S56agoipYSe3BaJjRk6FnJNA"] Vary:[Accept-Encoding] Date:[Tue, 13 Nov 2018 04:16:23 GMT] X-Powered-By:[Express] Cache-Control:[private, no-cache, no-store, must-revalidate] Expires:[-1] Pragma:[no-cache] Content-Length:[271] Connection:[keep-alive]] 0xc000096380 271 [] false false map[] 0xc000212000 <nil>}

✅   [Header] >>  &{200 OK 200 HTTP/1.1 1 1 map[Content-Length:[271] Date:[Tue, 13 Nov 2018 04:16:23 GMT] Connection:[keep-alive] X-Powered-By:[Express] Expires:[-1] Pragma:[no-cache] Content-Type:[application/json; charset=utf-8] Etag:[W/"10f-a44Le1ktcCivtzArytQkmZNZ6jI"] Vary:[Accept-Encoding] Cache-Control:[private, no-cache, no-store, must-revalidate]] 0xc000096400 271 [] false false map[] 0xc000140200 <nil>}

✅   [Header] >>  &{200 OK 200 HTTP/1.1 1 1 map[Cache-Control:[private, no-cache, no-store, must-revalidate] Expires:[-1] Etag:[W/"10f-a44Le1ktcCivtzArytQkmZNZ6jI"] Vary:[Accept-Encoding] X-Powered-By:[Express] Pragma:[no-cache] Content-Type:[application/json; charset=utf-8] Content-Length:[271] Date:[Tue, 13 Nov 2018 04:16:23 GMT] Connection:[keep-alive]] 0xc000276300 271 [] false false map[] 0xc000140000 <nil>}

✅   [Header] >>  &{200 OK 200 HTTP/1.1 1 1 map[Pragma:[no-cache] Content-Type:[application/json; charset=utf-8] Content-Length:[312] Date:[Tue, 13 Nov 2018 04:16:23 GMT] Connection:[keep-alive] X-Powered-By:[Express] Expires:[-1] Etag:[W/"138-z3XtxOE8FvVaY2YdFkXU61B3lhI"] Vary:[Accept-Encoding] Cache-Control:[private, no-cache, no-store, must-revalidate]] 0xc000096480 312 [] false false map[] 0xc00015a000 <nil>}

✅   [Header] >>  &{200 OK 200 HTTP/1.1 1 1 map[X-Powered-By:[Express] Expires:[-1] Pragma:[no-cache] Content-Length:[312] Date:[Tue, 13 Nov 2018 04:16:23 GMT]Connection:[keep-alive] Cache-Control:[private, no-cache, no-store, must-revalidate] Content-Type:[application/json; charset=utf-8] Etag:[W/"138-FC0TaKY5QzSWuYi+6yFHQAvjdyk"] Vary:[Accept-Encoding]] 0xc000276380 312 [] false false map[] 0xc000122000 <nil>}

✅   [Header] >>  &{200 OK 200 HTTP/1.1 1 1 map[Cache-Control:[private, no-cache, no-store, must-revalidate] Expires:[-1] Pragma:[no-cache] Content-Type:[application/json; charset=utf-8] Etag:[W/"15a-XqD+yUuT/8jzFIQCTmst7bbWbSM"] Date:[Tue, 13 Nov 2018 04:16:23 GMT] X-Powered-By:[Express] Vary:[Accept-Encoding] Connection:[keep-alive] Content-Length:[346]] 0xc000276400 346 [] false false map[] 0xc000140100 <nil>}

✅   [Header] >>  &{200 OK 200 HTTP/1.1 1 1 map[Cache-Control:[private, no-cache, no-store, must-revalidate] Expires:[-1] Content-Type:[application/json; charset=utf-8] Etag:[W/"15a-mIV7882sNn4X2RdD6cCeZ7Gy9mo"] X-Powered-By:[Express] Pragma:[no-cache] Content-Length:[346] Vary:[Accept-Encoding] Date:[Tue, 13 Nov 2018 04:16:23 GMT] Connection:[keep-alive]] 0xc000276480 346 [] false false map[] 0xc000226000 <nil>}

✅   [Header] >>  &{200 OK 200 HTTP/1.1 1 1 map[Etag:[W/"15a-RiiWQJyObfqtzvACbZpvODlF/CA"] Vary:[Accept-Encoding] Date:[Tue, 13 Nov 2018 04:16:23 GMT] Expires:[-1] Pragma:[no-cache] Content-Length:[346] Connection:[keep-alive] X-Powered-By:[Express] Cache-Control:[private, no-cache, no-store, must-revalidate] Content-Type:[application/json; charset=utf-8]] 0xc000022580 346 [] false false map[] 0xc000140300 <nil>}

✅   [Header] >>  &{200 OK 200 HTTP/1.1 1 1 map[X-Powered-By:[Express] Content-Type:[application/json; charset=utf-8] Etag:[W/"15a-XQlF08aPJQP1Efv390Kp5TJUNzY"] Vary:[Accept-Encoding] Date:[Tue, 13 Nov 2018 04:16:23 GMT] Cache-Control:[private, no-cache, no-store, must-revalidate] Expires:[-1] Pragma:[no-cache] Content-Length:[346] Connection:[keep-alive]] 0xc000022600 346 [] false false map[] 0xc000160000 <nil>}


⌛  [Swarm #1]   Elapsed: 41.237334ms

==============================================

🔍  Request 11
🔍  Request 12
🔍  Request 13
🔍  Request 14
🔍  Request 15
🔍  Request 16
🔍  Request 17
🔍  Request 18
🔍  Request 19
🔍  Request 20

✅   [Header] >>  &{200 OK 200 HTTP/1.1 1 1 map[Cache-Control:[private, no-cache, no-store, must-revalidate] Content-Length:[316] Date:[Tue, 13 Nov 2018 04:16:24 GMT] Connection:[keep-alive] X-Powered-By:[Express] Expires:[-1] Pragma:[no-cache] Content-Type:[application/json; charset=utf-8] Etag:[W/"13c-rcySn5L89rTr5FTtcPFC3ww/aXQ"] Vary:[Accept-Encoding]] 0xc00026e6c0 316 [] false false map[] 0xc0001a4200 <nil>}

✅   [Header] >>  &{200 OK 200 HTTP/1.1 1 1 map[Pragma:[no-cache] Content-Type:[application/json; charset=utf-8] Vary:[Accept-Encoding] Connection:[keep-alive]X-Powered-By:[Express] Cache-Control:[private, no-cache, no-store, must-revalidate] Expires:[-1] Content-Length:[346] Etag:[W/"15a-XQlF08aPJQP1Efv390Kp5TJUNzY"] Date:[Tue, 13 Nov 2018 04:16:24 GMT]] 0xc000276600 346 [] false false map[] 0xc000140500 <nil>}

✅   [Header] >>  &{200 OK 200 HTTP/1.1 1 1 map[Expires:[-1] Pragma:[no-cache] Content-Type:[application/json; charset=utf-8] Date:[Tue, 13 Nov 2018 04:16:24 GMT] Connection:[keep-alive] X-Powered-By:[Express] Cache-Control:[private, no-cache, no-store, must-revalidate] Content-Length:[316] Etag:[W/"13c-rcySn5L89rTr5FTtcPFC3ww/aXQ"] Vary:[Accept-Encoding]] 0xc00026e740 316 [] false false map[] 0xc000140400 <nil>}

✅   [Header] >>  &{200 OK 200 HTTP/1.1 1 1 map[Vary:[Accept-Encoding] Connection:[keep-alive] Pragma:[no-cache] Cache-Control:[private, no-cache, no-store, must-revalidate] Expires:[-1] Content-Type:[application/json; charset=utf-8] Content-Length:[271] Etag:[W/"10f-QthcjG1AnZPGik3Qcw9iDKIMv0k"] Date:[Tue, 13 Nov 2018 04:16:24 GMT] X-Powered-By:[Express]] 0xc000276680 271 [] false false map[] 0xc00015a100 <nil>}

✅   [Header] >>  &{200 OK 200 HTTP/1.1 1 1 map[Content-Length:[312] Vary:[Accept-Encoding] Connection:[keep-alive] X-Powered-By:[Express] Cache-Control:[private, no-cache, no-store, must-revalidate] Pragma:[no-cache] Date:[Tue, 13 Nov 2018 04:16:24 GMT] Expires:[-1] Content-Type:[application/json; charset=utf-8] Etag:[W/"138-z3XtxOE8FvVaY2YdFkXU61B3lhI"]] 0xc000022780 312 [] false false map[] 0xc000122200 <nil>}

✅   [Header] >>  &{200 OK 200 HTTP/1.1 1 1 map[Cache-Control:[private, no-cache, no-store, must-revalidate] Pragma:[no-cache] Date:[Tue, 13 Nov 2018 04:16:24 GMT] Connection:[keep-alive] X-Powered-By:[Express] Expires:[-1] Content-Type:[application/json; charset=utf-8] Content-Length:[312] Etag:[W/"138-opGPf8qKAJFX9fIKaNoiEgJa1ws"] Vary:[Accept-Encoding]] 0xc000276700 312 [] false false map[] 0xc0001a4100 <nil>}

✅   [Header] >>  &{200 OK 200 HTTP/1.1 1 1 map[Content-Length:[271] Etag:[W/"10f-Kf3S56agoipYSe3BaJjRk6FnJNA"] Vary:[Accept-Encoding] Expires:[-1] Pragma:[no-cache] Content-Type:[application/json; charset=utf-8] Connection:[keep-alive] X-Powered-By:[Express] Cache-Control:[private, no-cache, no-store, must-revalidate] Date:[Tue, 13 Nov 2018 04:16:24 GMT]] 0xc000276780 271 [] false false map[] 0xc000226100 <nil>}

✅   [Header] >>  &{200 OK 200 HTTP/1.1 1 1 map[Vary:[Accept-Encoding] X-Powered-By:[Express] Expires:[-1] Pragma:[no-cache] Content-Length:[271] Connection:[keep-alive] Cache-Control:[private, no-cache, no-store, must-revalidate] Content-Type:[application/json; charset=utf-8] Etag:[W/"10f-QthcjG1AnZPGik3Qcw9iDKIMv0k"] Date:[Tue, 13 Nov 2018 04:16:24 GMT]] 0xc00026e7c0 271 [] false false map[] 0xc000212100 <nil>}

✅   [Header] >>  &{200 OK 200 HTTP/1.1 1 1 map[Content-Length:[346] Vary:[Accept-Encoding] Connection:[keep-alive] X-Powered-By:[Express] Cache-Control:[private, no-cache, no-store, must-revalidate] Content-Type:[application/json; charset=utf-8] Etag:[W/"15a-XQlF08aPJQP1Efv390Kp5TJUNzY"] Date:[Tue, 13 Nov 2018 04:16:24 GMT] Expires:[-1] Pragma:[no-cache]] 0xc000276800 346 [] false false map[] 0xc000122300 <nil>}

✅   [Header] >>  &{200 OK 200 HTTP/1.1 1 1 map[X-Powered-By:[Express] Expires:[-1] Pragma:[no-cache] Content-Length:[352] Date:[Tue, 13 Nov 2018 04:16:24 GMT]Cache-Control:[private, no-cache, no-store, must-revalidate] Content-Type:[application/json; charset=utf-8] Etag:[W/"160-rvLeI6h+W5YVpa3RUCsFfu4riUg"] Vary:[Accept-Encoding] Connection:[keep-alive]] 0xc000096640 352 [] false false map[] 0xc00032a000 <nil>}


⌛  [Swarm #2]   Elapsed: 43.624333ms
```

```terminal
➜ go run main.go
==============================================

🔍  Request 1
🔍  Request 2
🔍  Request 3
🔍  Request 4
🔍  Request 5
🔍  Request 6
🔍  Request 7
🔍  Request 8
🔍  Request 9
🔍  Request 10

✅   [BODY]  {"response":0.4430138339981695}

✅   [BODY]  {"response":0.793714230446541}

✅   [BODY]  {"response":0.706769358253222}

✅   [BODY]  {"response":0.9522459782391726}

✅   [BODY]  {"response":0.3927239683091339}

✅   [BODY]  {"response":0.4755491071351652}

✅   [BODY]  {"response":0.0897603037050636}

✅   [BODY]  {"response":0.6560212119090747}

✅   [BODY]  {"response":0.502549026726649}

✅   [BODY]  {"response":0.9302333508703451}


⌛  [Swarm #1] 	Elapsed: 36.322127ms

==============================================

🔍  Request 11
🔍  Request 12
🔍  Request 13
🔍  Request 14
🔍  Request 15
🔍  Request 16
🔍  Request 17
🔍  Request 18
🔍  Request 19
🔍  Request 20

✅   [BODY]  {"response":0.8359390967226725}

✅   [BODY]  {"response":0.8725942409517111}

✅   [BODY]  {"response":0.7187650945903132}

✅   [BODY]  {"response":0.9267171576758975}

✅   [BODY]  {"response":0.053200464941225434}

✅   [BODY]  {"response":0.8883227394358442}

✅   [BODY]  {"response":0.7421644154296805}

✅   [BODY]  {"response":0.3177492127404833}

✅   [BODY]  {"response":0.45470919030627943}

✅   [BODY]  {"response":0.2820161846848719}


⌛  [Swarm #2] 	Elapsed: 26.767254ms
```
