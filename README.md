# load_testing

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

## Log response header

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
