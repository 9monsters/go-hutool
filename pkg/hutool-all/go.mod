module github.com/nine-monsters/go-hutool/hutool-all

go 1.19

require (
	github.com/nine-monsters/go-hutool/hutool-core latest
)

replace (
	github.com/nine-monsters/go-hutool/hutool-cor latest => ../hutool-core latest
)