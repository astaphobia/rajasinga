FROM node:latest
WORKDIR /client
COPY . .

RUN yarn install
EXPOSE 3002
CMD ["yarn", "dev"]
