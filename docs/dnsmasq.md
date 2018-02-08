# Dnsmasq Configuration

## DNS NS Beacon
The DNS NS beacon sends an NS request for id.domain.com, where id is our implant id. Our DNS frontend needs to be configured to listen for these requests and forward them on to the backend DNS server. In addition, the DNS frontend needs to reply to other potential queries, particularly those made by defenders who are researching the domain. We can configure Dnsmasq to recognize and forward these NS requests while still responding to other legitimate queries. Install dnsmasq on the frontend server and replace the configuration file with the following.

    domain-needed
    bogus-priv
    filterwin2k

    no-resolv
    server=/*.i.domain.com/<backend-dns-server>
    host-record=www.domain.com,<web-server-ip>
    host-record=domain.com,<web-server-ip>
