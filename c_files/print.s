.text
.globl add
.ent add

add:
  add $v0, $a0, $a1 # args passed to a0 and a1
  jr $ra

  nop
  .end add
