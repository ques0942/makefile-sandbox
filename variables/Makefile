VAR := make_var

## VAR is not envvar in shell
.PHONY: case1
case1: ## VAR is empty
	@go run penv.go

## `export` doesn't affect other commands in the same rule
.PHONY: case2
case2: ## VAR is empty
	export VAR=$(VAR)
	@go run penv.go

## envvar in the same line affects command
.PHONY: case3
case3: ## VAR is make_var
	VAR=$(VAR) go run penv.go

## `export` in eval affects other commands in the same rule
.PHONY: case4
case4: ## VAR is make_var
	$(eval export VAR=$(VAR))
	@go run penv.go

## `export` in eval affects other commands in the same rule
.PHONY: case5
case5: ## VAR is updated
	$(eval export VAR=updated)
	@go run penv.go

## can use eval and shell at the same time
.PHONY: case6
case6: ## VAR is $(hostname)
	$(eval export VAR=$(shell hostname))
	@go run penv.go

.PHONY: define
define: ## define
	$(eval export VAR=defined)
	@echo ${VAR}

## target is affected by dependent target
.PHONY: case7
case7: define ## define\nVAR is define
	@go run penv.go

## called target doesn't affect other commands in the same rule
.PHONY: case8
case8: ## define\nVAR is empty
	$(MAKE) define
	@go run penv.go

## can use other target as shell command
.PHONY: case9
case9: ## VAR is define
	$(eval export VAR=$(shell $(MAKE) define))
	@go run penv.go

## parent envvar affects child target
.PHONY: case10
case10:
	VAR=parent $(MAKE) case1 ## VAR is make_var ## This is hard to understand. does make update envvar by own variables?
	$(MAKE) VAR=parent case1 ## VAR is parent
	$(eval export VAR=export_parent)
	$(MAKE) case1 ## VAR is make_var ## This is hard to understand. does make update envvar by own variables?
