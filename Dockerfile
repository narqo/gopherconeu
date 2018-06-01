FROM scratch

ENV PORT=8000 \
    HEALTH_PORT=8001

EXPOSE $PORT $HEALTH_PORT

COPY ./bin/linux-amd64/gopherconeu /
CMD ["/gopherconeu"]
