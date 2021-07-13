module.exports = {
  "parser": "@typescript-eslint/parser",
  "plugins": ["@typescript-eslint"],
  "extends": [
    //"airbnb-typescript",
    "airbnb/hooks",
    "plugin:@typescript-eslint/recommended",
    "plugin:@typescript-eslint/recommended-requiring-type-checking"
  ],
  "parserOptions": {
    "project": "./tsconfig.json",
    "tsconfigRootDir": __dirname
  },
  "env": {
    "browser": true
  },
  "ignorePatterns": [
    "build",
    "public",
    "src/env-vars.ts",
    "src/index.jsx" // TODO: convert to tsx
  ],
  "rules": {
    // "export default" is explicitly avoided because it creates refactoring problems if the filename changes.
    // default exports use the name of filename to name the object. If the filename is changed _all_ its uses _also_ have to change.
    "react/prefer-stateless-function": 0,
    // Many components are sufficiently complex that they should benefit from react components props/state management,
    // even if they could be represented as pure stateless components.
    "jsx-a11y/interactive-supports-focus": 0,
    // Overkill. Prevents "interactive roles" from being able to trigger events, in the particular instance a div from being clickable.
    "max-len": [2, {"code": 200, "tabWidth": 2, "ignoreUrls": true}],
    // 100 lines is too restrictive

    // Fix rules:
    "@typescript-eslint/no-unsafe-call": 0,
    "@typescript-eslint/no-unsafe-member-access": 0,
    "@typescript-eslint/no-unsafe-assignment": 0,
    "@typescript-eslint/no-explicit-any": 0,
    "@typescript-eslint/no-floating-promises": 0,
    "jsx-a11y/click-events-have-key-events": 0,
    "@typescript-eslint/unbound-method": 0,

    // Added rules:
    "@typescript-eslint/indent": 0,
    "@typescript-eslint/no-unused-vars": 0, // Seemed to be incorrect a lot of the time :-/
    "@typescript-eslint/explicit-module-boundary-types": 0,
    "@typescript-eslint/no-unsafe-return": 0,
  }
};

// Iffy Rules:
// "prefer-object-spread": 0,
// I may come to regret this one. Rule wants to convert Object.assign to object spreads: {...obj1, ...obj2}.