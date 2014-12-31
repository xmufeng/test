all: deps install

deps:
	@echo "handling dependents ..."
	@godep restore
	@echo "dependents handled"
install:
	@echo "installing application..."
	@godep go install .
	@echo "install done"
