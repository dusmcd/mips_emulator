addi $s0, $zero, 10
Loop: beq $t0, $s0, Done
addi $t1, $t1, 2
addi $t0, $t0, 1
j Loop

Done: addi $v0, $zero, 10 
