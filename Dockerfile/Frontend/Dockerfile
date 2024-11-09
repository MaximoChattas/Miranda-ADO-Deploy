FROM node:18-alpine as base

WORKDIR /app
COPY . .

RUN npm install -g serve

EXPOSE 80

CMD ["serve", "-s", "build", "-l", "80"]