let go = new Go();
let wasmModule;

window.fetchFromGo = async function (url) {
    let response = await fetch(url, { mode: 'no-cors' });
    let body = await response.text();
    return body;
}

async function loadWasmModule() {
    const response = await fetch('test.wasm');
    const binary = await response.arrayBuffer();
    const result = await WebAssembly.instantiate(binary, go.importObject);
    wasmModule = result.instance;

    // Log the available exports
    const functionNames = Object.keys(wasmModule.exports).filter(key => typeof wasmModule.exports[key] === 'function');
    console.log("Exported functions:", functionNames);

    go.run(wasmModule);
}

async function callGoAndGetResponse(x) {
    getFileSystemAccess();
    return await callGoFetch(x);
}

loadWasmModule();


async function getFileSystemAccess() {
    let fs;
    if (window.showDirectoryPicker) {
        fs = await window.showDirectoryPicker();
    } else {
        fs = await window.showOpenFilePicker();
    }



    // read
    const file = await fs.getFileHandle('test.txt');
    const fileContents = await file.getFile();
    const text = await fileContents.text();
    console.log(text);

    // write to the file
    const writable = await fs.getFileHandle('test.txt', { create: true });
    const writableStream = await writable.createWritable();
    await writableStream.write(text + '\nHello World');
    writableStream.close();

}