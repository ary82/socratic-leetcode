<script lang="ts">
	import { PUBLIC_BACKEND_URL } from '$lib';
	import { onMount } from 'svelte';

	let content: string = '';

	let loc: number = 0;
	let lang: string = '';
	let question: string = '';
	let code: string = '';

	let codeHTML: string;
	let container: Element;

	let insights: { line: number; q: string }[] = [];

	$: loc = code.split(/\r\n|\r|\n/).length;
	$: if (codeHTML !== null) {
		if (codeHTML && container !== undefined) {
			container.innerHTML = '';
			container.innerHTML = codeHTML;
		}
	}
	$: if (insights !== null) {
		insights.forEach((insight: { line: number; q: string }) => {
			container?.children[insight.line]?.classList.add(
				'bg-info',
				'text-info-content'
				// 'tooltip'
			);
			// container?.children[insight.line]?.setAttribute('data-tip', insight.q);
		});
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
					console.log(codeHTML);
				}
			}
		);
	};

	const getInsights = async () => {
		const children = container.children;
		for (let index = 0; index < children.length; index++) {
			const child = children[index];
			child.className = '';
		}
		insights = [];
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

	<h2 class="text-lg font-bold text-center">URL: {question} -- LANG: {lang} -- LOC: {loc}</h2>
	<div class="flex gap-2 p-1 justify-center">
		<button class="btn p-2" on:click={update}>Manually update question or input</button>
		<button class="btn btn-primary p-2" on:click={getInsights}>Teach me!</button>
	</div>
	{#if codeHTML !== '' && codeHTML !== undefined}
		<h2 class="text-lg font-extrabold text-center m-2">Your Code</h2>
		<div class="mockup-code p-2 m-2">
			<code style="white-space: pre-wrap;" bind:this={container}></code>
		</div>
	{/if}
	{#if insights.length !== 0}
		<h2 class="text-lg font-extrabold text-center m-2">Insights</h2>
		<div class="mockup-window bg-base-300">
			<div class="bg-base-200 p-2">
				{#each insights as i}
					<div class="chat chat-start">
						<div class="chat-bubble chat-bubble-primary">{i.q}</div>
					</div>
				{/each}
			</div>
		</div>
	{/if}
</div>
