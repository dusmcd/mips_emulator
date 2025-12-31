FLAGS= -mabi=32 -march=mips32
CMD= mipsel-linux-gnu-gcc

all: run

print.o: print.s
	$(CMD) $(FLAGS) -c $<

crt0.o: crt0.s
	$(CMD) $(FLAGS) -c $<

main.o: main.c
	$(CMD) $(FLAGS) -c $<

run: crt0.o print.o main.o
	$(CMD) $(FLAGS) -nostdlib -Wl,-Ttext-segment=0x10000 $^ -o main


