.PHONY: clean all

all: c_static_lib go_executable

c_static_lib:
	/home/pdeng/Workspace2/toolchain-mipsel_24kec+dsp_gcc-4.8-linaro_uClibc-0.9.33.2/bin/mipsel-openwrt-linux-gcc -c greetings/*.c
	mkdir -p greetings/build/
	mv *.o greetings/build/
	/home/pdeng/Workspace2/toolchain-mipsel_24kec+dsp_gcc-4.8-linaro_uClibc-0.9.33.2/bin/mipsel-openwrt-linux-ar rs greetings/build/greetings.a greetings/build/*.o
	mv greetings/build/greetings.a greetings/build/liblanguage.a
	rm -rf greetings/build/*.o

go_executable:
	GOOS=linux GOARCH=mipsle CGO_ENABLED=1 CXX="/home/pdeng/Workspace2/toolchain-mipsel_24kec+dsp_gcc-4.8-linaro_uClibc-0.9.33.2/bin/mipsel-openwrt-linux-g++" CC="/home/pdeng/Workspace2/toolchain-mipsel_24kec+dsp_gcc-4.8-linaro_uClibc-0.9.33.2/bin/mipsel-openwrt-linux-gcc" CGO_CFLAGS="--sysroot /home/pdeng/Workspace2/toolchain-mipsel_24kec+dsp_gcc-4.8-linaro_uClibc-0.9.33.2/" CGO_LDFLAGS="--sysroot /home/pdeng/Workspace2/toolchain-mipsel_24kec+dsp_gcc-4.8-linaro_uClibc-0.9.33.2/" go build -o static_c_lib_mipsel

clean:
	rm -rf *.o *.a greetings/build static_c_lib_mipsel
