# Use the official Python image as the base image
FROM python:3.10.13-bookworm

RUN apt update
#RUN apt install -y python3-venv

# Установите рабочую директорию в контейнере
WORKDIR /var/www/applic

# Установите переменные среды
ENV PATH="/var/www/.local/bin:${PATH}"

RUN usermod -u 1000 www-data && groupmod -g 1000 www-data
RUN chown -R www-data:www-data /var/www
USER "1000:1000"
# RUN pip install -r requirements.txt
