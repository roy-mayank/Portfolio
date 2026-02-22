<script lang="ts">
    import plane from '$lib/images/plane.png';
    import resume from '$lib/files/resume.pdf';
    import ProjectModal from '$lib/components/ProjectModal.svelte';
    import { projects } from '$lib/data/projects.js';
    import type { ProjectEntry } from '$lib/data/projects.js';

    let y = $state(0);
    let innerHeight = $state(0);
    let selectedProject = $state<ProjectEntry | null>(null);

    let scale = $derived(1 + (y / 100));
    let translateY = $derived(-(y * 1.5));
    let opacity = $derived(innerHeight ? Math.min((y / innerHeight) * 3, 1) : 0);
</script>

<!-- Binding the window scroll to variable y -->
<svelte:window bind:scrollY={y} bind:innerHeight={innerHeight} /> 

<svelte:head>
    <title>Mayank Roy's Portfolio</title>
    <meta name="description" content="Mayank's Portfolio site" />
</svelte:head>

<div class="relative flex flex-col items-center justify-center h-screen overflow-hidden text-white px-10">
    
    <img 
        src={plane} 
        alt="plane" 
        style:transform="scale({scale}) translateY({translateY}px)"
        style:opacity={opacity}
        class="absolute bottom-0 w-96 pointer-events-none z-0 will-change-transform"
    />

    <div class="flex flex-col-reverse justify-center items-start pb-20 z-10">
        <h1 class="font-bold font-jost text-6xl md:text-8xl xl:text-9xl tracking-tight cursor-default">
            MAYANK ROY
        </h1>
        <p class="text-xl font-semibold font-jost ml-1 mb-2 oldstyle-nums tracking-tighter">
            Maximizing shareholder value since 2003
        </p>
    </div>
</div>
<div class="font-jost flex flex-col items-center w-full px-6 py-16 md:py-24">
    <div class="max-w-3xl w-full space-y-8 text-center md:text-left bg-white/5 border border-white/10 p-8 md:p-12 rounded-3xl backdrop-blur-sm">
        <p class="text-lg md:text-xl text-white/90 leading-relaxed font-light">
            I am an <span class="font-semibold">Indian-Origin student</span> pursuing my <span class="font-semibold">Master's in Computer and Information Science </span> degree at the <span class="text-white font-semibold">University of Pennsylvania's</span> School of Engineering and Applied Sciences.
        </p>
        <p class="text-lg md:text-xl text-white/80 leading-relaxed font-light">
            I place significant value in critical expansion of my knowledge base, possibly a knock-on effect of being from a family of academically motivated individuals.
        </p>
        <p class="text-lg md:text-xl text-white/80 leading-relaxed font-light">
            Additionally, I bring the ability to <span class="font-semibold">objectively and pragmatically assess situations</span>, while also retaining <span class="font-semibold">deep empathy and curiosity</span> in my thought.
        </p>
        
    </div>
</div>
<div class="font-jost flex flex-col text-white items-center">
    <h2 class="uppercase text-sm opacity-60 mb-4">About Me</h2>
    <p class="text-4xl md:text-5xl font-bold mb-16">My Interests</p>

    <div class="grid grid-cols-1 md:grid-cols-3 gap-8 w-full">
        
        <div class="group flex flex-col items-center justify-center p-10 rounded-2xl bg-white/5 border border-white/10 hover:bg-white/10 transition-all duration-300">
            <span class="material-symbols-outlined text-5xl mb-6 text-blue-400 group-hover:scale-110 transition-transform">code</span>
            <p class="text-lg font-medium tracking-tight">Systems & Software Engineering</p>
        </div>

        <div class="group flex flex-col items-center justify-center p-10 rounded-2xl bg-white/5 border border-white/10 hover:bg-white/10 transition-all duration-300">
            <span class="material-symbols-outlined text-5xl mb-6 text-emerald-400 group-hover:scale-110 transition-transform">map</span>
            <p class="text-lg font-medium tracking-tight">Maps & Geography</p>
        </div>

        <div class="group flex flex-col items-center justify-center p-10 rounded-2xl bg-white/5 border border-white/10 hover:bg-white/10 transition-all duration-300">
            <span class="material-symbols-outlined text-5xl mb-6 text-amber-400 group-hover:scale-110 transition-transform">flight_takeoff</span>
            <p class="text-lg font-medium tracking-tight">Commercial Aviation & Travel</p>
        </div>

    </div>
</div>
<div class="font-jost flex flex-col text-white items-center py-16 w-2/3 ">
    <p class="text-4xl md:text-5xl font-bold mb-16">My Projects</p>
    <div class="grid grid-cols-1 gap-8 w-full">
        {#each projects as project (project.id)}
            <button
                type="button"
                class="group flex flex-col items-center justify-center p-10 rounded-2xl bg-white/5 border border-white/10 hover:bg-white/10 transition-all duration-300 text-left w-full"
                onclick={() => selectedProject = project}
            >
                <p class="text-lg font-medium tracking-tight text-white">{project.name}</p>
                <p class="text-white/80 text-sm mt-2 text-center">{project.shortDescription}</p>
            </button>
        {/each}
    </div>
</div>
<div class="font-jost flex flex-col text-white items-center w-full aspect-video">
    <p class="text-4xl font-bold opacity-80">Latest Resume</p>
    <p class="text-sm text-gray-400 mb-6 font-bold"> Last Updated: February 7, 2026</p>

    <iframe 
        src={resume} 
        class="w-full max-w-4xl h-full rounded-lg border border-white/0 mb-20"
        title="Resume">
    </iframe>
</div>

<ProjectModal project={selectedProject} onClose={() => selectedProject = null} />

<style>

</style>