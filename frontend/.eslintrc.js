module.exports = {
    parser: "babel-eslint",
    env: {
        browser: true,
        es6: true,
        node: true,
    },
    extends: ["airbnb"],
    globals: {
        Atomics: "readonly",
        SharedArrayBuffer: "readonly",
    },
    parserOptions: {
        ecmaFeatures: {
            jsx: true,
        },
        ecmaVersion: 2018,
        sourceType: "module",
    },
    plugins: ["react"],
    rules: {
        "react/jsx-filename-extension": [1, { extensions: [".js", ".jsx"] }],
        "import/no-unresolved": 0,
        "react/jsx-one-expression-per-line": 0,
        "react/jsx-fragments": 0,
        "import/prefer-default-export": 0,
        "jsx-a11y/anchor-is-valid": [
            "error",
            {
                components: ["Link"],
                specialLink: ["hrefLeft", "hrefRight"],
                aspects: ["invalidHref", "preferButton"],
            },
        ],
    },
    settings: { "import/resolver": { node: { paths: ["src"] } } },
};
