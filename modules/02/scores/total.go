package scores

func MaxBreak() int {
	return (RedValue+BlackValue)*Reds + YellowValue + GreenValue +
		BrownValue + BlueValue + PinkValue + BlackValue
}
