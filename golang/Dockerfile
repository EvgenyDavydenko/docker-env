# Use the official Golang image as the base image
FROM golang:1.22.2-bookworm
RUN apt update

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /var/www/applic

# Создаем нового пользователя 'golang' и группу 'golang'
RUN usermod -u 1000 www-data && groupmod -g 1000 www-data
RUN chown -R www-data:www-data /var/www

USER "1000:1000"