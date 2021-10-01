<script lang="ts">
    import Dropdown from "$ui/Dropdown.svelte";
    import DropdownGroup from "$ui/DropdownGroup.svelte";
    import DropdownItem from "$ui/DropdownItem.svelte";

    export let title: string = null
    export let description: string = null
</script>

<div class="bg-white px-4 py-5 border-b border-gray-200 dark:border-gray-900 sm:px-6">
    <div class="-ml-4 -mt-2 flex items-center justify-between flex-wrap sm:flex-nowrap">
        <div class="ml-4 mt-2">
            <h3 class="text-lg leading-6 font-medium text-gray-900">
                {title}
            </h3>

            {#if description}
            <p class="mt-1 text-sm text-gray-500">
                {description}
            </p>
            {/if}
        </div>

        {#if $$slots.actions}
        <div class="ml-4 mt-2 flex-shrink-0">
            <slot name="actions" />
        </div>
        {/if}

        {#if $$slots.more}
        <div class="flex-shrink-0 self-center flex z-40">
            <Dropdown>
                <button slot="trigger" class="pointer-click -m-2 p-2 rounded-full flex items-center text-gray-400 hover:text-gray-600">
                    <span class="sr-only">Open options</span>
                    <i class="fas fa-ellipsis-v"></i>
                </button>

                <DropdownGroup>
                    <DropdownItem href="/">test</DropdownItem>
                </DropdownGroup>
            </Dropdown>
        </div>
        {/if}
    </div>
</div>