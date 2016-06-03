package scheduler

import (
	"github.com/klarna/eremetic/types"
	mesos "github.com/mesos/mesos-go/mesosproto"
)

const (
	constraintEqualsOp = "EQUALS"
)

type constraint struct {
	Field    string
	Operator string
	Value    string
}

func checkConstraints(offer *mesos.Offer, constraints []types.Constraint) bool {
	for _, cons := range constraints {
		if len(cons) == 3 {
			cons := &constraint{
				Field:    cons[0],
				Operator: cons[1],
				Value:    cons[2],
			}

			if cons.Operator == constraintEqualsOp {
				if !equalMatches(offer, cons) {
					return false
				}
			}
		}
	}
	return true
}

func equalMatches(offer *mesos.Offer, cons *constraint) bool {
	for _, a := range offer.GetAttributes() {
		if a.GetName() == cons.Field && a.GetText().GetValue() == cons.Value {
			return true
		}
	}
	return false
}
