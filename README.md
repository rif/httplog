## HttpLog

Prints detailed info about the received requests.

### Install
go get -u github.com/rif/httplog

### Usage

```
$ httplog 8000 // default port is 8080

$ curl --data "test=1"  http://localhost:8000/mama
```

### Example output

```
{
 "Method": "POST",
 "URL": {
  "Scheme": "",
  "Opaque": "",
  "User": null,
  "Host": "",
  "Path": "/mama",
  "RawQuery": "",
  "Fragment": ""
 },
 "Proto": "HTTP/1.1",
 "ProtoMajor": 1,
 "ProtoMinor": 1,
 "Header": {
  "Accept": [
   "*/*"
  ],
  "Content-Length": [
   "6"
  ],
  "Content-Type": [
   "application/x-www-form-urlencoded"
  ],
  "User-Agent": [
   "curl/7.34.0"
  ]
 },
 "Body": {
  "Reader": {
   "R": {},
   "N": 0
  }
 },
 "ContentLength": 6,
 "TransferEncoding": null,
 "Close": false,
 "Host": "localhost:8080",
 "Form": {
  "test": [
   "1"
  ]
 },
 "PostForm": {
  "test": [
   "1"
  ]
 },
 "MultipartForm": null,
 "Trailer": null,
 "RemoteAddr": "[::1]:59065",
 "RequestURI": "/mama",
 "TLS": null
}
```
