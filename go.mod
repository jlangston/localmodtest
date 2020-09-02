module github.com/jlangston/localmodtest

go 1.15

require (
	github.com/jlangston/localmodtest/hello v0.0.0
	github.com/jlangston/localmodtest/utils v0.0.0

)

replace (
	github.com/jlangston/localmodtest/hello => ./hello
	github.com/jlangston/localmodtest/utils => ./utils
)
