<script lang="ts">
    import { scale } from 'svelte/transition'
    import { cubicIn, cubicOut } from "svelte/easing";
    import { clickOutside } from "$lib/actions";

    let open: boolean = false

    function handleClick() {
        open = !open
    }
</script>

<div class="relative inline-block text-left nightwind-prevent-block" use:clickOutside={{ enabled: open, cb: () => open = false }}>
    <div on:click={handleClick}>
        <slot name="trigger" />
    </div>

    {#if open}
    <div
        class="origin-top-right z-50 absolute right-0 mt-2 w-56 rounded-md shadow-lg bg-white ring-1 ring-black ring-opacity-5 divide-y divide-gray-100 focus:outline-none"
        role="menu"
        aria-orientation="vertical"
        aria-labelledby="menu-button"
        tabindex="-1"
        in:scale={{ duration: 100, start: 0.95, easing: cubicOut }}
        out:scale={{ duration: 75, start: 1, easing: cubicIn }}
    >
        <slot />
    </div>
    {/if}
</div>
