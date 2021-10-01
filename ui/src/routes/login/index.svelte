<script lang="ts" context="module">
    /**
     * @type {import('@sveltejs/kit').Load}
     */
    export async function load({ page }) {
        const url = `http://${page.host}/auth`
        const res = await fetch(url)

        if (res.ok) {
            return {
                props: {
                    data: await res.json()
                }
            }
        }

        return {
            status: res.status,
            error: new Error(`Could not fetch auth providers from '${url}'`)
        }
    }
</script>

<script lang="ts">
    import { page } from '$app/stores'

    export let data: any

    let from = $page.query.get('from')
    if (from === null) {
        from = "/portal"
    }
</script>

<div class="min-h-screen bg-white flex">
    <div class="flex-1 flex flex-col justify-center py-12 px-4 sm:px-6 lg:flex-none lg:px-20 xl:px-24">
        <div class="mx-auto w-full max-w-sm lg:w-96">
            <div>
                <h2 class="mt-6 text-3xl font-extrabold text-gray-900">
                    Sign in to Response
                </h2>
                <p class="mt-2 text-sm text-gray-600">
                    <a href="#" class="font-medium text-indigo-600 hover:text-indigo-500">
                        New here? Click here to get started.
                    </a>
                </p>
            </div>

            <div class="mt-8">
                {#if data.oauth2.length > 0}
                <div>
                    <div class="space-y-2">
                        <p class="text-sm font-medium text-gray-700">
                            Sign in with
                        </p>

                        <div class={`mt-1 grid grid-cols-${data.oauth2.length < 3 ? data.oauth2.length : 3} gap-3`}>
                            {#each data.oauth2 as provider}
                                <div>
                                    <a href={`/auth/${provider.key}/login?from=${from}`} rel="external" class="w-full inline-flex justify-center py-2 px-2 border border-gray-300 rounded-md shadow-sm bg-white text-md font-medium text-gray-500 hover:bg-gray-50">
                                        <span class="sr-only">Sign in with {provider.name}</span>
                                        <i class={`fab fa-discord`}></i>
                                    </a>
                                </div>
                            {/each}
                        </div>
                    </div>

                    {#if data.password}
                    <div class="mt-6 relative">
                        <div class="absolute inset-0 flex items-center" aria-hidden="true">
                            <div class="w-full border-t border-gray-300"></div>
                        </div>
                        <div class="relative flex justify-center text-sm">
                          <span class="px-2 bg-white text-gray-500">
                            Or continue with
                          </span>
                        </div>
                    </div>
                    {/if}
                </div>
                {/if}

                {#if data.password}
                <div class="mt-6">
                    <form action="#" method="POST" class="space-y-6">
                        <div>
                            <label for="email" class="block text-sm font-medium text-gray-700">
                                Email address
                            </label>
                            <div class="mt-1">
                                <input id="email" name="email" type="email" autocomplete="email" required class="appearance-none block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm placeholder-gray-400 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm">
                            </div>
                        </div>

                        <div class="space-y-1">
                            <label for="password" class="block text-sm font-medium text-gray-700">
                                Password
                            </label>
                            <div class="mt-1">
                                <input id="password" name="password" type="password" autocomplete="current-password" required class="appearance-none block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm placeholder-gray-400 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm">
                            </div>
                        </div>

                        <div class="flex items-center justify-between">
                            <div class="flex items-center">
                                <input id="remember-me" name="remember-me" type="checkbox" class="h-4 w-4 text-indigo-600 focus:ring-indigo-500 border-gray-300 rounded">
                                <label for="remember-me" class="ml-2 block text-sm text-gray-900">
                                    Remember me
                                </label>
                            </div>

                            <div class="text-sm">
                                <a href="#" class="font-medium text-indigo-600 hover:text-indigo-500">
                                    Forgot your password?
                                </a>
                            </div>
                        </div>

                        <div>
                            <button type="submit" class="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500">
                                Sign in
                            </button>
                        </div>
                    </form>
                </div>
                {/if}
            </div>
        </div>
    </div>

    <div class="hidden lg:block relative w-0 flex-1 bg-black">
    </div>
</div>
