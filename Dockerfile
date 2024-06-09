FROM node:20.0-alpine as build

WORKDIR /usr/src/app

COPY package*.json ./

COPY tsconfig*.json ./

RUN npm install -g pnpm

RUN pnpm i

RUN pnpm build


FROM node:20.0-alpine 

WORKDIR /usr/src/app

COPY --from=build /usr/src/app/dist ./dist
COPY --from=build  /usr/src/app/node_modules ./node_modules


ENTRYPOINT [ "node" ]


CMD [ "dist/main.js" ]