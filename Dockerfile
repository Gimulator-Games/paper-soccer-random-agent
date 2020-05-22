FROM ubuntu
COPY ./bin/paper-soccer-random-agent /app/agent
WORKDIR /app
CMD ["bash", "-c", "./agent"]