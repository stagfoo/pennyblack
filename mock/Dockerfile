FROM alpine:latest

WORKDIR /

# Install runtime dependencies + SSH server
RUN apk add --no-cache xvfb maim xdotool openssh bash \
    xorg-server xwininfo imagemagick

# Set up SSH
RUN ssh-keygen -A
RUN echo 'root:password' | chpasswd
RUN sed -i 's/#PermitRootLogin prohibit-password/PermitRootLogin yes/' /etc/ssh/sshd_config
RUN sed -i 's/#PasswordAuthentication yes/PasswordAuthentication yes/' /etc/ssh/sshd_config

# Copy the pre-built binary
COPY main /app/main

# Copy the entrypoint script
COPY mock/entrypoint.sh /app/entrypoint.sh
RUN chmod +x /app/entrypoint.sh

# Expose SSH port and display port
EXPOSE 22 99

# Start SSH daemon instead of the entrypoint script
CMD ["/usr/sbin/sshd", "-D"]