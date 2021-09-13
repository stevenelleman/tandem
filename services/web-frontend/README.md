# Web Frontend Service 

## Table of Contents 
- [Local Service Setup](#local-service-setup)
- [Docker Setup](#docker-setup)
- [Available Scripts](#available-scripts)
    - [depcheck](#depcheck)
- [Examples](#examples)
- [Original README](#original-readme)

## Setup 
1. `yarn start`
2. Navigate to `http://localhost:3000/`

## Docker Setup 
1. `docker build -t web-frontend .`
2. `docker run -p 80:80 web-frontend` (running over HTTP, `443:443` for HTTPS)
3. Run [`Public-API` service](./public-api) (`make run` in `./public-api`) 
3. Open `http://localhost`
        
How will user-data be organized? Also should there be a more comprehensive user action history? In other words, a completely private history of metadata. It would be pretty cool if this was default scope. What would a user history scope look like? Also it would likely need to be several scopes: one for each FI and then in combination. This seems like a fascinating application of scopes - can the same sources (log entries) be tied together into 

## Available Scripts

In the project directory, you can run:

### `depcheck`

To check if dependencies are unused or missing.

### `Examples`
For examples of how redux actions/reducers are used in-code, refer to: https://github.com/stevenelleman/tandem/commit/5e6120f3eb0f93b83b148fbc05404373275f6d6a

## Original README

This project was bootstrapped with [Create React App](https://github.com/facebook/create-react-app).

### `yarn start`

Runs the app in the development mode.<br />
Open [http://localhost:3000](http://localhost:3000) to view it in the browser.

The page will reload if you make edits.<br />
You will also see any lint errors in the console.

### `yarn test`

Launches the test runner in the interactive watch mode.<br />
See the section about [running tests](https://facebook.github.io/create-react-app/docs/running-tests) for more information.

### `yarn build`

Builds the app for production to the `build` folder.<br />
It correctly bundles React in production mode and optimizes the build for the best performance.

The build is minified and the filenames include the hashes.<br />
Your app is ready to be deployed!

See the section about [deployment](https://facebook.github.io/create-react-app/docs/deployment) for more information.

### `yarn eject`

**Note: this is a one-way operation. Once you `eject`, you can’t go back!**

If you aren’t satisfied with the build tool and configuration choices, you can `eject` at any time. This command will remove the single build dependency from your project.

Instead, it will copy all the configuration files and the transitive dependencies (webpack, Babel, ESLint, etc) right into your project so you have full control over them. All of the commands except `eject` will still work, but they will point to the copied scripts so you can tweak them. At this point you’re on your own.

You don’t have to ever use `eject`. The curated feature set is suitable for small and middle deployments, and you shouldn’t feel obligated to use this feature. However we understand that this tool wouldn’t be useful if you couldn’t customize it when you are ready for it.

## Learn More

You can learn more in the [Create React App documentation](https://facebook.github.io/create-react-app/docs/getting-started).

To learn React, check out the [React documentation](https://reactjs.org/).

### Code Splitting

This section has moved here: https://facebook.github.io/create-react-app/docs/code-splitting

### Analyzing the Bundle Size

This section has moved here: https://facebook.github.io/create-react-app/docs/analyzing-the-bundle-size

### Making a Progressive Web App

This section has moved here: https://facebook.github.io/create-react-app/docs/making-a-progressive-web-app

### Advanced Configuration

This section has moved here: https://facebook.github.io/create-react-app/docs/advanced-configuration

### Deployment

This section has moved here: https://facebook.github.io/create-react-app/docs/deployment

### `yarn build` fails to minify

This section has moved here: https://facebook.github.io/create-react-app/docs/troubleshooting#npm-run-build-fails-to-minify
