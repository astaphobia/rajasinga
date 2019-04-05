FROM node:latest as base

WORKDIR /rajasinga/client
COPY . .
RUN yarn install
RUN yarn build

FROM base as dev 
ENTRYPOINT [ "node", "serve.js" ]

FROM base as prod
ENTRYPOINT [ "node", "serve.js" ]
