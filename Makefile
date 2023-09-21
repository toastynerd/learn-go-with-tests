FOLDERS = hello_world integers array_and_slice test_lib

testall:
	for dir in $(FOLDERS); do \
		make -C $$dir test; \
	done
