FROM nginx:alpine

WORKDIR /usr/share/nginx/html

RUN rm -rf ./*

COPY ./ .

#RUN echo VITE_base_url='$API_URL_QA' > ./.env
RUN echo VITE_base_url="https://miranda-back-qa.azurewebsites.net" > ./.env

EXPOSE 80
ENTRYPOINT ["nginx", "-g", "daemon off;"]