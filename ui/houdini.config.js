/** @type {import('houdini').ConfigFile} */
const config = {
	schemaPath: './schema.graphql',
	sourceGlob: 'src/**/*.svelte',
	module: 'esm',
	framework: 'kit',
	static: true,
	apiUrl: 'http://localhost:9080/api/graphql',
	scalars: {
		Map: {
			type: 'Object',
			unmarshal(val) {
				return val;
			},
			marshal(map) {
				return map;
			}
		}
	}
};

export default config;
