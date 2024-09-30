<script lang="ts">
	const getCode = async () => {
		const [tab] = await chrome.tabs.query({ active: true, lastFocusedWindow: true });
		console.log(tab.id);

		chrome.scripting.executeScript({
			target: { tabId: tab.id },
			world: 'MAIN',
			func: (url: string | undefined) => {
				const lg = window.monaco?.editor?.getModels()[0].getLanguageId();
				const code = window.monaco?.editor?.getModels()[0].getValue();
				console.log(url);
				console.log(lg, code);
			},
			args: [tab.url]
		});

		chrome.runtime.sendMessage(
			{ method: 'getLocalStorage', key: 'lc-curr-code' },
			function (response) {
				console.log(response.data);
			}
		);
	};
</script>

<h1>Socratic Code</h1>
<button on:click={getCode}>Get code</button>
