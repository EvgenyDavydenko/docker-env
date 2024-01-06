# Use the official PHP image as the base image
FROM php:8.2.13-cli-bookworm
RUN apt update
RUN apt install -y git zip curl

# Install Composer
RUN curl -sS https://getcomposer.org/installer | php -- --install-dir=/usr/local/bin --filename=composer

WORKDIR /var/www/html
RUN usermod -u 1000 www-data && groupmod -g 1000 www-data
RUN chown -R www-data:www-data /var/www
USER "1000:1000"
