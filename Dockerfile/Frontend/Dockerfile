FROM nginx:alpine

WORKDIR /usr/share/nginx/html

RUN rm -rf ./*

COPY ./ .

# Exponer el puerto 80 para servir la aplicación Angular
EXPOSE 80
CMD sh -c 'nginx -g "daemon off;"'