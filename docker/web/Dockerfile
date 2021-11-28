FROM node:lts-alpine3.14 AS builder
RUN apk update && apk upgrade && apk add --no-cache bash git openssh
WORKDIR /app
COPY ../../web /app
#RUN yarn set version stable
RUN npm install
RUN npm run build
CMD ["npm", "run", "start"]
EXPOSE 3000
