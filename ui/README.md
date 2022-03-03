# ISTS 2022 DigiDrop UI

## Requirements

- node.js (`>=17.0.1`)
- npm (`>=8.3.0`)

## `.env` File

You need a `.env` file to work with both the development and production builds of the UI. Here is a template for the `.env` file:

```shell
# .env
REACT_APP_API_URI=<http(s)://uri-of-backend>/api/query
```

## Dev Server

To spin up the dev server, first you need to install dependencies:

```shell
$ npm install
```

Then you can start the dev server with:

```shell
$ npm start
```

## Production Build

To generate a production build, run:

```shell
$ npm run build
```

The statically compiled site will end up in the `build` directory.
