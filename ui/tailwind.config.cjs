const { indigo } = require("tailwindcss/colors");

const config = {
    darkMode: "class",
    mode: "jit",
    purge: [
        "./src/**/*.{html,js,svelte,ts}",
    ],
    theme: {
        extend: {
            colors: {
                primary: indigo,
            }
        },
        nightwind: {
            typography: true,
        }
    },
    plugins: [
        require('@tailwindcss/forms'),
        require('@tailwindcss/typography'),
        require('nightwind'),
    ],
};

module.exports = config;
