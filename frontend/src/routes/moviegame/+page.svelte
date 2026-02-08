<!-- Imported straight from older version of portfolio with svelte alterations -->

<script lang="ts">
    import axios from 'axios';
    import tmdb from '$lib/images/tmdb.png';
    
    let submitted = false;
    let movies = ['', '', ''];
    let recommendList: any[] = [];
    let results: number[] = [];
    let recommendcount = 0;
    
    let errors = {
        movie1: '',
        movie2: '',
        movie3: ''
    };

    const displayRatings = async (names: string[]) => {
        const replies = names.map((name) =>
            axios.get(
                `https://api.themoviedb.org/3/search/movie?api_key=3bf98afbe95c4d4b4fd292ed2e9eedab&query=${name}`
            )
        );

        const responses = await Promise.all(replies);

        const movieIds = responses
            .map((response) => response.data.results[0]?.id)
            .filter((id) => id);

        const movieReplies = movieIds.map((id) =>
            axios.get(
                `https://api.themoviedb.org/3/movie/${id}?api_key=3bf98afbe95c4d4b4fd292ed2e9eedab`
            )
        );

        const movieResponses = await Promise.all(movieReplies);

        const voteAverages = movieResponses.map(
            (movieResponse) => movieResponse.data.vote_average
        );

        const allVoteAverage =
            voteAverages.reduce((acc, curr) => acc + curr, 0) / voteAverages.length;
        results = [allVoteAverage];

        const recommendationReplies = movieIds.map((result) =>
            axios.get(
                `https://api.themoviedb.org/3/movie/${result}/recommendations?api_key=3bf98afbe95c4d4b4fd292ed2e9eedab`
            )
        );

        const recommendationResponse = await Promise.all(recommendationReplies);

        const recommendations = recommendationResponse.map(
            (recommendationResponse) =>
                recommendationResponse.data.results.map((movie: any) => ({
                    title: movie.title,
                    posterUrl: `https://image.tmdb.org/t/p/w500/${movie.poster_path}`,
                }))
        );
        recommendList = recommendations;
    };

    const handleMovies = (event: Event, index: number) => {
        const target = event.target as HTMLInputElement;
        movies[index] = target.value;
    };

    const validateForm = () => {
        let isValid = true;
        errors = { movie1: '', movie2: '', movie3: '' };

        if (!movies[0]) {
            errors.movie1 = 'This is required!';
            isValid = false;
        }
        if (!movies[1]) {
            errors.movie2 = 'This is required!';
            isValid = false;
        }
        if (!movies[2]) {
            errors.movie3 = 'This is required!';
            isValid = false;
        }

        return isValid;
    };

    const showRatings = async (event: Event) => {
        event.preventDefault();
        
        if (!validateForm()) {
            return;
        }

        await displayRatings(movies);
        submitted = true;
        recommendcount = 0;
    };

    const scrollToBottom = () => {
        window.scrollTo({
            top: document.body.scrollHeight,
            behavior: 'smooth',
        });
    };
</script>

<svelte:head>
    <title>Movie Taste Game</title>
</svelte:head>

