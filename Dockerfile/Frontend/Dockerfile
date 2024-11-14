FROM nginx:alpine

WORKDIR /usr/share/nginx/html

RUN rm -rf ./*

COPY ./ .

#CMD sh -c echo '{ "apiUrl":"'$API_URL'" }' > ./config.json
#CMD sh -c echo '{ "apiUrl":"https://miranda-back-qa.azurewebsites.net" }' > /usr/share/nginx/html/config.json

EXPOSE 80
#ENTRYPOINT ["nginx", "-g", "daemon off;"]
CMD sh -c echo '{ "apiUrl":"https://miranda-back-qa.azurewebsites.net" }' > /usr/share/nginx/html/config.json && nginx -g "daemon off;"'