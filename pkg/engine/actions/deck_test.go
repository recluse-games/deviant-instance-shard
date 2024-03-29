package actions

import (
	"testing"

	"github.com/recluse-games/deviant-instance-shard/pkg/engine/enginetest"
	deviant "github.com/recluse-games/deviant-protobuf/genproto/go/instance_shard"
	"go.uber.org/zap/zaptest"
)

func TestDrawCard(t *testing.T) {
	logger := zaptest.NewLogger(t)
	logger.Sync()

	entityWithOneCardInDeckAndHand := &deviant.Encounter{
		ActiveEntity: &deviant.Entity{
			Id:   "0001",
			Deck: enginetest.GenerateDeckLiteral(1),
			Hand: enginetest.GenerateHandLiteral(1),
		},
	}

	if DrawCard(entityWithOneCardInDeckAndHand, logger.Sugar()) != true {
		t.Fail()
	}

	if len(entityWithOneCardInDeckAndHand.ActiveEntity.Deck.Cards) != 0 {
		t.Fail()
	}

	if len(entityWithOneCardInDeckAndHand.ActiveEntity.Hand.Cards) != 2 {
		t.Fail()
	}

	entityWithTwoCardsInDeckAndNoneInHand := &deviant.Encounter{
		ActiveEntity: &deviant.Entity{
			Id:   "0001",
			Deck: enginetest.GenerateDeckLiteral(2),
			Hand: enginetest.GenerateHandLiteral(0),
		},
	}

	if DrawCard(entityWithTwoCardsInDeckAndNoneInHand, logger.Sugar()) != true {
		t.Fail()
	}

	if len(entityWithTwoCardsInDeckAndNoneInHand.ActiveEntity.Deck.Cards) != 1 {
		t.Fail()
	}

	if len(entityWithTwoCardsInDeckAndNoneInHand.ActiveEntity.Hand.Cards) != 1 {
		t.Fail()
	}

	if DrawCard(entityWithTwoCardsInDeckAndNoneInHand, logger.Sugar()) != true {
		t.Fail()
	}

	if len(entityWithTwoCardsInDeckAndNoneInHand.ActiveEntity.Deck.Cards) != 0 {
		t.Fail()
	}

	if len(entityWithTwoCardsInDeckAndNoneInHand.ActiveEntity.Hand.Cards) != 2 {
		t.Fail()
	}
}
