<script>

    export let data = {}
    export let level = 0

    $: indent = "&nbsp;".repeat(level*4)
</script>

{#each Object.entries(data) as [key,val],index}
    {#if val && typeof(val)==='object'}
        {#if key != index}
            <div class="text-left">{@html indent}{key}:</div>
        {/if}
    {:else}
        {#if val}
            <div class="text-left">{@html indent}{key}: {val}</div>
        {/if}
    {/if}
{/each}

<style>
    .text-left {
        text-align: left;
    }
</style>