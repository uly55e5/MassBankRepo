<script>
    export let page = 1
    export let totalItems
    export let pagesize = 10

    $: offset=(page-1)*pagesize
    $: totalpages=Math.ceil(totalItems/pagesize)
    function changePageRelative(n) {
        console.log(n)
        page = page + n > totalpages ? totalpages : page + n < 1 ? 1 : page + n
    }
</script>

<div>
    <span on:click="{() => changePageRelative(-10)}" disabled={page<3 || null}>&#8810;</span>
    <span on:click="{() => changePageRelative(-1)}" disabled={page<2 || null}>&#60;</span>
    {page}
    <span on:click="{() => changePageRelative(1)}" disabled={page>totalpages-1 || null}>&#62;</span>
    <span on:click="{() => changePageRelative(10)}" disabled={page>totalpages-2 || null}>&#8811;</span>
    <span >({totalpages})</span>
</div>

<style>
    span[disabled] {
        color: lightgray;
    }
</style>