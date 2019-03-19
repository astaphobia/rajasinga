FROM node:latest as base
WORKDIR /client
COPY . .
RUN yarn install

FROM base as prod
CMD ["yarn", "build"]
ENTRYPOINT ["yarn", "start"]

FROM base as dev
ENTRYPOINT ["yarn", "start"]
