module github.com/goombaio/goomba

require (
	github.com/google/uuid v1.0.0
	github.com/goombaio/ansicolor v0.0.0-20180925191811-fa6507c8ad5d
	github.com/goombaio/cli v0.0.0-20180925150851-02c676757da7
	github.com/goombaio/log v0.0.0-20180925151324-4c3fe5c2e684
	github.com/goombaio/namegenerator v0.0.0-20180925151310-6b8631dcf92d
)

replace (
	github.com/goombaio/ansicolor => ../ansicolor
	github.com/goombaio/cli => ../cli
	github.com/goombaio/log => ../log
	github.com/goombaio/namegenerator => ../namegenerator
)
