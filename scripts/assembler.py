#! /usr/bin/env python3
from rtype_encoder import encode_assembly as encode_rtype
from itype_encoder import encode as encode_itype
import os

types = {
        "addi": 1, "add": 0, "sub": 0
}

def main():
    if os.path.exists("bin"):
        os.remove("bin")

    with open("test.asm", "r") as file:
        assembly = file.read()

    with open("bin", "wb") as file:
        lines = filter(lambda line: len(line) > 0, assembly.split("\n"))
        for line in lines:
            instruction = None
            instr_text = line.split()[0].strip()
            print(f"assembly instruction: {instr_text}")
            if types[line.split(maxsplit=1)[0].strip()] == 1:
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

