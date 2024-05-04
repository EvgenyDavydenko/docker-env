# Use the official PHP image as the base image
FROM php:8.2.18-cli-bookworm

RUN apt update
RUN apt install -y git zip curl \
    ca-certificates gnupg \
    libpq-dev
RUN docker-php-ext-install pdo_pgsql pgsql

# Install Composer
RUN curl -sS https://getcomposer.org/installer | php -- --install-dir=/usr/local/bin --filename=composer

# Install symfony CLI
RUN curl -sS https://get.symfony.com/cli/installer | bash
RUN mv /root/.symfony5/bin/symfony /usr/local/bin/symfony

# Install nodejs
RUN curl -fsSL https://deb.nodesource.com/gpgkey/nodesource-repo.gpg.key | gpg --dearmor -o /etc/apt/keyrings/nodesource.gpg
RUN echo "deb [signed-by=/etc/apt/keyrings/nodesource.gpg] https://deb.nodesource.com/node_20.x nodistro main" > /etc/apt/sources.list.d/nodesource.list
RUN apt update
RUN apt install -y nodejs

WORKDIR /var/www/html
# для согласования прав доступа к файлам контейнера и хост-системы
# изменим UID и GID служебного аккаунта www-data с системного до пользовательского
RUN usermod -u 1000 www-data && groupmod -g 1000 www-data
RUN chown -R www-data:www-data /var/www
USER "1000:1000"
