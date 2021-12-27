package template

type RealTimeStrategyGame interface {
	initialize()
	start()
	end()
}

func play(g RealTimeStrategyGame) {
	g.initialize()
	g.start()
	g.end()
}

