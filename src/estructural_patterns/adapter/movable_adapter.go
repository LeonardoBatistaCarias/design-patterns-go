package main

type MovableAdapter interface {
}

type MovableAdapterImpl struct {
	luxuryCar Movable
}

func (m *MovableAdapterImpl) GetSpeed() float32 {
	return convertMPHToKMPH(m.luxuryCar.GetSpeed())
}

func convertMPHToKMPH(mph float32) float32 {
	return mph * 1.60934
}
