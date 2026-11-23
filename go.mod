module github.com/cashfree/cashfree-pg/v6

go 1.18

require github.com/getsentry/sentry-go v0.29.1
require github.com/stretchr/testify v1.11.1

require (
)

retract (
	v6.0.8 // does not compile: duplicate model types, undefined Nullable* types
	v6.0.9 // does not compile: duplicate model types
)