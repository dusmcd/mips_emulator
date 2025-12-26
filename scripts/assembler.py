#! /usr/bin/env python3
from rtype_encoder import encode_assembly as encode_rtype
from itype_encoder import encode as encode_itype
import os

types = {
        "addi": 1, "add": 0, "sub": 0, "beq": 1
}

# checks whether the first word in the instruction
# is a label or an instruction
def get_instr_text(line):
    words = line.split()
    first_word = words[0].strip()
    if first_word[-1] == ":":
        return words[1], False
    return first_word, True

def main():
    if os.path.exists("bin"):
        os.remove("bin")

    with open("test.asm", "r") as file:
        assembly = file.read()

    with open("bin", "wb") as file:
        label_addrs = {}
        lines = filter(lambda line: len(line) > 0, assembly.split("\n"))
        for i, line in enumerate(lines):
            instruction = None

            (instr_text, is_instr) = get_instr_text(line)
            if not is_instr:
                label = line.split(maxsplit=1)[0].rstrip(": ")
                label_addrs[label] = i
                line = line.lstrip(f"{label}:")

            print(f"assembly instruction: {instr_text}")
            if types[instr_text] == 1:
                instruction = encode_itype(line)
            else:
                instruction = encode_rtype(line)

            
            data = int(instruction, 16).to_bytes(4, byteorder="little", signed=False)
            if file.write(data) != 4:
                raise Exception("error writing file")

    print("\nMachine Code (Hex)")
    print("===================")
    with open("bin", "rb") as file:
        while True:
            code = file.read(4)
            if not code:
                break
            code = hex(int.from_bytes(code, byteorder="little", signed=False))
            print(code)
    

if __name__ == "__main__":
    try:
        main()
    except Exception as e:
        print(f"error: {e}")
    except KeyError as ke:
        print(f"error: key not found: {ke}")

