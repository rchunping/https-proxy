# https-proxy
A simple https/http proxy, with Basic Authentication


## Build
```shell script
go build
```

## Usage


### HTTP Proxy
```shell script
./https-proxy -listen=":8080"
```
#### Test
```shell script
curl --proxy http://localhost:8080 https://www.baidu.com
```

### HTTPS Proxy
```shell script
# self-signed certificate
./cert.sh

./https-proxy -listen=":8080" -proto=https -key server.key -pem server.pem
```

#### Test
```shell script
curl --proxy https://localhost:8080 --proxy-cacert server.pem  https://www.baidu.com
```
OR 
```shell script
curl --proxy https://localhost:8080 --proxy-insecure https://www.baidu.com
```

### HTTPS Proxy With Basic Authentication
```shell script
./https-proxy --listen=":8080" -proto=https -key server.key -pem server.pem -users "foo:123;bar:456"
```

#### Test
```shell script
curl --proxy https://foo:123@localhost:8080 --proxy-insecure https://www.baidu.com
```
