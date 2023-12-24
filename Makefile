.PHONY: watch
watch:
	@echo "Watching for changes..."
	go run github.com/wolfeidau/reflex -c reflex.conf

.PHONY: install-assets
install-assets:
	@echo "Installing assets..."
	cd assets && npm install

.PHONY: build-assets
build-assets:
	@echo "Building assets..."
	@cd assets && rm -rf ./dist && ./node_modules/.bin/esbuild src/index.ts \
			--bundle --bundle --minify --sourcemap --target=chrome58,firefox57,safari11,edge16 \
			--outfile=dist/bundle.js
