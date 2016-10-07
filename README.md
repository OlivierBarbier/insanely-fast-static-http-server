# Insanely fast static http server

Start your static http server in seconds:

`docker run -d -p 3000:3000 -v PATH_TO_STATIC_CONTENT:/static hokmah/insanely-fast-static-http-server:latest`

Just replace PATH_TO_STATIC_CONTENT with the absolute path of the static content you want to serve.

Then open your browser to http://localhost:3000/ and voilà!
