<script context="module">
    import { setEnvironment } from '$houdini';
    import env from '../environment';

    setEnvironment(env);
</script>

<script lang="ts">
    import 'leaflet/dist/leaflet.css'
    import '@fortawesome/fontawesome-free/js/brands.min'
    import '@fortawesome/fontawesome-free/js/solid.min'
    import '@fortawesome/fontawesome-free/js/fontawesome.min'
    import "../app.postcss";
    import { updateSettings, view } from "$lib/stores";
    import nightwind from "nightwind/helper"
    import { graphql, query, GetConsoleSettings } from "$houdini";

    nightwind.init()

    // const { data } = query<GetConsoleSettings>(graphql`
    //     query GetConsoleSettings {
    //         consoleSettings
    //     }
    // `)
    //
    // $: {
    //     updateSettings($data?.consoleSettings)
    // }
</script>

<svelte:head>
    <title>{$view.title} - Response</title>
    <meta name="description" content={$view.description}>
</svelte:head>

<div class="antialiased max-h-screen min-h-screen">
    <slot></slot>
</div>