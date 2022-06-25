FROM node:latest

WORKDIR /app

COPY package.json ./

RUN npm install

COPY . .

CMD ["npm", "start"]


# pull official base image
# FROM node:16-alpine

# set working directory
# WORKDIR /app

# RUN chown -R node /app/node_modules

# # add `/app/node_modules/.bin` to $PATH
# ENV PATH /app/node_modules/.bin:$PATH

# # install app dependencies
# COPY ./package.json ./
# COPY ./package-lock.json ./
# RUN npm install
# # RUN npm install react-scripts@3.4.1 -g

# # add app
# COPY ./ ./

# # start app
# CMD ["npm", "start"]