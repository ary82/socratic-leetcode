<script lang="ts">
	import Github from '$lib/icons/github.svelte';
	import Linkedin from '$lib/icons/linkedin.svelte';
	import Mail from '$lib/icons/mail.svelte';
	import '../app.css';
</script>

<slot />

<footer
	class="flex footer bg-neutral text-neutral-content items-center p-4 justify-between text-xs"
>
	<aside class="grid-flow-col items-center">
		<p>Copyright © {new Date().getFullYear()} - Aryan Goyal</p>
	</aside>
	<nav class="grid-flow-col gap-4 md:place-self-center md:justify-self-end items-end">
		<a href="https://github.com/ary82" target="_blank" rel="noopener noreferrer"><Github /></a>
		<a href="https://linkedin.com/in/aryan-goyal1" target="_blank" rel="noopener noreferrer"
			><Linkedin /></a
		>
		<a href="mailto:mail@ary82.dev"><Mail /></a>
	</nav>
</footer>
