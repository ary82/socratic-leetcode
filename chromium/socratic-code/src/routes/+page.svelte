<script lang="ts">
	import { onMount } from 'svelte';
	import { PUBLIC_BACKEND_URL } from '$env/static/public';

	let content: string = '';

	let loc: number = 0;
	let lang: string = '';
	let question: string = '';
	let code: string = '';

	let codeHTML: string;
	let container: Element;

	let insights: any = {};

	$: loc = code.split(/\r\n|\r|\n/).length;
	$: if (codeHTML !== null) {
		if (codeHTML && container !== undefined) {
			container.innerHTML = '';
			container.innerHTML = codeHTML;
		}
	}
	$: if (insights !== null) {
	}

	const update = () => {
		chrome.runtime.sendMessage(
			{
				type: 'get-data'
			},
			(res) => {
				if (res) {
					lang = res.langText;
					content = res.problemText;
					question = res.questionText;
					code = res.codeText;
					codeHTML = res.codeHtml;
				}
			}
		);
	};

	const getInsights = async () => {
		const res = await fetch(`${PUBLIC_BACKEND_URL}/generate`, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({
				question,
				loc,
				lang,
				code
			})
		});
		const json = await res.json();
		if (res.ok) {
			insights = json;
			console.log(insights);
		}
	};

	onMount(() => {
		if (content === '' && code === '' && lang === '') {
			update();
		}
	});
</script>

<div class="p-4 min-w-96">
	<h1 class="text-2xl font-extrabold text-center">Socratic Code</h1>
	<p class="text-neutral-content text-center">
		**make sure only one leetcode tab is open for best performance
	</p>

	<h2 class="textarea-md font-bold text-center">URL: {question} -- LANG: {lang} -- LOC: {loc}</h2>
	<div class="flex gap-2 p-1">
		<button class="btn p-2" on:click={update}>Manually update question or input</button>
		<button class="btn btn-primary p-2" on:click={getInsights}>Teach me!</button>
	</div>
	{#if codeHTML !== '' && codeHTML !== undefined}
		<div style="white-space: pre-wrap;" class="mockup-code p-2 m-2">
			<code bind:this={container}></code>
		</div>
	{/if}
</div>
