.PHONY: schema
schema:
	pulumi package get-schema . > schema.json
	
.PHONY: publish
publish: schema
	pulumi package publish --readme README.md schema.json