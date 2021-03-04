'use strict';

const WASM_URL = '/lib/add.wasm';
let wasm;

function jsAdd(x, y) {
    return x + y
}
function init() {
    const go = new Go(); // Defined in wasm_exec.js

    go.importObject.env = go.importObject.env || {};
    go.importObject.env['command-line-arguments.add'] = jsAdd

    WebAssembly.instantiateStreaming(fetch(WASM_URL), go.importObject).then(function (obj) {
        wasm = obj.instance;
        go.run(wasm);
    })

    console.log(go.importObject);
    console.log('wasm loaded');
}

init();
