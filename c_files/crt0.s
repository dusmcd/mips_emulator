.text
.globl __start
.ent __start

__start:
  jal main
  nop

  move $a0, $v0

  li $v0, 10
  syscall

  .end __start
