# NGINX Ingress Controller Migration Tool

The NGINX Ingress Controller Migration Tool use for migration of the [Kubernetes ingress-nginx](https://kubernetes.github.io/ingress-nginx/) configurations YAML to [Nginx Ingress Controller](https://github.com/nginxinc/kubernetes-ingress).
 

## Overview

The annotations of the [Kubernetes ingress-nginx](https://kubernetes.github.io/ingress-nginx/) uses that [Nginx Ingress Controller](https://github.com/nginxinc/kubernetes-ingress) do use or do it different. This tool help to automate migration of the configurations to yaml of [Nginx Ingress Controller](https://github.com/nginxinc/kubernetes-ingress).

There is three option of migrating using NGINX Ingress Resources

1, Convert Ingress to Ingress or CRD

•For functions that can be achieved through Ingress resource types and annotations, it is supported to convert to the same Ingress resource type and use NGINX annotations.
<br>•For advanced functions such as canary functionality, the NGINX Ingress Controller cannot be implemented through ingress with annotations, and is transformed into the resource type of CRD.
<br>•It is also supported, and all functions are converted to CRD resources.

![image](https://user-images.githubusercontent.com/59547386/171353803-e8a68e20-dadc-4bd4-8134-6e22e3be94b0.png)


2, New Ingess/CRD resource with  specific IngressClass Name

•Do not modify the existing Ingress resources in the cluster and would not affect the existing access traffic.
<br>•A new set of Ingress resources for NGINX ICs can be used with user-specified IngressClass names so that ICs that also declare the --ingress-class can watch to this set of new Ingress resources.

![image](https://user-images.githubusercontent.com/59547386/171353852-b4e9af0b-8ea4-4465-8e58-c8bcc01db4d0.png)


3, New Mergeable Ingress resource with  specific IngressClass Name

•Do not modify the existing Ingress resources in the cluster and would not affect the existing access traffic.
<br>•In the CE, there are multiple ingresses with the same hostname with different paths, and we can't convert directly due to host collision detection, so we need to identify them first and then convert them to Mergeable Ingress resources.

![image](https://user-images.githubusercontent.com/59547386/171353885-e84e4b68-4770-4721-8253-dfe9a795750c.png)


## Showcase example

![image](https://user-images.githubusercontent.com/59547386/171353909-e7818c5b-2d8c-4b53-a0a1-3ecf2547a3e2.png)


![image](https://user-images.githubusercontent.com/59547386/171353953-ba03c0a3-fe66-457b-bc05-e2852e7c7cc6.png)



## Getting Started

In this section, we show how to quickly run NGINX Ingress Controller Migration tool.

### A Note about NGINX Ingress Controller

If you’d like to use the tool with [NGINX Ingress Controller](https://github.com/nginxinc/kubernetes-ingress/) for Kubernetes, see [this doc](https://docs.nginx.com/nginx-ingress-controller/) for the installation instructions.

### Prerequisites

We assume that you have already installed NGINX or NGINX Plus Ingress Controller. 

## Usage

### Command-line Arguments

```
Usage of ./nginx-migration-tool:
  -original_ingress_name
        Define the original ingress to be converted. Default name is "test" temporarily. 
  -namespace_name
        Defines the namespace in which the original inrgess to be converted is located. Default is the "default" namespace.
  -new_namespace_name
        Defines the namespace of the new ingrerss resource object after conversion. Default is the "default" namespace.
  -new_ingress_classname
        Defines the inress class name of the new ingrerss resource object after conversion. Default is "nginx-plus" temporarily. 
  -new_ingress_name
        Defines the inress name of the new ingrerss resource object after conversion. Default is "test-migration-edtion" temporarily. 
  
```

## Different configuration

### Advanced annotations in Ingress type with corresponding keys:
Community Ingress Controller | NGINX Ingress Controller
----|----|
nginx.ingress.kubernetes.io/configuration-snippet | nginx.org/location-snippets |
nginx.ingress.kubernetes.io/proxy-body-size | nginx.org/client-max-body-size | 
nginx.ingress.kubernetes.io/proxy-connect-timeout | nginx.org/proxy-connect-timeout | 
nginx.ingress.kubernetes.io/proxy-send-timeout | nginx.org/proxy-send-timeout |
nginx.ingress.kubernetes.io/proxy-read-timeout | nginx.org/proxy-read-timeout |
nginx.ingress.kubernetes.io/rewrite-target | nginx.org/rewrites` | 
nginx.ingress.kubernetes.io/server-snippet | nginx.org/server-snippets | 
nginx.ingress.kubernetes.io/ssl-redirect   | ingress.kubernetes.io/ssl-redirect | 
nginx.ingress.kubernetes.io/stream-snippet | nginx.org/stream-snippets | 
nginx.ingress.kubernetes.io/upstream-hash-by: "$request_url" | nginx.org/lb-method: "hash $request_uri consistent" | 
nginx.ingress.kubernetes.io/load-balance | nginx.org/lb-method | 
nginx.ingress.kubernetes.io/proxy-buffering | nginx.org/proxy-buffering|
nginx.ingress.kubernetes.io/proxy-buffers-number | nginx.org/proxy-buffers|
nginx.ingress.kubernetes.io/proxy-buffer-size | nginx.org/proxy-buffer-size|
nginx.ingress.kubernetes.io/proxy-max-temp-file-size | nginx.org/proxy-max-temp-file-size | 

#### Cookie (Plus only)
The sticky session function of the Community Ingress Controller is implemented through [LUA](https://github.com/kubernetes/ingress-nginx/blob/main/rootfs/etc/nginx/lua/balancer/sticky.lua), and the sticky session function of the NGINX Ingress Controller is a commercially limited version, so this section does not apply to the NGINX open source version.
Community Ingress Controller | NGINX Ingress Controller(Ingress) | NGINX Ingress Controller(CRD)
----|----|----|
nginx.ingress.kubernetes.io/affinity: "cookie"<br>nginx.ingress.kubernetes.io/affinity-mode: "balanced" \| "persistent"<br>nginx.ingress.kubernetes.io/affinity-canary-behavior: "sticky" \| "legacy"<br>nginx.ingress.kubernetes.io/session-cookie-name: "cookieName"<br>nginx.ingress.kubernetes.io/session-cookie-expires: "2"<br>nginx.ingress.kubernetes.io/session-cookie-path: "/example"<br>nginx.ingress.kubernetes.io/session-cookie-secure: “true”<br>nginx.ingress.kubernetes.io/session-cookie-change-on-failure: "true" \| "false"<br>nginx.ingress.kubernetes.io/session-cookie-samesite<br>nginx.ingress.kubernetes.io/session-cookie-conditional-samesite-none: "true" | nginx.com/sticky-cookie-services: "serviceName=example-svc cookie_name expires=1h domain=.example.com httponly samesite=strict\|lax\|none secure path=/example" | name: example<br>service: example-svc<br>port: 80<br>sessionCookie:<br>&emsp;&emsp;enable: true<br>&emsp;&emsp;name: cookieName<br>&emsp;&emsp;path: /example<br>&emsp;&emsp;expires: 2h<br>&emsp;&emsp;domain: .example.com<br>&emsp;&emsp;httpOnly: true<br>&emsp;&emsp;secure: true | 

#### Redirect
The redirect function of the Community Ingress Controller is implemented through LUA. Therefor is just simple convert it to rewrites of NGINX Ingress Controller.
Community Ingress Controller | NGINX Ingress Controller
----|----|
nginx.ingress.kubernetes.io/force-ssl-redirect: "true" | nginx.org/redirect-to-https, ingress.kubernetes.io/ssl-redirect
nginx.ingress.kubernetes.io/from-to-www-redirect: "true" | 
nginx.ingress.kubernetes.io/permanent-redirect: "http://www.google.com" |
nginx.ingress.kubernetes.io/permanent-redirect-code: "308" |
nginx.ingress.kubernetes.io/temporal-redirect |
nginx.ingress.kubernetes.io/rewrite-target: "URI" | nginx.org/rewrites |
nginx.ingress.kubernetes.io/enable-rewrite-log |
nginx.ingress.kubernetes.io/ssl-redirect: "true,false" | ingress.kubernetes.io/ssl-redirect: "True,False" |

### Advanced annotations in Ingress type with snippets:
Community Ingress Controller | NGINX Ingress Controller
----|----|
nginx.ingress.kubernetes.io/app-root | nginx.org/server-snippets: "if($request_uri = '/'){   <br> return 302 $http_x_forward_proto://$host{{ $location.Rewrite.AppRoot }};   <br> }" |
nginx.ingress.kubernetes.io/client-body-buffer-size | "nginx.org/location-snippets": "client_body_buffer_size 30M; "` |
nginx.ingress.kubernetes.io/custom-http-errors: "code" | nginx.org/location-snippets: \|   <br> location / {  <br>   error_page 404 = @fallback;  <br>} |
nginx.ingress.kubernetes.io/default-backend: "default-svc" | nginx.org/location-snippets: \|   <br> location @fallback {  <br>   proxy_pass http://backend;  <br>} |
nginx.ingress.kubernetes.io/proxy-cookie-domain | nginx.org/location-snippets: proxy_cookie_domain off; | 
nginx.ingress.kubernetes.io/proxy-cookie-path | nginx.org/location-snippets: proxy_cookie_path off; | 

#### CORS
Community Ingress Controller | NGINX Ingress Controller
----|----|
nginx.ingress.kubernetes.io/enable-cors: "true"<br>nginx.ingress.kubernetes.io/cors-allow-origin: "*"<br>nginx.ingress.kubernetes.io/cors-allow-methods: "PUT,GET,POST,OPTIONS"<br>nginx.ingress.kubernetes.io/cors-allow-headers: "DNT,Keep-Alive,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,Range,Authorization"<br>nginx.ingress.kubernetes.io/cors-expose-headers: ""<br>nginx.ingress.kubernetes.io/cors-allow-origin: "*"<br>nginx.ingress.kubernetes.io/cors-allow-credentials："true"<br>nginx.ingress.kubernetes.io/cors-max-age: "600" | nginx.org/server-snippets: "add_header {{ $h.Name }} "{{ $h.Value }}" always;" | 

#### SSL Passthrough
The Community Ingress Controller to send TLS connections directly to the backend instead of letting NGINX decrypt the communication.
Community Ingress Controller | NGINX Ingress Controller
----|----|
nginx.ingress.kubernetes.io/ssl-passthrough: [true\|false] | command:- -enable-tls-passthrough=true


#### Mirror
Community Ingress Controller | NGINX Ingress Controller
----|----|
nginx.ingress.kubernetes.io/mirror-request-body: "off" | location-snippets |
nginx.ingress.kubernetes.io/mirror-target: https://test.env.com/$request_uri |





### Advanced annotations in CRD type
For the following scenarios, it is recommended to convert to CRD type resource objects, so that the configuration is more clearer than ingress.

#### Redirect
The redirect function of the Community Ingress Controller is implemented through LUA. Therefor is just simple convert it to rewrites of NGINX Ingress Controller.
Community Ingress Controller | NGINX Ingress Controller
----|----|
nginx.ingress.kubernetes.io/rewrite-target: "URI" | rewritePath: /beans |
nginx.ingress.kubernetes.io/ssl-redirect: "true,false" | VirtualServer.TLS.Redirect<br>enable: true<br>code: 301<br>basedOn: scheme

#### Canary
Community Ingress Controller | NGINX Ingress Controller
----|----|
nginx.ingress.kubernetes.io/canary: "true" | 
nginx.ingress.kubernetes.io/canary-by-header: "always,never,custom" | matches:<br>- conditions:<br>&emsp;&emsp;- header: httpHeader<br>&emsp;&emsp;value: never<br>action:<br>&emsp;&emsp;pass: echo<br>&emsp;&emsp;- header: httpHeader<br>&emsp;&emsp;value: always<br>&emsp;&emsp;action:<br>&emsp;&emsp;pass: echo-canary<br>action:<br>pass: echo
nginx.ingress.kubernetes.io/canary-by-header-value: custom_value | matches:<br>- conditions:<br>&emsp;&emsp;- header: httpHeader<br>&emsp;&emsp;value: my-value<br>action:<br>&emsp;&emsp;pass: echo-canary<br>action:<br>pass: echo
nginx.ingress.kubernetes.io/canary-by-header-pattern: regex | - |
nginx.ingress.kubernetes.io/canary-by-cookie: "always,never,custom" | matches:<br>- conditions:<br>&emsp;&emsp;- cookie: cookieName<br>&emsp;&emsp;value: never<br>&emsp;&emsp;action:<br>&emsp;&emsp;pass: echo<br>&emsp;&emsp;- cookie: cookieName<br>&emsp;&emsp;value: always<br>&emsp;&emsp;action:<br>&emsp;&emsp;pass: echo-canary<br>action:<br>&emsp;&emsp;pass: echo
nginx.ingress.kubernetes.io/canary-weight: "0-100"<br>nginx.ingress.kubernetes.io/canary-weight-total: "100,custom" | splits:<br>- weight: 90<br>&emsp;&emsp;action:<br>&emsp;&emsp;pass: echo<br>- weight: 10<br>&emsp;&emsp;action:<br>&emsp;&emsp;pass: echo-canary

#### CORS
Community Ingress Controller | NGINX Ingress Controller
----|----|
nginx.ingress.kubernetes.io/enable-cors: "true"<br>nginx.ingress.kubernetes.io/cors-allow-origin: "*"<br>nginx.ingress.kubernetes.io/cors-allow-methods: "PUT,GET,POST,OPTIONS"<br>nginx.ingress.kubernetes.io/cors-allow-headers: \|<br>&emsp;&emsp;"DNT,Keep-Alive,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,Range,Authorization"<br>nginx.ingress.kubernetes.io/cors-expose-headers: ""<br>nginx.ingress.kubernetes.io/cors-allow-origin: "*"<br>nginx.ingress.kubernetes.io/cors-allow-credentials："true"<br>nginx.ingress.kubernetes.io/cors-max-age: "600" | responseHeaders:<br>&emsp;&emsp;add:<br>&emsp;&emsp;- name: Access-Control-Allow-Origin<br>&emsp;&emsp;value: "*"<br>&emsp;&emsp;- name: Access-Control-Allow-Credentials<br>&emsp;&emsp;value: "true"<br>&emsp;&emsp;- name: Access-Control-Allow-Methods<br>&emsp;&emsp;value: "PUT,GET,POST"<br>OPTIONS:<br>&emsp;&emsp;- name: Access-Control-Allow-Headers<br>&emsp;&emsp;value: "X-Forward-For"<br>&emsp;&emsp;- name: Access-Control-Max-Age<br>&emsp;&emsp;value: "600" |

#### Rate Limiting
Community Ingress Controller | NGINX Ingress Controller
----|----|
nginx.ingress.kubernetes.io/limit-connections<br>nginx.ingress.kubernetes.io/limit-rps: "number"<br>nginx.ingress.kubernetes.io/limit-rpm:“number”<br>nginx.ingress.kubernetes.io/limit-burst-multiplier: “multiplier”<br>nginx.ingress.kubernetes.io/limit-rate:“number”<br>nginx.ingress.kubernetes.io/limit-rate-after: “number”<br>nginx.ingress.kubernetes.io/limit-whitelist: “CIDR”| rateLimit:<br>&emsp;&emsp;rate: {number}r/m<br>&emsp;&emsp;burst: {number} * {multiplier}<br>&emsp;&emsp;key: ${binary_remote_addr}<br>&emsp;&emsp;zoneSize: 5m |

#### Authentication
Community Ingress Controller | NGINX Ingress Controller
----|----|
nginx.ingress.kubernetes.io/auth-type: [basic\|digest]<br>nginx.ingress.kubernetes.io/auth-secret: secretName<br>nginx.ingress.kubernetes.io/auth-secret-type: [auth-file\|auth-map]<br>nginx.ingress.kubernetes.io/auth-realm: "realm string" | nginx.org/location-snippets | 

#### Client Certificate Authentication
Community Ingress Controller | NGINX Ingress Controller
----|----|
nginx.ingress.kubernetes.io/auth-tls-secret: namespace/secretName<br>nginx.ingress.kubernetes.io/auth-tls-verify-depth: "1"<br>nginx.ingress.kubernetes.io/auth-tls-verify-client: "on,off,optional,optional_no_ca"<br>nginx.ingress.kubernetes.io/auth-tls-error-page<br>nginx.ingress.kubernetes.io/auth-tls-pass-certificate-to-upstream: "true,false" | ingressMTLS:<br>&emsp;&emsp;clientCertSecret: secretName<br>&emsp;&emsp;verifyClient: “on”<br>&emsp;&emsp;verifyDepth: 1 |


#### External Authentication
Community Ingress Controller | NGINX Ingress Controller
----|----|
nginx.ingress.kubernetes.io/auth-url: "URL to the authentication service"<br>nginx.ingress.kubernetes.io/auth-keepalive<br>nginx.ingress.kubernetes.io/auth-keepalive-requests<br>nginx.ingress.kubernetes.io/auth-keepalive-timeout<br>nginx.ingress.kubernetes.io/auth-method<br>nginx.ingress.kubernetes.io/auth-signin<br>nginx.ingress.kubernetes.io/auth-signin-redirect-param<br>ginx.ingress.kubernetes.io/auth-response-headers<br>nginx.ingress.kubernetes.io/auth-proxy-set-headers<br>nginx.ingress.kubernetes.io/auth-request-redirect<br>nginx.ingress.kubernetes.io/auth-cache-key<br>nginx.ingress.kubernetes.io/auth-cache-duration<br>nginx.ingress.kubernetes.io/auth-always-set-cookie<br>nginx.ingress.kubernetes.io/auth-snippet | nginx.org/location-snippets |

#### Backend Certificate Authentication
Community Ingress Controller | NGINX Ingress Controller
----|----|
nginx.ingress.kubernetes.io/proxy-ssl-secret: "secretName"<br>nginx.ingress.kubernetes.io/proxy-ssl-ciphers: "DEFAULT"<br>nginx.ingress.kubernetes.io/proxy-ssl-name: "server-name"<br>nginx.ingress.kubernetes.io/proxy-ssl-protocols: "TLSv1.2"<br>nginx.ingress.kubernetes.io/proxy-ssl-verify: "on,off"<br>nginx.ingress.kubernetes.io/proxy-ssl-verify-depth: "1"<br>nginx.ingress.kubernetes.io/proxy-ssl-server-name: "on,off" | egressMTLS:<br>&emsp;&emsp;tlsSecret: secretName<br>&emsp;&emsp;verifyServer: true\|false<br>&emsp;&emsp;verifyDepth: 1<br>&emsp;&emsp;protocols: TLSv1.2<br>&emsp;&emsp;ciphers: DEFAULT<br>&emsp;&emsp;sslName: server-name<br>&emsp;&emsp;serverName: true\|false |


### Megertable Ingresses
to-do

### Other that can not be convert

#### ModSecurity
Community Ingress Controller | NGINX Ingress Controller
----|----|
nginx.ingress.kubernetes.io/enable-modsecurity | - |
nginx.ingress.kubernetes.io/enable-owasp-core-rules | - | 
nginx.ingress.kubernetes.io/modsecurity-transaction-id| - |
nginx.ingress.kubernetes.io/modsecurity-snippet | - | 

#### Global Rate Limiting
Community Ingress Controller | NGINX Ingress Controller
----|----|
nginx.ingress.kubernetes.io/global-rate-limit | - |
nginx.ingress.kubernetes.io/global-rate-limit-window | - |
nginx.ingress.kubernetes.io/global-rate-limit-key | - |
nginx.ingress.kubernetes.io/global-rate-limit-ignored-cidrs | - |


#### Influxdb
Community Ingress Controller | NGINX Ingress Controller
----|----|
nginx.ingress.kubernetes.io/enable-influxdb: "true|false" | - |
nginx.ingress.kubernetes.io/influxdb-measurement: "nginx-reqs" | - |
nginx.ingress.kubernetes.io/influxdb-port: "8089" | - |
nginx.ingress.kubernetes.io/influxdb-host: "127.0.0.1" | - |
nginx.ingress.kubernetes.io/influxdb-server-name: "nginx-ingress" | - |


## Troubleshooting

The program logs errors to the standard output. When using Docker, if it doesn’t work as expected, check its logs using [docker logs](https://docs.docker.com/engine/reference/commandline/logs/) command.


