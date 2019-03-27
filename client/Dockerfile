FROM node:latest as base

WORKDIR /rajasinga/client
COPY . .
RUN yarn install

FROM base as dev 
ENTRYPOINT [ "node", "serve.js" ]

FROM base as prod
RUN yarn build
ENTRYPOINT [ "node", "serve.js" ]
