FOLDERS = hello_world integers array_and_slice

testall:
	for dir in $(FOLDERS); do \
		make -C $$dir test; \
	done
