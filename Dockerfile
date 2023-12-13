FROM rust:latest
WORKDIR /usr/src/myapp
COPY . .
RUN cargo install --path .

# Expose the necessary ports
EXPOSE 8080

# Define environment variables for MongoDB connection
ENV MONGODB_URI=mongodb://localhost:27017/transfers
