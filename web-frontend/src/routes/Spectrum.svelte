<script>
    import {onMount} from "svelte";
    import {apiUrl} from "../js/config";
    import SpectrumRow from "../components/SpectrumRow.svelte";
    export let params={}
    $: endpoint = $apiUrl+"/spectra/"+params.accession
    let data={}

    onMount(async function () {
        reloadData()
    });

    async function reloadData() {
        const url = endpoint
        const response = await fetch(url);
        data = await response.json();
        console.log(data);
    }
    $: params && reloadData()

</script>

<h2>Edit Massbank</h2>
<SpectrumRow data={data}></SpectrumRow>


