#!/usr/bin/sh
# link_script generates the executable file recipe-aggregator
cat <<'EOF' >recipe-aggregator
	#!/usr/bin/sh
	docker run --interactive \
		--tty \
		--rm \
		recipe-aggregator application "$@"
EOF