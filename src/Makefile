
.PHONY: code-quality
code-quality:
	gometalinter --vendor --tests \
		--disable=gotype --disable=errcheck --disable=gas --disable=dupl \
		--deadline=1500s --checkstyle --sort=linter ./... > static-analysis.xml

.PHONY: dependencies
dependencies:
	glide install
