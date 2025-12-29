#! /usr/bin/env python3
from defs import regs
from rtype_encoder import get_correct_bits

CURRENT_LINE=10
LABEL_MAP = {"done": 25}

ops = {
        "lw": 0x23, "sw": 0x2B, "addi": 0x08, "beq": 0x04,
        "bne": 0x05, "blez": 0x06, "bgtz": 0x07
}

def encode(assembly, current_line=0, label_map=None):
    components = assembly.split(maxsplit=1)
    instr = components[0]
    registers = components[1].split(",")
    rt = registers[0].strip() # destination address

    op_bin = get_correct_bits(ops[instr], 6)
    rs = None
    imm = None
    
    if instr in ["lw", "sw"]:
        # get offset = immediate
        # get base addr in parens
        addr = registers[1].strip()
        (imm, rs) = addr.split("(")
        rs = rs.rstrip(")")
    elif instr in ["beq", "bne"]:
        # imm = label, will need to calculate
        # PC offset
        rs = registers[0].strip()
        rt = registers[1].strip()
        label = registers[2].rstrip(": ")
        label_line = label_map[label.strip()]

        # need to account for PC being on next instruction
        imm = label_line - (current_line + 1)   
    elif instr in ["blez", "bgtz"]:
        rt = "$zero"
        rs = registers[0].strip()
        label = registers[1].strip()
        label_line = label_map[label]

        imm = label_line - (current_line + 1)
    else:
        rs = registers[1].strip()
        imm = registers[2].strip()


    rs_bin = get_correct_bits(regs[rs], 5)
    rt_bin = get_correct_bits(regs[rt], 5)
    imm_bin = get_correct_bits(int(imm), 16)

    binary_str = f"0b{op_bin}{rs_bin}{rt_bin}{imm_bin}"
    return hex(int(binary_str, 2))

def main():
    assemblyInstruction = input("Assembly Instruction: ")
    machineCode = encode(assemblyInstruction, CURRENT_LINE, LABEL_MAP)
    print(f"Machine code hex: {machineCode}")

if __name__ == "__main__":
    main()
