<script>
    import { onMount } from "svelte";
    import Pagination from "./pagination.svelte"
    import {apiUrl} from "../js/config";"../js/config"
    const endpoint = $apiUrl+"/spectrainfo";
    let data = {};
    let pagesize = 10
    let page = 1
    $: totaldata = data.totalCount ? data.totalCount : 0
    $: spectra = data.spectra ? data.spectra : []
    onMount(async function () {
        reloadData()
    });

    async function reloadData() {
        const url = endpoint+"?limit="+pagesize+"&page="+page
        const response = await fetch(url);
        data = await response.json();
        console.log(data);
    }
    $: page && reloadData()
</script>

{#each spectra as mb}
    <p>{mb.accession}</p>
{/each}
<Pagination totalItems={totaldata} bind:page></Pagination>
