<script>
    import { view } from "$lib/stores";
    import Footer from "$ui/layout/Footer.svelte"
    import Header from "$ui/layout/Header.svelte";
</script>

<div class="min-h-screen bg-gray-200">
    <Header />

    <main class="pb-4 pt-8">
        <div class="max-w-3xl mx-auto px-4 sm:px-6 lg:max-w-full lg:px-32">
            <h1 class="sr-only">{$view.title}</h1>

            <slot />
        </div>
    </main>

    <Footer />
</div>
