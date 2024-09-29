<script lang="ts">
	const getUrl = async () => {
		const [tab] = await chrome.tabs.query({ active: true, lastFocusedWindow: true });
		console.log(tab.url);
	};

	const getCode = async () => {
		const [tab] = await chrome.tabs.query({ active: true, lastFocusedWindow: true });
    console.log(tab.id)
		if (tab.id === undefined) {
		}
		chrome.scripting.executeScript({
			target: { tabId: tab.id },
			func: () => {
				if (window.monaco) {
					console.log('got it');
				} else {
					console.log('nope, cant see');
				}
			}
		});
	};
</script>

<h1>Welcome to SvelteKit</h1>
<p>Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentation</p>
<button on:click={getUrl}>Get url</button>
<button on:click={getCode}>Get code</button>
