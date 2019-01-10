FROM scratch
ADD bin/hellobin /

EXPOSE 50051
CMD ["/hellobin"]
