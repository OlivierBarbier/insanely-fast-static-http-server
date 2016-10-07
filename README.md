# insanely-fast-static-http-server
Insanely fast static http server with golang and docker

Start your static http server in seconds:

`docker run --rm -p 3000:3000 -v PATH_TO_STATIC_CONTENTS:/static hokmah/insanely-fast-static-http-server:latest`

Just replace PATH_TO_STATIC_CONTENTS with the absolute path of the static content you want to serve.

Then open your browser to http://localhost/ and voilà!
