.PHONY: generate
schema:
	pulumi package get-schema . > schema.json