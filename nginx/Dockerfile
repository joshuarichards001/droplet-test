# Use the official Nginx base image from Docker Hub
FROM nginx:latest

# Remove the default Nginx configuration file
RUN rm /etc/nginx/conf.d/default.conf

# Copy your custom Nginx config to the container
COPY nginx/nginx.conf /etc/nginx/conf.d/

# Expose port 80 to the host
EXPOSE 80
EXPOSE 443

# Start Nginx when the container launches
CMD ["nginx", "-g", "daemon off;"]