<div class="min-h-screen bg-fixed font-jost text-white align-middle items-center flex flex-col justify-center">
    <div class="flex flex-col movie-body px-2 mx-28 gap-6">
        <div class="movie-headers flex flex-col gap-4">
            <h1 class="text-4xl text-back font-semibold">
                Legacy movie taste game
            </h1>
        </div>
        <div class="movie-deets flex flex-col gap-2">
            <p>Welcome to a simple movie taste game! (imported straight from a previous version of the portfolio website)</p>
            <p>
                How this works is: You pick 3 tv/movies that come right
                outta yer head and we calculate a score for ye based on its
                popularity and ratings.
            </p>
            <p> Powered by Axios rest API calls to TMDB </p>
        </div>
        <div class="movie-form flex flex-col mx-4">
            <form class="flex flex-col gap-6 text-gray-600" on:submit={showRatings}>
                <div class="input-field flex flex-col">
                    <input
                        type="text"
                        placeholder="Enter Movie/Series one:"
                        on:input={(event) => handleMovies(event, 0)}
                        class="rounded-3xl ring-offset-1 ring-2 p-1 px-3 drop-shadow-md bg-[#e8f0fe]"
                    />
                    {#if errors.movie1}
                        <span class="text-red-500">{errors.movie1}</span>
                    {/if}
                </div>
                <div class="input-field flex flex-col">
                    <input
                        type="text"
                        placeholder="Enter Movie/Series two:"
                        on:input={(event) => handleMovies(event, 1)}
                        class="rounded-3xl ring-offset-1 ring-2 p-1 px-3 drop-shadow-md bg-[#e8f0fe]"
                    />
                    {#if errors.movie2}
                        <span class="text-red-500">{errors.movie2}</span>
                    {/if}
                </div>
                <div class="input-field flex flex-col">
                    <input
                        type="text"
                        placeholder="Enter Movie/Series three:"
                        on:input={(event) => handleMovies(event, 2)}
                        class="rounded-3xl ring-offset-1 ring-2 p-1 px-3 drop-shadow-md bg-[#e8f0fe]"
                    />
                    {#if errors.movie3}
                        <span class="text-red-500">{errors.movie3}</span>
                    {/if}
                </div>
                <div class="flex flex-col submit-box bg-toppage m-auto p-3 rounded-xl mt-3 drop-shadow-lg">
                    <button type="submit" class="font-extrabold text-white text-lg cursor-pointer hover:scale-105 transition-transform">
                        Submit
                    </button>
                </div>
            </form>
        </div>
        
        {#if submitted && results.length > 0}
            <div class="movie-ratings">
                <div class="flex flex-col gap-4 p-6 bg-white/50 rounded-xl">
                    <h2 class="text-2xl font-bold">Your Taste Score</h2>
                    <p class="text-4xl font-bold text-blue-600">{results[0].toFixed(1)}/10</p>
                </div>
            </div>
        {/if}
        
        
        {#if submitted}
            <div class="flex flex-col gap-4 my-10">
                <h1 class="text-4xl text-back font-semibold">
                    RECOMMENDATIONS FROM YOUR PICKS
                </h1>
                <h2 class="tracking-wider font-bold">
                    Picked randomly from API (don't @ me)
                </h2>
                <div class="recommended-movies flex flex-row gap-16 flex-wrap">
                    {#each recommendList as recommendmovie, i}
                        {#if recommendmovie.length !== 0}
                            {@const randomIndex = Math.floor(Math.random() * recommendmovie.length)}
                            {@const recommendedMovie = recommendmovie[randomIndex]}
                            {#if recommendedMovie}
                                <div class="flex flex-col gap-2 items-center">
                                    <img 
                                        src={recommendedMovie.posterUrl} 
                                        alt={recommendedMovie.title}
                                        class="rounded-lg shadow-lg w-48"
                                    />
                                    <p class="text-center font-semibold">{recommendedMovie.title}</p>
                                </div>
                            {/if}
                        {/if}
                    {/each}
                    {#if recommendList.every(movie => movie.length === 0)}
                        <div class="flex flex-col gap-2">
                            <h2 class="tracking-wider">
                                Unfortunately, no recommendations :(
                            </h2>
                            <h2 class="tracking-wider">Try again!</h2>
                        </div>
                    {/if}
                </div>
            </div>
        {/if}
    </div>
</div>

<div class="fixed right-2 bottom-2 flex flex-row items-end gap-2">
    <p class="align-bottom text-white">Powered By:</p>
    <a href="https://www.themoviedb.org/">
        <img class="cursor-pointer" src={tmdb} alt="tmdb" width={100} />
    </a>
</div>

<style>
</style>