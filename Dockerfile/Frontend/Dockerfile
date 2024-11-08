FROM node
WORKDIR /frontend
COPY package.json .
COPY package-lock.json .
COPY . .
EXPOSE 5173
CMD ["npm", "run", "dev"]