all: compile

compile:
		cd main; go build . &&  mv main ../bin/httpDns;

clean:
		cd ../bin; go clean; cd logic; go clean;

