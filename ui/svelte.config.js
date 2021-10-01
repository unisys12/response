import preprocess from 'svelte-preprocess';
import houdini from 'houdini-preprocess'
import adapter from '@sveltejs/adapter-static'
import path from 'path'
import {isoImport} from "vite-plugin-iso-import";

/** @type {import('@sveltejs/kit').Config} */
const config = {
    // Consult https://github.com/sveltejs/svelte-preprocess
    // for more information about preprocessors
    preprocess: [houdini(), preprocess({
        postcss: true
    })],

    kit: {
        // hydrate the <div id="svelte"> element in src/app.html
        adapter: adapter(({
            fallback: 'index.html',
            pages: 'dist',
            assets: 'dist',
        })),
        target: '#response',
        paths: {
            base: ''
        },
        vite: {
            plugins: [isoImport()],
            resolve: {
                alias: {
                    $houdini: path.resolve('.', '$houdini'),
                    $ui: path.resolve('src', 'lib-ui')
                }
            }
        }
    }
};

export default config;
