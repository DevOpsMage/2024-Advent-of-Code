const fs = require('fs');

function isSafe(row) {
    /**
     * Determines if a row is 'safe'.
     * A row is considered safe if the differences between each pair of
     * successive numbers are all in the range [1, 3] (increasing)
     * or all in the range [-3, -1] (decreasing).
     */
    // Rows with less than 2 elements are considered safe
    if (row.length < 2) {
        return true;
    }

    // Calculate the first difference to determine the direction
    const firstDiff = row[1] - row[0];

    if (1 <= firstDiff && firstDiff <= 3) {
        // Sequence should be increasing
        for (let i = 1; i < row.length - 1; i++) {
            const diff = row[i + 1] - row[i];
            if (!(1 <= diff && diff <= 3)) {
                return false;
            }
        }
        return true;
    } else if (-3 <= firstDiff && firstDiff <= -1) {
        // Sequence should be decreasing
        for (let i = 1; i < row.length - 1; i++) {
            const diff = row[i + 1] - row[i];
            if (!(-3 <= diff && diff <= -1)) {
                return false;
            }
        }
        return true;
    } else {
        // Not increasing or decreasing within allowed increments
        return false;
    }
}

function readData(filename) {
    /**
     * Reads the input data from the specified file.
     * Each line in the file is expected to be a sequence of integers separated by spaces.
     * Returns an array of arrays of integers.
     */
    // Read the file synchronously
    const content = fs.readFileSync(filename, 'utf8');

    // Split into lines, filter out empty lines
    const lines = content.split('\n').filter(line => line.trim() !== '');

    // For each line, split by spaces and parse to integers
    return lines.map(line =>
        line.trim().split(/\s+/).map(num => parseInt(num, 10))
    );
}

function main() {
    // Check if input filename is provided
    if (process.argv.length < 3) {
        console.log('Usage: node script.js <input_file>');
        process.exit(1);
    }

    const filename = process.argv[2];

    // Read data from the input file
    const data = readData(filename);

    // Count the number of safe rows
    const safeCount = data.reduce((count, row) => count + (isSafe(row) ? 1 : 0), 0);
    console.log(safeCount);

    // For each row, check if removing any one element makes it safe
    const safeCountAfterRemoval = data.reduce((count, row) => {
        const isAnySafe = row.some((_, i) => {
            const newRow = row.slice(0, i).concat(row.slice(i + 1));
            return isSafe(newRow);
        });
        return count + (isAnySafe ? 1 : 0);
    }, 0);
    console.log(safeCountAfterRemoval);
}

main();
