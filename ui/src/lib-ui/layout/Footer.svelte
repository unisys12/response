<script lang="ts">
    import { browser } from '$app/env'
</script>

<footer>
    <div class="max-w-3xl mx-auto px-4 sm:px-6 lg:px-32 lg:max-w-full">
        <div class="py-8 text-sm text-gray-500 text-center sm:text-left flex md:flex-row md:justify-between">
            <div>
                <span class="block sm:inline">&copy; {(new Date).getFullYear()} Contaim, LLC and <a target="_blank" href="https://github.com/responserms/response/graphs/contributors" class="font-medium text-indigo-600 hover:text-indigo-500">Contributors</a>. All rights reserved.</span>
            </div>

            <div>
                <span class="block sm:inline">Powered by <a href="https://responserms.com" class="font-medium text-indigo-600 hover:text-indigo-500" target="_blank">Response {browser ? window.ResponseMeta.version.replace('{{ .Meta.Version }}', '') : 'Loading...'}</a> &middot; <a href="https://docs.responserms.com/users" target="_blank" class="font-medium text-indigo-600 hover:text-indigo-500">Read the Docs</a> &middot; <a href="https://www.mozilla.org/en-US/MPL/2.0/" target="_blank" class="font-medium text-indigo-600 hover:text-indigo-500">Open Source</a></span>
            </div>
        </div>
    </div>
</footer>