# Use Ubuntu as base image
FROM ubuntu:22.04

# Avoid prompts from apt
ENV DEBIAN_FRONTEND=noninteractive

# Install necessary packages including SSH server
RUN apt-get update && apt-get install -y \
    xvfb \
    x11-utils \
    imagemagick \
    scrot \
    wget \
    curl \
    git \
    build-essential \
    libgl1-mesa-glx \
    libxrandr2 \
    libxss1 \
    libgconf-2-4 \
    libxi6 \
    libxcursor1 \
    libxcomposite1 \
    libxdamage1 \
    libxext6 \
    libxfixes3 \
    libxinerama1 \
    libxrandr2 \
    libxrender1 \
    libxss1 \
    libxtst6 \
    ca-certificates \
    fonts-liberation \
    libappindicator1 \
    libnss3 \
    lsb-release \
    xdg-utils \
    openssh-server \
    sudo \
    && rm -rf /var/lib/apt/lists/*

# Install Go
ENV GO_VERSION=1.24.0
RUN wget -q https://golang.org/dl/go${GO_VERSION}.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz && \
    rm go${GO_VERSION}.linux-amd64.tar.gz

# Set Go environment
ENV PATH="/usr/local/go/bin:${PATH}"
ENV GOPATH="/go"
ENV PATH="${GOPATH}/bin:${PATH}"

# Set display for virtual screen
ENV DISPLAY=:99

# Create directories
RUN mkdir -p /app /screenshots /var/run/sshd

# Create a user for SSH access
RUN useradd -m -s /bin/bash developer && \
    echo 'developer:password' | chpasswd && \
    usermod -aG sudo developer

# Configure SSH
RUN sed -i 's/#PermitRootLogin prohibit-password/PermitRootLogin yes/' /etc/ssh/sshd_config && \
    sed -i 's/#PasswordAuthentication yes/PasswordAuthentication yes/' /etc/ssh/sshd_config && \
    echo 'root:rootpassword' | chpasswd

# Set working directory
WORKDIR /app

# Start SSH service and keep container running
CMD service ssh start && tail -f /dev/null
