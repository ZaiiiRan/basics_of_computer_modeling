package main

import (
	"lab8/internal/simulation"
)

func main() {
	sim := simulation.NewTunnelSim(1000, 1.0/12.0, 1.0/9.0, 30, 30, 60)
	sim.Run()
	sim.Analyze()
	sim.OptimizeDurations()
}
