const fs = require('fs');

const data = fs.readFileSync(process.argv[2], 'utf8').trim();

let left = [];
let right = [];

data.split('\n').forEach(line => {
    let [nl, nr] = line.split('   ');
    left.push(parseInt(nl));
    right.push(parseInt(nr));
});

let sortedLeft = left.slice().sort((a, b) => a - b);
let sortedRight = right.slice().sort((a, b) => a - b);

let p1 = sortedLeft.map((nl, i) => Math.abs(nl - sortedRight[i]));
console.log(`Part 1: ${p1.reduce((a, b) => a + b, 0)}`);

let c = {};

for (let n of right) {
    c[n] = (c[n] || 0) + 1;
}

let p2 = left.map(n => n * (c[n] || 0));
console.log(`Part 2: ${p2.reduce((a, b) => a + b, 0)}`);
