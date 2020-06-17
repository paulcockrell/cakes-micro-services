FROM node:13.12.0-alpine as builder

WORKDIR /app

ENV PATH /app/node_modules/.bin:$PATH
copy package.json ./
copy yarn.lock ./
RUN npm install --silent
RUN npm install react-scripts@3.4.1 -g --silent

COPY . ./
RUN yarn build

FROM nginx:stable-alpine

COPY --from=builder /app/build /usr/share/nginx/html

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
