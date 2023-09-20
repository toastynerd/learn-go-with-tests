FOLDERS = hello_world integers

testall:
	for dir in $(FOLDERS); do \
		make -C $$dir test; \
	done
