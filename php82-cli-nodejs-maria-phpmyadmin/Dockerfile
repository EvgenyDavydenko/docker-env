# Use the official PHP image as the base image
FROM php:8.2.13-cli-bookworm

RUN apt update
RUN apt install -y git zip curl
RUN docker-php-ext-install pdo_mysql

# install composer
RUN curl -sS https://getcomposer.org/installer | php -- --install-dir=/usr/local/bin --filename=composer

# install nodejs
RUN curl -sL https://deb.nodesource.com/setup_16.x | bash -
RUN apt install -y nodejs

RUN usermod -u 1000 www-data && groupmod -g 1000 www-data
RUN chown -R www-data:www-data /var/www
USER "1000:1000"
