vcl 4.0;

import std;

acl purge {
  "host.docker.internal";
}

backend default {
  .host = "host.docker.internal";
  .port = "8090";
}

sub vcl_recv {
  unset req.http.X-Body-Len;
  # Only allow BAN requests from IP addresses in the 'purge' ACL.
  if (req.method == "BAN") {
    # Same ACL check as above:
    if (!client.ip ~ purge) {
      return (synth(403, "Not allowed."));
    }

    # Logic for the ban, using the X-Cache-Tags header.
    if (req.http.X-Cache-Tag) {
      ban("obj.http.x-cache-tags ~ " + req.http.X-VS-Cache-Tag);
      return (synth(200, "Ban added."));
    }

    # Throw a synthetic page so the request won't go to the backend.
    return (synth(403, "Nothing to do"));

  }

   if (req.method == "GET") {
     return (hash);
   }

  return (pipe);

}

sub vcl_hash {
  hash_data("");
}

sub vcl_backend_response {
    # Set ban-lurker friendly custom headers.
    if (beresp.http.content-type ~ "json") {
      set beresp.do_gzip = true;
    }
    set beresp.http.X-Url = bereq.url;
    set beresp.http.X-Host = bereq.http.host;
}

sub vcl_deliver {
    if (obj.hits > 0) {
      set resp.http.X-Cache = "Hit";
      set resp.http.X-Cache-Hits = obj.hits;
    } else {
      set resp.http.X-Cache = "Miss";
    }
    unset resp.http.X-Varnish;
    unset resp.http.Via;
    unset resp.http.Age;
    unset resp.http.X-Purge-URL;
    unset resp.http.X-Purge-Host;
    # Remove ban-lurker friendly custom headers when delivering to client.
    unset resp.http.X-Url;
    unset resp.http.X-Host;
    # Comment these for easier Drupal cache tag debugging in development.
//    unset resp.http.X-Cache-Tags;
    unset resp.http.X-Cache-Contexts;
}