<script lang="ts">
    import { onMount } from 'svelte';
    import axios from 'axios';

    let docText = "";
    let socket: WebSocket;
    let activeUserCount = 0;

    onMount(() => {
        socket = new WebSocket("ws://localhost:8080/ws");
        socket.onmessage = (event) => {
            docText = event.data;
        };
        const userJWT = localStorage.getItem("userJWT") || "";

        const interval = setInterval(async () => {
            await axios.post("http://localhost:8080/heartbeat", {}, {
                headers: { Authorization: `Bearer ${userJWT}` }
            });
        }, 10000);

        return () => {
            socket.close();
            clearInterval(interval);
        };
    });

    function handleTyping(e) {
        socket.send(e.target.value);
    }
</script>

<div class="p-10 bg-slate-900 text-white">
    <h2>Active Users: {activeUserCount}</h2>
    
    <textarea 
        class="w-full h-64 bg-slate-800 p-4 rounded"
        value={docText}
        on:input={handleTyping}
        placeholder="Type here to collaborate..."
    ></textarea>
</div>