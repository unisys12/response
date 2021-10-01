<script lang="ts">
    export let href: string = null
</script>

<a {href} on:click class="text-gray-700 hover:text-gray-900 hover:bg-gray-100 block px-4 py-2 text-sm cursor-pointer" role="menuitem" tabindex="-1" on:click on:dblclick>
    <slot />
</a>