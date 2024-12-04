const fs = require('fs');

function processMultiplicationsBasic(text) {
    const pattern = /mul\((\d+),(\d+)\)/g;
    const results = [];
    let match;
    
    while ((match = pattern.exec(text)) !== null) {
        const x = parseInt(match[1]);
        const y = parseInt(match[2]);
        
        if (x >= 1 && x <= 999 && y >= 1 && y <= 999) {
            results.push(x * y);
        }
    }
    
    return results.reduce((sum, val) => sum + val, 0);
}

function processMultiplicationsWithControl(text) {
    const pattern = /mul\((\d+),(\d+)\)/;
    const results = [];
    const chars = text.split('');
    let multiplyEnabled = true;
    let i = 0;
    
    while (i < chars.length) {
        // Check for "don't" - must check first before "do"
        if (i + 4 < chars.length && chars.slice(i, i + 5).join('') === "don't") {
            multiplyEnabled = false;
            i += 5;
            continue;
        }
        
        // Check for standalone "do" (not part of "don't")
        if (i + 1 < chars.length && chars.slice(i, i + 2).join('') === "do") {
            // Make sure this "do" is not part of a "don't"
            if (i === 0 || chars.slice(i - 3, i).join('') !== "don") {
                multiplyEnabled = true;
            }
            i += 2;
            continue;
        }
        
        // Only process multiplication if enabled
        if (multiplyEnabled && chars[i] === 'm') {
            const substring = chars.slice(i, i + 15).join('');
            const match = substring.match(pattern);
            
            if (match) {
                const x = parseInt(match[1]);
                const y = parseInt(match[2]);
                if (x >= 1 && x <= 999 && y >= 1 && y <= 999) {
                    results.push(x * y);
                    i += match[0].length - 1; // Skip the entire multiplication pattern
                }
            }
        }
        
        i++;
    }
    
    return results.reduce((sum, val) => sum + val, 0);
}

function main() {
    if (process.argv.length !== 3) {
        console.log("Usage: node script.js <input_file>");
        process.exit(1);
    }
    
    try {
        const content = fs.readFileSync(process.argv[2], 'utf8');
        const basicTotal = processMultiplicationsBasic(content);
        const controlledTotal = processMultiplicationsWithControl(content);
        
        console.log(`Basic sum (without do/don't): ${basicTotal}`);
        console.log(`Controlled sum (with do/don't): ${controlledTotal}`);
        
    } catch (error) {
        console.error(`Error processing file: ${error.message}`);
        process.exit(1);
    }
}

main();