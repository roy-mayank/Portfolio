<script lang="ts">
		import plane from '$lib/images/plane.png';
</script>

<svelte:head>
	<title>Home</title>
	<meta name="description" content="Svelte demo app" />
</svelte:head>

<div class="group relative flex flex-col items-center justify-end h-screen overflow-hidden cursor-crosshair text-white">
    <img 
        src={plane} 
        alt="plane" 
        class="absolute bottom-0 w-84 opacity-100 scale-75 translate-y-24 transition-all duration-1000 ease-out group-hover:opacity-100 group-hover:-translate-y-12 group-hover:scale-120 pointer-events-none"
    />

    <div class="flex flex-col-reverse justify-center items-start pb-20 z-10">
        <h1 class="font-bold font-jost text-8xl tracking-wider cursor-default">
            MAYANK ROY
        </h1>
        <p class="text-xl font-semibold font-jost ml-1 mb-2 oldstyle-nums">Maximizing shareholder value since 2003</p>
    </div>
</div>


<style>


</style>
