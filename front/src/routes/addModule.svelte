<script>
	async function CreateVM() {
		const res = await fetch('http://localhost:8080/vm', {
			method: 'POST'
		});
	}
</script>

<div class="flex justify-center p-4">
	<div class="flex flex-row items-center">
		<button class="border border-gray-400 p-4" on:click={CreateVM}>Create</button>
		<div class="flex justify-center font-bold text-xl p-5 w-full">VM</div>
		<a class="font-extralight underline mr-3" href="/">Add</a>
	</div>
</div>
