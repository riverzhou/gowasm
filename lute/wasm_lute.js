'use strict';

const WASM_URL = '/lib/lutewasm.wasm';
let wasm;

function init() {
    const go = new Go(); // Defined in wasm_exec.js

    WebAssembly.instantiateStreaming(fetch(WASM_URL), go.importObject).then(function (obj) {
        wasm = obj.instance;
        go.run(wasm);
    })

    console.log(go.importObject);
    console.log('wasm loaded');
}

init();

