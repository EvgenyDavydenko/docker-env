# Use the official PHP image as the base image
FROM php:7.4-apache-bullseye

RUN apt update
RUN apt install -y git zip curl

# Install Composer
RUN curl -sS https://getcomposer.org/installer | php -- --install-dir=/usr/local/bin --filename=composer

# configure Apache httpd
RUN sed -i '/<Directory "\/var\/www">/,/<\/Directory>/ s/AllowOverride None/AllowOverride All/' /etc/apache2/apache2.conf
RUN a2enmod rewrite

RUN usermod -u 1000 www-data && groupmod -g 1000 www-data
RUN chown -R www-data:www-data /var/www
USER "1000:1000"