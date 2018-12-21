all: main.go
	go build $<

run: main.go
	go run $<

clean:
	rm -f lib.so main .DS_Store

so:
	go build -v -x -buildmode=c-shared -o lib.so

lib:
	go install -buildmode=shared -linkshared lib

exe:
	go build -buildmode=exe

help: 
	@echo "===============A common Makefilefor c programs==============" 
	@echo "The following targets aresupport:" 
	@echo 
	@echo " all              - (==make) compile and link" 
	@echo " clean            - clean target" 
	@echo " run        		 - run the programs" 
	@echo " help             - print help information" 
	@echo 
	@echo "To make a target, do 'make[target]'" 
	@echo "========================= Version2.2 ======================="  