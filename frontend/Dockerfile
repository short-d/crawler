# Create production container
FROM nginx:1.17.5-alpine

WORKDIR /usr/share/nginx/html

COPY build/ .

RUN rm -rf /etc/nginx/conf.d/default.conf

COPY nginx/default.conf /etc/nginx/conf.d/