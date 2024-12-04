import re
import sys

def process_multiplications_basic(text):
    pattern = r'mul\((\d+),(\d+)\)'
    results = []
    
    matches = re.finditer(pattern, text)
    for match in matches:
        x, y = int(match.group(1)), int(match.group(2))
        if 1 <= x <= 999 and 1 <= y <= 999:
            results.append(x * y)
    
    return sum(results)

def process_multiplications_with_control(text):
    pattern = r'mul\((\d+),(\d+)\)'
    results = []
    
    chars = list(text)
    multiply_enabled = True
    i = 0
    
    while i < len(chars):
        if i + 4 < len(chars) and ''.join(chars[i:i+5]) == "don't":
            multiply_enabled = False
            i += 5
            continue
            
        if i + 2 < len(chars) and ''.join(chars[i:i+2]) == "do":
            if i + 2 >= len(chars) or chars[i+2] != 'n':
                multiply_enabled = True
            i += 2
            continue
            
        if multiply_enabled and chars[i] == 'm':
            substring = ''.join(chars[i:i+15])
            match = re.match(pattern, substring)
            
            if match:
                x, y = int(match.group(1)), int(match.group(2))
                if 1 <= x <= 999 and 1 <= y <= 999:
                    results.append(x * y)
        
        i += 1
    
    return sum(results)

def main():
    if len(sys.argv) != 2:
        print("Usage: python script.py <input_file>")
        sys.exit(1)
        
    try:
        with open(sys.argv[1], 'r') as file:
            content = file.read()
        
        basic_total = process_multiplications_basic(content)
        controlled_total = process_multiplications_with_control(content)
        
        print(f"Basic sum (without do/don't): {basic_total}")
        print(f"Controlled sum (with do/don't): {controlled_total}")
        
    except FileNotFoundError:
        print(f"Error: File '{sys.argv[1]}' not found")
        sys.exit(1)
    except Exception as e:
        print(f"Error processing file: {e}")
        sys.exit(1)

if __name__ == "__main__":
    main()