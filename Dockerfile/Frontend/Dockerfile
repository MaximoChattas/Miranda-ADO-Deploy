FROM nginx:alpine

WORKDIR /usr/share/nginx/html

RUN rm -rf ./*

COPY ./ .

EXPOSE 80

CMD ["/bin/sh", "-c", "echo '{ \"apiUrl\":\"'$API_URL'\" }' > /usr/share/nginx/html/config.json && nginx -g 'daemon off;'"]