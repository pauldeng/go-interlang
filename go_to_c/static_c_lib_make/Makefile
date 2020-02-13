.PHONY: clean all

all: c_static_lib go_executable

c_static_lib:
	gcc -c greetings/*.c
	mkdir -p greetings/build/
	mv *.o greetings/build/
	ar rs greetings/build/greetings.a greetings/build/*.o
	mv greetings/build/greetings.a greetings/build/liblanguage.a
	rm -rf greetings/build/*.o

go_executable:
	go build -o static_c_lib

clean:
	rm -rf *.o *.a greetings/build static_c_lib
