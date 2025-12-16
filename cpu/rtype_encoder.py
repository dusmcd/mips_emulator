regs = {
        "$zero": 0, "$at": 1, "$v0": 2, "$v1": 3,
        "$a0:": 4, "$a1": 5, "$a2": 6, "$a3": 7,
        "$t0": 8, "$t1": 9, "$t2": 10, "$t3": 11,
        "$t4": 12, "$t5": 13, "$t6": 14, "$t7": 15,
        "$s0": 16, "$s1": 17, "$s2": 18, "$s3": 19,
        "$s4": 20, "$s5": 21, "$s6": 22, "$s7": 23,
        "$t8": 24, "$t9": 25, "$k0": 26, "$k1": 27,
        "$gp": 28, "$sp": 29, "$fp": 30, "$ra": 31
    }

funcs = {
    "add": 0x20, "addu": 0x21, "sub": 0x22, "subu": 0x23,
    "mult": 0x18, "multu": 0x19, "div": 0x1A, "divu": 0x1B,
    "and": 0x24, "or": 0x25, "xor": 0x26, "nor": 0x27
}

def get_correct_bits(num, bits):
    binary_str = bin(num)[2:]
    binary_str = "00000000" + binary_str # padding
    return binary_str[-bits:] # return last n number of bits


def encode_assembly(asm):
    components = asm.split(maxsplit=1)
    binary_str = "0b000000"
    instr = components[0]
    registers = components[1].split(",")
    if instr in ["mult", "div", "multu", "divu"]:
        rd = "$zero"
        rs = registers[0].strip()
        rt = registers[1].strip()
    else:
        rd = registers[0].strip()
        rs = registers[1].strip()
        rt = registers[2].strip()

    rs_bin = get_correct_bits(regs[rs], 5)
    rt_bin = get_correct_bits(regs[rt], 5)
    rd_bin = get_correct_bits(regs[rd], 5)
    func_bin = get_correct_bits(funcs[instr], 6)
    shift = "00000"

    binary_str += rs_bin + rt_bin + rd_bin + shift + func_bin

    return hex(int(binary_str, 2))

def main():
    assembly = input("Assembly Instruction: ")
    print(f"Encoded hex instruction: {encode_assembly(assembly)}")

if __name__ == "__main__":
    main()
