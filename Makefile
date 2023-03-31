GO:=docker compose run --rm go

# $(call go-run,type,name)
define go-run
	@if [ ! -r "./$(1)/$(2)/$(2).go" ]; then \
		echo "ERROR: Please enter a valid name. Usage: 'make run-$(1) name=<$(1)-name>'"; \
		exit 1; \
	else \
		$(GO) run $(1)/$(2)/$(2).go; \
	fi
endef

version:
	$(GO) version
run-demo:
	$(call go-run,demo,$(name))

run-exercise:
	$(call go-run,exercise,$(name))
