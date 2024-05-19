
generate a server response like: `https://httpbin.org/get`

```shell
$ curl 'https://httpbin.org/get?a1=v1&a2=v2'
{
  "args": {
    "a1": "v1", 
    "a2": "v2"
  }, 
  "headers": {
    "Accept": "*/*", 
    "Host": "httpbin.org", 
    "User-Agent": "curl/7.79.1", 
    "X-Amzn-Trace-Id": "Root=1-6262c85b-5e6c222563487fce4ff0769d"
  }, 
  "origin": "127.0.0.1", 
  "url": "https://httpbin.org/get?a1=v1&a2=v2"
}

$ curl 'http://localhost:28001/get?a1=v1&a2=v2'
{"args":{"a1":"v1","a2":"v2"},"headers":{"Accept":"*/*","User-Agent":"curl/7.86.0"},"origin":"127.0.0.1","url":"http://localhost:28001/get?a1=v1&a2=v2"}
```