module github.com/goombaio/goomba

require (
	github.com/google/uuid v1.0.0
	github.com/goombaio/ansicolor v0.0.0-20180925191811-fa6507c8ad5d
	github.com/goombaio/cli v0.0.0-20181004220610-49e3e37f249c
	github.com/goombaio/log v0.0.0-20181004215944-30400b5e9d52
	github.com/goombaio/namegenerator v0.0.0-20180925151310-6b8631dcf92d
)

replace (
	github.com/goombaio/ansicolor => ../ansicolor
	github.com/goombaio/cli => ../cli
	github.com/goombaio/log => ../log
	github.com/goombaio/namegenerator => ../namegenerator
)
