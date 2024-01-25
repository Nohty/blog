run:
	@templ generate
	@npx tailwindcss -i assets/tailwind.css -o assets/tailwind.min.css --minify
	@go run main.go