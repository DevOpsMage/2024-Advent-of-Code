import sys

def is_safe(row):
    """
    Determines if a row is 'safe'.
    A row is considered safe if the differences between each pair of
    successive numbers are all in the range [1, 3] (increasing)
    or all in the range [-3, -1] (decreasing).
    """
    # Rows with less than 2 elements are considered safe
    if len(row) < 2:
        return True

    # Calculate the first difference to determine the direction
    first_diff = row[1] - row[0]

    if 1 <= first_diff <= 3:
        # Sequence should be increasing
        for i in range(1, len(row) - 1):
            diff = row[i + 1] - row[i]
            if not (1 <= diff <= 3):
                return False
        return True
    elif -3 <= first_diff <= -1:
        # Sequence should be decreasing
        for i in range(1, len(row) - 1):
            diff = row[i + 1] - row[i]
            if not (-3 <= diff <= -1):
                return False
        return True
    else:
        # Not increasing or decreasing within allowed increments
        return False

def read_data(filename):
    """
    Reads the input data from the specified file.
    Each line in the file is expected to be a sequence of integers separated by spaces.
    Returns a list of lists of integers.
    """
    with open(filename) as f:
        # Read each line, split by spaces, and convert to integers
        return [
            [int(num) for num in line.strip().split()]
            for line in f if line.strip()
        ]

def main():
    # Check if input filename is provided
    if len(sys.argv) < 2:
        print("Usage: python script.py <input_file>")
        sys.exit(1)
    filename = sys.argv[1]

    # Read data from the input file
    data = read_data(filename)

    # Count the number of safe rows
    safe_count = sum(is_safe(row) for row in data)
    print(safe_count)

    # For each row, check if removing any one element makes it safe
    safe_count = sum(
        any(
            is_safe(row[:i] + row[i+1:])
            for i in range(len(row))
        )
        for row in data
    )
    print(safe_count)

if __name__ == "__main__":
    main()
