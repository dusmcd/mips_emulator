#! /usr/bin/env python3

def encode(assembly, label_map):
    components = assembly.split()
    instr = components[0].strip()
    label = components[1].strip()
    target_addr = (label_map[label] + 1) * 4
    binary_str = "0b"
    if instr == "j":
        binary_str += "000010"
    elif instr == "jal":
        binary_str += "000011"

    target_bin = bin(target_addr)[2:]
    # add padding and cut off last two bits
    target_bin = "00000000000000000000000000000000" + target_bin[:-2]

    # get last 26 bits
    target_bin = target_bin[-26:]
    binary_str += target_bin

    return hex(int(binary_str, 2))

def main():
    assembly = input("Enter assembly: ")
    label_map = {"done": 4000}
    machine_code = encode(assembly, label_map)
    print(f"Machine code hex: {machine_code}")

if __name__ == "__main__":
    main()
