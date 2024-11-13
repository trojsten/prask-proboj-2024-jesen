#!/bin/python3
from weapons import *
from proboj import *

class Hrac(Game):
    def make_turn(self) -> Turn:
        self.log("WeaponNone.stats.Range: ", WeaponNone.stats.Range)
        return MoveTurn(XY(0,0))


if __name__ == '__main__':
    g = Hrac()
    g.log(f"SOM f{g.player}")
    g.run()
