<script>
	import MbList from "./components/mblist.svelte"
	import Home from "./routes/Home.svelte"
	import Spectrum from "./routes/Spectrum.svelte"
	import Router from "svelte-spa-router"

const routes = {
		'/': Home,
		'/spectrum/:accession': Spectrum
}


</script>

<main>
	<h1>MassBank Tools</h1>
	<h2><a href="#/">Home</a></h2>
	<Router {routes}></Router>
</main>

<style>
	main {
		text-align: center;
		padding: 1em;
		max-width: 240px;
		margin: 0 auto;
	}

	h1 {
		color: #ff3e00;
		text-transform: uppercase;
		font-Size: 4em;
		font-weight: 100;
	}

	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}
</style>