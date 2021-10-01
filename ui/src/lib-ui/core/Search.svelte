<script lang="ts">
    let search: string
</script>

<label for="mobile-search" class="sr-only">Search</label>

<div class="relative text-white focus-within:text-gray-700 dark:focus-within:text-gray-700 w-full">
    <div class="pointer-events-none absolute inset-y-0 left-0 pl-3 flex items-center">
        <svg class="h-5 w-5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
            <path fill-rule="evenodd" d="M8 4a4 4 0 100 8 4 4 0 000-8zM2 8a6 6 0 1110.89 3.476l4.817 4.817a1 1 0 01-1.414 1.414l-4.816-4.816A6 6 0 012 8z" clip-rule="evenodd" />
        </svg>
    </div>

    <input id="mobile-search" bind:value={search} class="block w-full bg-white nightwind-prevent bg-opacity-10 py-2 pl-10 pr-3 border border-transparent rounded-md leading-5 text-gray-300 focus:text-gray-700 dark:focus:text-gray-700 placeholder-gray-300 focus:outline-none focus:bg-opacity-80 focus:border-transparent focus:placeholder-gray-700 focus:ring-0 sm:text-sm" placeholder="Looking for something?" type="search" name="search">
</div>