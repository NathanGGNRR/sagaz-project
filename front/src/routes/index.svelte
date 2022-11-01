<script>
	import Card from '../card.svelte';
	import { onMount } from 'svelte';

	let tab = [];
	onMount(async function () {
		tab = [
			{
				Title: 'Module 1',
				Creator: 'Charly',
				Description: '2 Vm',
				DateAdd: '20 mars 2022',
				IdModule: 1
			},
			{
				Title: 'Module 2',
				Creator: 'Nathan',
				Description: '1 Vm',
				DateAdd: '10 avril 2022',
				IdModule: 2
			},
			{
				Title: 'Module 3',
				Creator: 'Charly',
				Description: '3 Vm',
				DateAdd: '3 mai 2022',
				IdModule: 3
			},
			{
				Title: 'Module 4',
				Creator: 'Kim',
				Description: '1 Vm',
				DateAdd: '11 janvier 2022',
				IdModule: 4
			},
			{
				Title: 'Module 5',
				Creator: 'Baptiste',
				Description: '1 Vm',
				DateAdd: '21 mars 2022',
				IdModule: 5
			}
		];
	});
</script>

<div class="flex flex-col justify-center">
	<div class="flex flex-row items-center">
		<div class="flex justify-center font-bold text-xl p-5 w-full">Liste des modules</div>
		<a class="font-extralight underline mr-3" href="/addModule">Add</a>
	</div>

	<div class="grid grid-cols-4 gap-4 justify-center">
		{#each tab as item, index}
			<Card
				Title={item.Title}
				Creator={item.Creator}
				Description={item.Description}
				DateAdd={item.DateAdd}
				IdModule={item.IdModule}
			/>
		{/each}
	</div>
</div>
