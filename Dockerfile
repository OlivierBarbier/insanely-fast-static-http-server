FROM scratch
ADD static-server /
ENV STATIC_CONTENT_PATH /static
CMD ["/static-server"]
