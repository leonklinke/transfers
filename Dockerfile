FROM rust:1.67 as builder
WORKDIR /usr/src/myapp
COPY . .
RUN cargo install --path .

FROM debian:bullseye-slim
RUN apt-get update && apt-get install -y  && rm -rf /var/lib/apt/lists/*
COPY --from=builder /usr/local/cargo/bin/myapp /usr/local/bin/myapp

# Expose the necessary ports
EXPOSE 8080

# Define environment variables for MongoDB connection
ENV MONGODB_URI=mongodb://localhost:27017/transfers

CMD ["myapp"]
