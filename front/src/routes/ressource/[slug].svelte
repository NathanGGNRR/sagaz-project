<script>
	import { append } from 'svelte/internal';

	import Card from '../../card.svelte';
	import { onMount } from 'svelte';
	import { page } from '$app/stores';

	console.log($page.params);

	let tab = [];
	onMount(async function () {
		tab = [
			{
				Title: 'Ressource 1',
				Creator: 'Charly',
				Description: 'Vm',
				DateAdd: '20 mars 2022'
			},
			{ Title: 'Ressource 2', Creator: 'Charly', Description: 'test', DateAdd: '10 avril 2022' },
			{ Title: 'Ressource 3', Creator: 'Charly', Description: '3 Vm', DateAdd: '3 mai 2022' },
			{ Title: 'Ressource 4', Creator: 'Charly', Description: '1 Vm', DateAdd: '11 janvier 2022' },
			{ Title: 'Ressource 5', Creator: 'Charly', Description: '1 Vm', DateAdd: '21 mars 2022' }
		];
	});
</script>

<div class="flex flex-col justify-center">
	<div class="flex flex-row items-center">
		<div class="flex w-full justify-center font-bold text-xl p-5">Liste des ressources</div>
		<a class="font-extralight underline mr-3" href="/">back</a>
	</div>

	<div class="grid grid-cols-4 gap-4 justify-center">
		{#each tab as item, index}
			<div class="flex justify-center ">
				<div class="card border border-gray-400 p-4 ">
					<div class="card-body">
						<h1 class="card-title text-xl font-bold mr-10">{item.Title}</h1>
						<h6 class="card-subtitle mb-2 text-muted">{item.Description}</h6>
						<h6 class="card-subtitle mb-2 font-black text-muted">{item.Creator}</h6>
						<h6 class="card-subtitle text-xs mb-2 text-muted">{item.DateAdd}</h6>
					</div>
				</div>
			</div>
		{/each}
	</div>
</div>
