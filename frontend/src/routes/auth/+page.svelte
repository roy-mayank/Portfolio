<script lang="ts">
    // import { auth } from '$lib/firebase';
    import { signInWithEmailAndPassword, createUserWithEmailAndPassword } from "firebase/auth";

    let isLogin = $state(true);
    let email = $state("");
    let password = $state("");
    let errorMessage = $state("");
    let loading = $state(false);

    function toggleMode() {
        isLogin = !isLogin;
        errorMessage = ""; // Clear errors when switching modes
    }

    async function handleSubmit(event: SubmitEvent) {
        event.preventDefault();
        
        // if (!auth) {
        //     errorMessage = "Firebase is not initialized";
        //     return;
        // }

        loading = true;
        errorMessage = "";

        try {
            if (isLogin) {
                // Login logic
                const userCredential = await signInWithEmailAndPassword(auth, email, password);
                console.log("Signed in:", userCredential.user);
                // Redirect or handle successful login
            } else {
                // Registration logic
                const userCredential = await createUserWithEmailAndPassword(auth, email, password);
                console.log("Signed up:", userCredential.user);
                // Redirect or handle successful registration
            }
        } catch (error: any) {
            console.error("Auth error:", error);
            errorMessage = error.message || "An error occurred";
        } finally {
            loading = false;
        }
    }
</script>

<div class="h-screen flex flex-col justify-center items-center p-4 font-jost">
    <p class="mb-6 text-2xl font-semibold tracking-tight">
        {isLogin ? 'Login' : 'Create Account'}
    </p>

    <form class="flex flex-col w-full max-w-xs" onsubmit={handleSubmit}>
        
        {#if isLogin}
            <input 
                type="email" 
                placeholder="Email Address" 
                bind:value={email}
                required
                class="border border-amber-200 p-3 rounded mb-4 outline-none focus:ring-2 focus:ring-amber-100 transition-all"
            />
        {:else}
            <input 
                type="email" 
                placeholder="Email Address" 
                bind:value={email}
                required
                class="border p-3 rounded mb-4 outline-none focus:ring-2 focus:ring-blue-100 transition-all"
            />
        {/if}

        <input 
            type="password" 
            placeholder="Password" 
            bind:value={password}
            required
            class="border p-3 rounded mb-6 outline-none focus:ring-2 focus:ring-blue-100 transition-all"
        />

        {#if errorMessage}
            <p class="text-red-500 text-sm mb-4">{errorMessage}</p>
        {/if}

        <button 
            type="submit"
            disabled={loading}
            class="bg-blue-500 text-white p-3 rounded font-semibold hover:bg-blue-600 active:scale-95 transition-all shadow-md disabled:opacity-50 disabled:cursor-not-allowed"
        >
            {loading ? 'Loading...' : (isLogin ? 'Sign In' : 'Register')}
        </button>
        
        <button 
            type="button" 
            onclick={toggleMode}
            class="mt-6 text-sm text-gray-300 hover:text-blue-500 transition-colors"
        >
            {isLogin ? "Don't have an account? Sign Up" : "Already have an account? Login"}
        </button>
    </form>
</div>