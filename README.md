# file-upload

[![Go Report Card](https://goreportcard.com/badge/github.com/dadez/file-upload)](https://goreportcard.com/report/github.com/dadez/file-upload) ![version](https://img.shields.io/badge/version-v.0.0.20-blue.svg?cacheSeconds=2592000)

## Goal

The purpose of this program is to be able to upload certificate files in a PEM format for later use as sidecar container together with https://github.com/dadez/cert-exporter  
Only upload of valid PEM files will be accepted

## Usage

### configuration

| environment variable | default value | description |
|:--- | :---:| ---|
| `STATIC_FILES_PATH` | `./web` | where the index.html is stored |
| `AUTH_FILES_PATH` | `./test` | where the username and password files are stored |
| `UPLOADS_DIRECTORY_PATH` | `uploads` |  where the uploaded files should be stored |
| `LOG_LEVEL` | `info` | loglevel for applications logs, possible values are: `trace`, `debug`, `info`, `warn`, `error`, `fatal`, `panic` see https://github.com/sirupsen/logrus#level-logging



## build & run

### go run

```bash
cd file-upload # move directory to this project
# run application
go run cmd/app/main.go
```

### run from compiled binary

get a release on the [release page](https://github.com/dadez/file-upload/releases)

```bash
# extract binary
tar xzf file-upload_<version>_<arch>.tar.gz

# run binary
./file-upload
```


### build docker image

```bash
cd file-upload # move directory to this project
docker build -t file-upload:latest -f build/package/Dockerfile .
```

### run docker image

```bash
docker run -p 4500:4500 file-upload:latest
```

## upload certificates

### a single certificate with curl

```bash
cd file-upload # move directory to this project
curl -X POST -u user:secret -F file=@test/github.crt.pem http://localhost:4500/upload
```

### multiple certificates with curl

just repeat the `-F file=@/path/to/file.pem` part in your command


### from web UI

just connect to http://localhost:4500, login and choose the file(s) to upload, press on submit.


# todo

- [ ] run behind https
- [ ] do not crash container on cert validation (ie: cert expired `time="2020-12-23T17:01:26Z" level=fatal msg="failed to verify certificatex509: certificate has expired or is not yet valid: current time 2020-12-23T17:01:26Z is after 2020-12-11T13:04:36Z" `)
- [ ] do not crash container when file is invalid (ie: `2020/12/23 16:52:46 http: superfluous response.WriteHeader call from github.com/dadez/file-upload/pkg/upload.UploadHandler (upload.go:57)`)
- [ ] prettier web part

---

> *this is the original readme from the forked repo*

# How to process file uploads in Go

This repo contains the complete code used in [this Freshman
tutorial](https://freshman.tech/file-upload-golang/). Clone this repo to your
computer and run `go run main.go` to start the server on PORT 4500.

