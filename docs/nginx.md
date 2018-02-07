# Nginx Configuration

## HTTP Auth Beacon
The HTTP auth beacon sends the id of our implant to the C2 server in the Authorization header. We need to configure Nginx to redirect our implant traffic to our C2 server and redirect all other traffic to a benign server. The following Nginx configuration directives will redirect any request that includes the Authorization header with a Bearer token to our C2 server and all other traffic to google.com.

    if ($http_authorization ~ "^Bearer") {
        rewrite ^(.*)$ http://c2.server.address:port/ last
    }

    return 301 https://google.com

A similar technique can be used to reroute traffic based on many other properties of the HTTP request. The Nginx rewrite rules are available at https://www.nginx.com/blog/creating-nginx-rewrite-rules/
