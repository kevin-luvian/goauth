# Use node alpine as it's a small node image
FROM node:alpine

RUN mkdir -p /app

WORKDIR /app

COPY package*.json /app

RUN yarn install

COPY . /app

EXPOSE 3000

CMD ["yarn", "dev"]