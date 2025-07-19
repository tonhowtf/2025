defmodule ExMon do
  alias ExMon.{Game, Player}

  def create_player(name, move_avg, move_rnd, move_heal) do
    Player.build(name, move_rnd, move_avg, move_heal)
  end

  def start_game(player) do
    "marvin"
    |> create_player(:punch, :kick, :heal)
    |> Game.start(player)
  end
end
