<script lang="ts">
    export let size: ("sm" | "default" | "lg") = "default"
    export let color: ("primary" | "secondary" | "danger") = "primary"

    let colorClasses
    $: {
        if (color === "primary") {
            colorClasses = "text-white bg-indigo-600 hover:bg-indigo-700 dark:hover:text-white focus:ring-indigo-500"
        }

        if (color === "secondary") {
            colorClasses = "text-white bg-indigo-100 hover:bg-indigo-200 dark:hover:text-white focus:ring-indigo-500"
        }

        if (color === "danger") {
            colorClasses = "text-white bg-red-600 hover:bg-red-700 dark:hover:text-white focus:ring-red-500"
        }
    }

    let sizeClasses
    $: {
        if (size === "sm") {
            sizeClasses = "px-2.5 py-1.5 text-xs font-medium rounded shadow-sm"
        }

        if (size === "default") {
            sizeClasses = "px-3 py-2 text-sm leading-4 font-medium rounded-md shadow-sm"
        }

        if (size === "lg") {
            sizeClasses = "px-4 py-2 text-sm font-medium rounded-md shadow-sm"
        }
    }
</script>

<button type="button" class={`inline-flex items-center ${sizeClasses} border border-transparent focus:outline-none focus:ring-2 focus:ring-offset-2 ${colorClasses}`}>
    <slot />
</button>