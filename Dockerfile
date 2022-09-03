# ############################ BUILDER IMAGE ############################
FROM golang:1.18.5-buster as builder

LABEL maintainer="wahyugnc@gmail.com"

# UPDATE BUILDER IMAGE
RUN apt-get update && apt-get install -y xz-utils pkg-config libaio1 unzip

# CREATE A WORKING DIRECTORY
RUN mkdir /app
WORKDIR /app

#COPY SOURCE CODE
COPY . .

#BUILD BINARY FILE
RUN go build -ldflags '-linkmode=external' -o usedeall_backend main.go

########################### DISTRIBUTION IMAGE DEBIAN ############################
FROM debian:buster-slim

LABEL maintainer="wahyugnc@gmail.com"

# UPDATE DISTRIBUTION IMAGE
RUN apt-get update && apt-get install -y xz-utils pkg-config libaio1 unzip


# SET TIMEZONE
ENV TZ="Asia/Jakarta"
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone && dpkg-reconfigure -f noninteractive tzdata

#CREATE WORKDIR
RUN mkdir /app
WORKDIR /app

#COPY BINARY FILE FROM BUILDER
COPY --from=builder /app/usedeall_backend /app
COPY --from=builder /app/.env /app


EXPOSE 3000
CMD /app/usedeall_backend