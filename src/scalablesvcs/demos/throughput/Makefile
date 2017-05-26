prepare:
	touch benches/writemem.out
benchmem:
	mv benches/writemem.out benches/writemem.old
	go test -bench=Write > benches/writemem.out

benchfile:
	go test -bench=Write -usefile true > benches/writemem.out

benchcmpmem:
	benchcmp benches/writemem.old benches/writemem.out > benches/memdelta.out
	cat benches/memdelta.out

compare: prepare benchmem benchfile benchcmpmem
	./getDelta.sh benches/memdelta.out
