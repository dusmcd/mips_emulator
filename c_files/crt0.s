.text
.globl __start
.ent __start

__start:
  jal main
  nop

  addi $a0, $v0, 0

  addi $v0, $zero, 4001 # exit the program
  syscall

  .end __start
