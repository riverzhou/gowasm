'use strict';

const WASM_URL = '/lib/cube.wasm';
let wasm;

function init() {
	if (!WebAssembly.instantiateStreaming) { // polyfill
		WebAssembly.instantiateStreaming = async (resp, importObject) => {
			const source = await (await resp).arrayBuffer();
			return await WebAssembly.instantiate(source, importObject);
		};
	}
	const go = new Go();
	let mod, inst;
	WebAssembly.instantiateStreaming(fetch(WASM_URL), go.importObject).then((result) => {
		mod = result.module;
		inst = result.instance;
		run().then((result) => {
			console.log("Ran WASM: ", result)
		}, (failure) => {
			console.log("Failed to run WASM: ", failure)
		})
	});
	async function run() {
		await go.run(inst);
		inst = await WebAssembly.instantiate(mod, go.importObject); // reset instance
	}
	console.log(go.importObject);
	console.log('wasm loaded');
}

//function draw() {
//    wasm.exports.draw();
//}

init();

