import { open } from 'node:fs/promises';

var current = 0,
    max = 0

function onUserElfChanged() {
    if (current >= max) {
        max = current;
    }
    current = 0;
}

const file = await open('../data.txt');
for await (const line of file.readLines()) {
    if (line === '') {
        onUserElfChanged()
        continue
    }
    current += parseInt(line)
}

console.log(max)
