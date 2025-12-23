#! /usr/bin/env python3
from defs import regs
from rtype_encoder import get_correct_bits

ops = {
        "lw": 0x23, "sw": 0x2B, "addi": 0x08
}

def encode(assembly):
    components = assembly.split(maxsplit=1)
    instr = components[0]
    registers = components[1].split(",")
    rt = registers[0].strip() # destination address

    op_bin = get_correct_bits(ops[instr], 6)
    rt_bin = get_correct_bits(regs[rt], 5)
    rs = None
    imm = None
    
    if instr in ["lw", "sw"]:
        # get offset = immediate
        # get base addr in parens
        addr = registers[1].strip()
        (imm, rs) = addr.split("(")
        rs = rs.rstrip(")")
    else:
        rs = registers[1].strip()
        imm = registers[2].strip()


    rs_bin = get_correct_bits(regs[rs], 5)
    imm_bin = get_correct_bits(int(imm), 16)

    binary_str = f"0b{op_bin}{rs_bin}{rt_bin}{imm_bin}"
    return hex(int(binary_str, 2))

def main():
    assemblyInstruction = input("Assembly Instruction: ")
    machineCode = encode(assemblyInstruction)
    print(f"Machine code hex: {machineCode}")

if __name__ == "__main__":
    main()
