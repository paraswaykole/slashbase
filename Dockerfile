FROM alpine:3.17.1

# install supervisor package
RUN apk add --no-cache supervisor

# install node and yarn packages
RUN apk add --update nodejs npm yarn

# install golang package
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN apk add --no-cache go

# Setup Folders
RUN mkdir /app /etc/supervisor.d/ /var/log/supervisor
WORKDIR /app

# Copy the source from the current directory to the working Directory inside the container
COPY ./ ./
COPY supervisord.conf /etc/supervisor/conf.d/supervisord.conf

RUN cp development.env.sample development.env

# install frontend packages and build web application front-end.
RUN cd frontend/ && yarn install && yarn build && yarn global add serve

# build backend server
RUN cd /app
RUN go build -o main .

# expose the port
EXPOSE 22022
EXPOSE 3000

# Run the executables
CMD ["/usr/bin/supervisord"]
