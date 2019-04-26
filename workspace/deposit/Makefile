contract:=deposit

default:go

sol:
	solc --abi ${contract}.sol | awk '/JSON ABI/{x=1;next}x' > ${contract}.json
	solc --bin ${contract}.sol | awk '/Binary:/{x=1;next}x' > ${contract}.bin

go:
	abigen --abi=${contract}.json --pkg=${contract} --out=${contract}.go

clean:
	rm ${contract}.go