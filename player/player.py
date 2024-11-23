#!/bin/python3
from proboj import *
import random

class Hrac(Game):
    def make_turn(self) -> Turn:
        self.log(self.player)
        self.log("WeaponNone.stats.Range: ", self.player.weapon.stats.Range)
        
        return MoveTurn(XY(random.randint(-100,100),random.randint(-100,100)))


if __name__ == '__main__':
    g = Hrac()
    g.run()
    g.log(f"SOM f{g.player}")
