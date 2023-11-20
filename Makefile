JEKYLL_VERSION=4.2.2

serve:
	docker run --rm \
	  --volume=$(PWD):/srv/jekyll \
	  -p 4000:4000 \
	  -it jekyll/jekyll:$(JEKYLL_VERSION) \
	  jekyll serve --livereload
