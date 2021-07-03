# Typescript Packages 

## How to Make Package? 
1. `yarn init -y`
2. `echo "node_modules\n/lib" >> .gitignore`
3. `yarn add typescript tsc`
4. Make `tsconfig.json`: 
    ``` 
    {
      "compilerOptions": {
        "target": "es5",
        "module": "commonjs",
        "declaration": true,
        "outDir": "./lib",
        "strict": true
      },
      "include": ["src"],
      "exclude": ["node_modules", "**/__tests__/*"]
    }
    ```
5. Add `"build": "tsc"` to `package.json` `scripts`. 
6. Run `yarn run build` to verify build. 
7. For any local libraries use [local path notation](https://docs.npmjs.com/cli/v7/configuring-npm/package-json#local-paths) in `package.json` dependencies. 

## Rationale
Share typescript between projects. Use libraries as a kind of package incubator, as they find more users spin off into separate libraries. As such, treat packages like any other packages, do not differentiate them from other libraries. 

# TODO:
- [ ] Automatic installer: 
    - [ ] Package edit triggers `yarn install --force` in downstream projects 
- [ ] Link `package.json` `repository` to shared tandem repo? Link to internal directory? Is that possible? 
- [ ] Link `package.json` `homepage` to tandem 
- [ ] Add jest testing to all components, follow itnext.io guide

## Resources: 
- https://itnext.io/step-by-step-building-and-publishing-an-npm-typescript-package-44fe7164964c [really good]
- https://docs.npmjs.com/cli/v7/configuring-npm/package-json#local-paths [local paths in package.json]
    - https://stackoverflow.com/questions/14381898/local-dependency-in-package-json
    - Use relative pathing in `package.json` and `yarn install -S` to import the library, 
- https://github.com/npm/npm/issues/13050
- https://medium.com/@alexishevia/the-magic-behind-npm-link-d94dcb3a81af
- https://medium.com/dailyjs/how-to-use-npm-link-7375b6219557
