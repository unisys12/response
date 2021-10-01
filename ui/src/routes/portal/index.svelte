<script lang="ts">
    import { updateTitle } from "$lib/stores";
    import Card from "$ui/Card/Card.svelte";
    import CardHeader from "$ui/Card/CardHeader.svelte";
    import Button from "$ui/Button.svelte";
    updateTitle('Portal Home')
</script>

<!-- Main 3 column grid -->
<div class="grid grid-cols-1 gap-4 items-start lg:grid-cols-3 lg:gap-8">
    <!-- Left column -->
    <div class="grid grid-cols-1 gap-4 lg:col-span-2">
       <Card>
           <CardHeader slot="header" title="Test" description="This is a test!">
               <div slot="more">
                   <Button size="sm" color="danger">test</Button>
                   <Button size="sm">test</Button>
               </div>
           </CardHeader>

           <p>test</p>
           <p>test</p>
           <p>test</p>
           <p>test</p>
           <p>test</p>
           <p>test</p>
           <p>test</p>
           <p>test</p>
           <p>test</p>
           <p>test</p>
       </Card>
    </div>

    <!-- Right column -->
    <div class="grid grid-cols-1 gap-4">
        <section aria-labelledby="section-2-title">
            <h2 class="sr-only" id="section-2-title">Section title</h2>
            <div class="rounded-lg bg-white overflow-hidden shadow">
                <div class="p-6">
                    <!-- Your content -->
                </div>
            </div>
        </section>
    </div>
</div>