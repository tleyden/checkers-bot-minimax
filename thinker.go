package main

import (
	"github.com/couchbaselabs/logg"
	cbot "github.com/tleyden/checkers-bot"
	core "github.com/tleyden/checkers-core"
)

type MinimaxThinker struct {
	ourTeamId cbot.TeamType
}

func (t MinimaxThinker) Think(gameState cbot.GameState) (bestMove cbot.ValidMove, ok bool) {

	ok = true
	ourTeam := gameState.Teams[t.ourTeamId]
	allValidMoves := ourTeam.AllValidMoves()
	if len(allValidMoves) > 0 {

		// convert into core.board representation
		board := gameState.Export()
		logg.LogTo("DEBUG", "Before move %v", board.CompactString(true))

		// generate best move (will be a core.move) -- initially, pick random
		move := t.generateBestMove(board)

		// search allValidMoves to find corresponding valid move
		found, bestValidMoveIndex := cbot.CorrespondingValidMoveIndex(move, allValidMoves)

		if !found {
			msg := "Could not find corresponding valid move: %v in %v"
			logg.LogPanic(msg, move, allValidMoves)
		} else {
			bestMove = allValidMoves[bestValidMoveIndex]
		}

		// this is just for debugging purposes
		player := cbot.GetCorePlayer(t.ourTeamId)
		boardPostMove := board.ApplyMove(player, move)
		logg.LogTo("DEBUG", "After move %v", boardPostMove.CompactString(true))

		return

	} else {
		ok = false
	}

	return
}

func (t MinimaxThinker) generateBestMove(board core.Board) core.Move {
	evalFunc := core.DefaultEvaluationFunction()
	player := cbot.GetCorePlayer(t.ourTeamId)
	depth := 8 // with 9 moves, I saw cases where it took 5m with no move
	bestMove, scorePostMove := board.Minimax(player, depth, evalFunc)
	logg.LogTo("DEBUG", "scorePostMove: %v", scorePostMove)
	return bestMove
}

func (t MinimaxThinker) GameFinished(gameState cbot.GameState) (shouldQuit bool) {
	shouldQuit = false
	return
}

func init() {
	logg.LogKeys["MAIN"] = true
	logg.LogKeys["CHECKERSBOT"] = true
	logg.LogKeys["DEBUG"] = true
}

func main() {
	checkersBotFlags := cbot.ParseCmdLine()
	thinker := &MinimaxThinker{}
	thinker.ourTeamId = checkersBotFlags.Team
	game := cbot.NewGame(thinker.ourTeamId, thinker)
	game.SetServerUrl(checkersBotFlags.SyncGatewayUrl)
	game.SetFeedType(checkersBotFlags.FeedType)
	game.SetDelayBeforeMove(checkersBotFlags.RandomDelayBeforeMove)
	game.GameLoop()
}
