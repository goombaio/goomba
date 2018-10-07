module github.com/goombaio/goomba

require (
	github.com/google/uuid v1.0.0
	github.com/goombaio/ansicolor v0.0.0-20180925191811-fa6507c8ad5d
	github.com/goombaio/cli v0.0.0-20181006234452-f8ff1984029b
	github.com/goombaio/log v0.0.0-20181006234330-b2d335e3400f
	github.com/goombaio/namegenerator v0.0.0-20181006234301-989e774b106e
)

replace (
	github.com/goombaio/ansicolor => ../ansicolor
	github.com/goombaio/cli => ../cli
	github.com/goombaio/log => ../log
	github.com/goombaio/namegenerator => ../namegenerator
)
