import { open } from 'node:fs/promises';

var current = 0,
    max = [0]

function onUserElfChanged() {

    const cur = current
    current = 0

    if (max.length < 3) {
        max.push(cur)
        return
    }

    let min = max[0]
    let minIndex = 0;
    for (let i = 0; i < max.length; ++i) {
        if (max[i] < min) {
            min = max[i]
            minIndex = i
        }
    }

    if (min < cur) {
        max[minIndex] = cur 
    }

}

const file = await open('../data.txt');
for await (const line of file.readLines()) {
    if (line === '') {
        onUserElfChanged()
        continue
    }
    current += parseInt(line)
}

console.log(max.reduce((a, b) => a + b, 0))
