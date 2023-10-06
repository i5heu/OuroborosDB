let go = new Go();
let wasmModule;

async function loadWasmModule() {
    const response = await fetch('test.wasm');
    const binary = await response.arrayBuffer();
    const result = await WebAssembly.instantiate(binary, go.importObject);
    wasmModule = result.instance;

    // Log the available exports
    console.log(wasmModule.exports);
    if (wasmModule.exports.add) {
        console.log("'add' function is present in exports.");
    } else {
        console.error("'add' function is NOT present in exports.");
    }
    const functionNames = Object.keys(wasmModule.exports).filter(key => typeof wasmModule.exports[key] === 'function');
    console.log("Exported functions:", functionNames);


    go.run(wasmModule);
}

loadWasmModule();