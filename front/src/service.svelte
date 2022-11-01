<!-- <script>
	export async function getModuleById(id: number): Promise<IModule> {
		var module: IModule;

		await fetch(variables.base_url + 'module/detail/' + id)
			.then((response) => response.json())

			.then(async (data) => {
				module = data.data;
			})
			.catch(async (error) => {
				console.log(error);
			});

		return await new Promise((resolve) => resolve(module));
	}
</script> -->
