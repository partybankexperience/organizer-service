# Specifies a parent image
FROM golang:1.19.2-bullseye
 
# Creates an rave-app directory to hold your rave-app’s source code
WORKDIR /rave
 
# Copies everything from your root directory into /rave-app
COPY . .
 
# Installs Go dependencies
RUN go mod download
 
# Builds your rave-app with optional configuration
RUN go build -o /rave bin .
 
# Tells Docker which network port your container listens on
EXPOSE 8082
 
# Specifies the executable command that runs when the container starts
CMD ["/rave/bin"]
