FROM rust:latest
WORKDIR /usr/src/myapp
COPY . .
RUN cargo install --path .
RUN cargo install cargo-watch

# Expose the necessary ports
EXPOSE 8080

# Define environment variables for MongoDB connection
ENV MONGODB_URI=mongodb://localhost:27017/transfers

# CMD ["myapp"]
