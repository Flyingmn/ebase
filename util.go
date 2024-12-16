package ebase

import (
	"github.com/Flyingmn/ebase/util/LinkTracking"
)

func (eb *Eb) initLinkTracking() {
	Producer := GetKafka(eb.Config.LinkTrack.KafkaProducerName)
	err := LinkTracking.InitLinkTracking(eb.Config, Producer)
	if err != nil {
		return
	}
}
