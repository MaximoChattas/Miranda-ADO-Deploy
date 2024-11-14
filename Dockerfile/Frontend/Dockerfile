FROM nginx:alpine

WORKDIR /usr/share/nginx/html

RUN rm -rf ./*

COPY ./ .

CMD sh -c echo '{ "apiUrl":"'$API_URL'" }' > ./config.json

EXPOSE 80
ENTRYPOINT ["nginx", "-g", "daemon off;"]