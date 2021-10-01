<script lang="ts">
    import { createEventDispatcher, setContext } from 'svelte'
    import L, { MapOptions, Map } from 'leaflet?client'
    import { EventBridge } from "$ui/map/event-bridge";
    import { browser } from "$app/env";

    export let options: Partial<MapOptions> = {}
    export let events = []

    setContext(L, {
        getMap: () => map,
    })

    let map: Map = null
    let eventBridge

    const dispatch = createEventDispatcher()

    function initialize(el: HTMLElement) {
        if (! browser) {
            return {
                destroy: () => null
            }
        }

        map = L.map(el, {})
        eventBridge = new EventBridge(map, dispatch, events)

        return {
            destroy: () => {
                eventBridge.unregister()
                map.remove()
            }
        }
    }

    export function getMap() {
        return map
    }
</script>

<div class="w-full h-full" use:initialize>
    {#if map}
        <slot />
    {/if}
</div>