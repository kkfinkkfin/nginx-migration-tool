# NGINX Ingress Controller Migration Tool

The NGINX Ingress Controller Migration Tool use for migration of the [Kubernetes ingress-nginx](https://kubernetes.github.io/ingress-nginx/) configurations yaml for to [Nginx Ingress Controller](https://github.com/nginxinc/kubernetes-ingress).
 

## Overview

The annotations of the [Kubernetes ingress-nginx](https://kubernetes.github.io/ingress-nginx/) uses that [Nginx Ingress Controller](https://github.com/nginxinc/kubernetes-ingress) do use or do it different. This tool help to automate migration of the configurations to yaml of [Nginx Ingress Controller](https://github.com/nginxinc/kubernetes-ingress).

There is three option of migrating using NGINX Ingress Resources

1, Convert Ingress to Ingress or CRD

•For functions that can be achieved through Ingress resource types and annotations, it is supported to convert to the same Ingress resource type and use NGINX annotations.
•For advanced functions such as canary functionality, the NGINX Ingress Controller cannot be implemented through ingress with annotations, and is transformed into the resource type of CRD.
•It is also supported, and all functions are converted to CRD resources.
![image](https://user-images.githubusercontent.com/59547386/171353803-e8a68e20-dadc-4bd4-8134-6e22e3be94b0.png)

•It is also supported, and all functions are converted to CRD resources.

2, New Ingess/CRD resource with  specific IngressClass Name

•Do not modify the existing Ingress resources in the cluster and would not affect the existing access traffic.
•A new set of Ingress resources for NGINX ICs can be used with user-specified IngressClass names so that ICs that also declare the --ingress-class can watch to this set of new Ingress resources.
![image](https://user-images.githubusercontent.com/59547386/171353852-b4e9af0b-8ea4-4465-8e58-c8bcc01db4d0.png)


3, New Mergeable Ingress resource with  specific IngressClass Name

•Do not modify the existing Ingress resources in the cluster and would not affect the existing access traffic.
•In the CE, there are multiple ingresses with the same hostname with different paths, and we can't convert directly due to host collision detection, so we need to identify them first and then convert them to Mergeable Ingress resources.
![image](https://user-images.githubusercontent.com/59547386/171353885-e84e4b68-4770-4721-8253-dfe9a795750c.png)


## Showcase example
![image](https://user-images.githubusercontent.com/59547386/171353909-e7818c5b-2d8c-4b53-a0a1-3ecf2547a3e2.png)

![image](https://user-images.githubusercontent.com/59547386/171353953-ba03c0a3-fe66-457b-bc05-e2852e7c7cc6.png)



## Getting Started

In this section, we show how to quickly run NGINX Ingress Controller Migration tool.

### A Note about NGINX Ingress Controller

If you’d like to use the tool with [NGINX Ingress Controller](https://github.com/nginxinc/kubernetes-ingress/) for Kubernetes, see [this doc](https://docs.nginx.com/nginx-ingress-controller/logging-and-monitoring/prometheus/) for the installation instructions.

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
`nginx.ingress.kubernetes.io/app-root` | `nginx.org/server-snippets: "if($request_uri = '/'){ \n return 302 $http_x_forward_proto://$host{{ $location.Rewrite.AppRoot }}; \n }"` |
`nginx.ingress.kubernetes.io/configuration-snippet` | `nginx.org/location-snippets` |
`nginx.ingress.kubernetes.io/proxy-body-size` | `nginx.org/client-max-body-size` | 
`nginx.ingress.kubernetes.io/proxy-connect-timeout` | `nginx.org/proxy-connect-timeout` | 
`nginx.ingress.kubernetes.io/proxy-send-timeout` | `nginx.org/proxy-send-timeout` |
`nginx.ingress.kubernetes.io/proxy-read-timeout` | `nginx.org/proxy-read-timeout` |
`nginx.ingress.kubernetes.io/rewrite-target` | `nginx.org/rewrites` | 
`nginx.ingress.kubernetes.io/server-snippet` | `nginx.org/server-snippets` | 
`nginx.ingress.kubernetes.io/ssl-redirect`   | `ingress.kubernetes.io/ssl-redirect` | 
`nginx.ingress.kubernetes.io/stream-snippet` | `nginx.org/stream-snippets` | 
`nginx.ingress.kubernetes.io/upstream-hash-by: "$request_url"` | `nginx.org/lb-method: "hash $request_uri consistent"` | 
`nginx.ingress.kubernetes.io/load-balance` | `nginx.org/lb-method` | 
`nginx.ingress.kubernetes.io/proxy-buffering` | `nginx.org/proxy-buffering`|
`nginx.ingress.kubernetes.io/proxy-buffers-number` | `nginx.org/proxy-buffers`|
`nginx.ingress.kubernetes.io/proxy-buffer-size` | `nginx.org/proxy-buffer-size`|
`nginx.ingress.kubernetes.io/proxy-max-temp-file-size` | `nginx.org/proxy-max-temp-file-size` | 

#### Cookie
Community Ingress Controller | NGINX Ingress Controller
----|----|

#### Redirect
Community Ingress Controller | NGINX Ingress Controller
----|----|

### Advanced annotations in Ingress type with snippets:
Community Ingress Controller | NGINX Ingress Controller
----|----|
`nginx.ingress.kubernetes.io/app-root` | `nginx.org/server-snippets: "if($request_uri = '/'){ \n return 302 $http_x_forward_proto://$host{{ $location.Rewrite.AppRoot }}; \n }"` |
`nginx.ingress.kubernetes.io/client-body-buffer-size` | `"nginx.org/location-snippets": "client_body_buffer_size 30M; "` |
`nginx.ingress.kubernetes.io/custom-http-errors: "code"` | `nginx.org/location-snippets: /| \n location / {\n   error_page 404 = @fallback;\n}` |
`nginx.ingress.kubernetes.io/default-backend: "default-svc"` | `nginx.org/location-snippets: /| \n location @fallback {\n   proxy_pass http://backend;\n}` |
`nginx.ingress.kubernetes.io/proxy-cookie-domain` | `nginx.org/location-snippets: proxy_cookie_domain off;` | 
`nginx.ingress.kubernetes.io/proxy-cookie-path` | `nginx.org/location-snippets: proxy_cookie_path off;` | 

#### Authentication
Community Ingress Controller | NGINX Ingress Controller
----|----|

#### External Authentication
Community Ingress Controller | NGINX Ingress Controller
----|----|

#### CORS
Community Ingress Controller | NGINX Ingress Controller
----|----|

#### Rate Limiting
Community Ingress Controller | NGINX Ingress Controller
----|----|

#### Mirror
Community Ingress Controller | NGINX Ingress Controller
----|----|

### Advanced annotations in CRD type
For the following scenarios, it is recommended to convert to CRD type resource objects, so that the configuration is more clearer than ingress.

#### Canary
Community Ingress Controller | NGINX Ingress Controller
----|----|

#### CORS
Community Ingress Controller | NGINX Ingress Controller
----|----|

#### Rate Limiting
Community Ingress Controller | NGINX Ingress Controller
----|----|


####Backend Certificate Authentication
Community Ingress Controller | NGINX Ingress Controller
----|----|


### Megertable Ingresses

### Other that can not be convert
####ModSecurity
Community Ingress Controller | NGINX Ingress Controller
----|----|
`nginx.ingress.kubernetes.io/enable-modsecurity` | - |
`nginx.ingress.kubernetes.io/enable-owasp-core-rules` | - | 
`nginx.ingress.kubernetes.io/modsecurity-transaction-id`| - |
`nginx.ingress.kubernetes.io/modsecurity-snippet` | - | 



## Troubleshooting

The program logs errors to the standard output. When using Docker, if it doesn’t work as expected, check its logs using [docker logs](https://docs.docker.com/engine/reference/commandline/logs/) command.

## Releases

### Docker images
It would be uploaded to [DockerHub](https://hub.docker.com/r/nginx/nginx-prometheus-exporter/) and [GitHub Container](https://github.com/nginxinc/nginx-prometheus-exporter/pkgs/container/nginx-prometheus-exporter). But not now yet, please waiting. 


### Binaries


## Building the Migration tool


