FROM scratch
ADD bin/timebin /

EXPOSE 50051
CMD ["/timebin"]

