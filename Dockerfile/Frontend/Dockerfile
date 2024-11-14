FROM nginx:alpine

WORKDIR /usr/share/nginx/html

RUN rm -rf ./*

COPY ./ .

RUN echo '{ "apiUrl":\"'$API_URL_QA'\" }' > ./config.json

EXPOSE 80
ENTRYPOINT ["nginx", "-g", "daemon off;"]