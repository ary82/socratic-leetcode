<script lang="ts">
	import { onMount } from 'svelte';
	import { PUBLIC_BACKEND_URL } from '$env/static/public';

	let content: string = '';

	let loc: number = 0;
	let lang: string = '';
	let question: string = '';
	let code: string = '';

	let insights: any = {};

	$: loc = code.split(/\r\n|\r|\n/).length;

	const update = () => {
		chrome.runtime.sendMessage(
			{
				type: 'get-data'
			},
			(res) => {
				if (res) {
					lang = res.langText;
					content = res.problemText;
					code = res.codeText;
					question = res.questionText;
				}
			}
		);
	};

	const getInsights = async () => {
		const res = await fetch(`${PUBLIC_BACKEND_URL}/generate`, {
			method: 'POST',
			headers: {
				// 'Access-Control-Allow-Origin': '*',
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

<h1>Socratic Code</h1>
<p>make sure only one leetcode tab is open for best performance</p>
<button on:click={update}>Update question and input</button>
<button on:click={getInsights}>Generate Questions</button>
<h2>{question} - {lang} - LOC: {loc}</h2>
<div>{content}</div>
<div>{code}</div>
<div>{insights}</div>
