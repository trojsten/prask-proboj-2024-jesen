from proboj import *

class Hrac(Game):
    def make_turn(self) -> Turn:
        return MoveTurn(XY(0,0))


if __name__ == '__main__':
    g = Hrac()
    g.run()