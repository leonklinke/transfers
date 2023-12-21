FROM golang:latest
WORKDIR /usr/src/api
COPY . .
RUN go build main.go

# Expose the necessary ports
EXPOSE 8080

# Define environment variables for MongoDB connection
ENV MONGODB_URI=mongodb://localhost:27017/transfers
