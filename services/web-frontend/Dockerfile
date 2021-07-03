# Stage 0, "build-stage", based on Node.js, to build and compile the frontend
FROM tiangolo/node-frontend:10 as build-stage
WORKDIR /app

COPY package*.json /app/

# Switched to `RUN yarn install --force` to also include local packages
RUN yarn install --force
COPY internal /app/

RUN yarn run build

# Stage 1, based on Nginx, to have only the compiled app, ready for production with Nginx
FROM nginx:1.18
COPY --from=build-stage /app/build/ /usr/share/nginx/html

# Based on the default nginx.conf provided by tiangolo/node-frontend
COPY /internal/nginx/default.conf /etc/nginx/conf.d/default.conf