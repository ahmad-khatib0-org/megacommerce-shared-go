.PHONY: release

release:
ifeq ($(VERSION),)
	@echo "❌ ERROR: No version specified for release."
	@exit 1
endif
ifeq ($(MSG),)
	@echo "❌ ERROR: No message specified for release. Use MSG='Your message'"
	@exit 1
endif
	@echo "Committing generated files and tagging release..."
	git add -A
	git commit -m "$(MSG)" || true

	# create annotated tag (skip if exists)
	if git rev-parse "refs/tags/v$(VERSION)" >/dev/null 2>&1; then \
	  echo "Tag v$(VERSION) already exists — updating commit (force)"; \
	  git tag -f -a "v$(VERSION)" -m "$(MSG)"; \
	else \
	  git tag -a "v$(VERSION)" -m "$(MSG)"; \
	fi
	git push origin --tags || true
	git push origin main
	@echo "✅ Release complete"
