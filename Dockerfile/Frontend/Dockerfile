FROM nginx:alpine

WORKDIR /usr/share/nginx/html

RUN rm -rf ./*

COPY ./ .

RUN echo VITE_base_url="$API_URL_QA" > ./.env

EXPOSE 80
ENTRYPOINT ["nginx", "-g", "daemon off;"]