<script>
	export let Title;
	export let Description;
	export let Creator;
	export let DateAdd;
	export let IdModule;

	function navigate() {}
</script>

<div class="flex justify-center w-full lg:max-w-full lg:flex">
	<div
		class="h-48 lg:h-auto flex-none bg-cover rounded-t lg:rounded-t-none lg:rounded-l text-center overflow-hidden"
		style="background-image: url('/mountain.jpg')"
		title="Mountain"
	/>
	<div
		class="border-r border-b border-l border-gray-400   lg:border-t lg:border-gray-400 bg-white rounded-b lg:rounded-b-none lg:rounded-r p-4 flex flex-col justify-between leading-normal"
	>
		<div class="mb-8">
			<div class="text-gray-900 font-bold text-xl mb-2 flex flex-row items-center">
				<div>{Title}</div>
				<a class="ml-4 text-xs underline" href="ressource/{IdModule}">Détails</a>
			</div>
			<p class="text-gray-700 text-base">{Description}</p>
		</div>
		<div class="flex items-center">
			<img class="w-8 h-8 rounded-full mr-4" src="http://i.pravatar.cc/300" alt="Avatar of User" />
			<div class="text-sm">
				<p class="text-gray-900 leading-none">{Creator}</p>
				<p class="text-gray-600">{DateAdd}</p>
			</div>
		</div>
	</div>
</div>
