import json
from abc import abstractstaticmethod


class XY:
    def __init__(self, x: float, y: float):
        self.x = x
        self.y = y

    @classmethod
    def from_json(cls, data: dict):
        return cls(data.get('x'), data.get('y'))


class Wall:
    def __init__(self, data: dict):
        self.a = XY.from_json(data.get('a'))
        self.b = XY.from_json(data.get('b'))


class Map:

    def __init__(self, data: dict):
        self.radius = data.get('radius')
        self.walls = []
        for wall in data.get('walls'):
            self.walls.append(Wall(wall))

    @classmethod
    def read_map(cls):
        map_json = input()
        input()
        return Map(json.load(map_json))


class Player:
    def __init__(self, data):
        self.xy = XY(data.get('x'), data.get('y'))
        self.id = data.get('id')
        self.health = data.get('health')
        self.weapon = data.get('weapon')
        self.laoded_ammo = data.get('laoded_ammo')
        self.reload_cooldown = data.get('reload_cooldown')

    @classmethod
    def read_player(cls, data: dict):
        return cls(data)


class EnemyPlayer:
    def __init__(self, data):
        self.xy = XY(data.get('x'), data.get('y'))
        self.id = data.get('id')
        self.weapon = data.get('weapon')

    @classmethod
    def read_player(cls, data: dict):
        return cls(data)

    def __eq__(self, other):
        return self.id == other.id
    def __hash__(self):
        return self.id

class Item:
    def __init__(self, data):
        self.xy = XY(data.get('x'), data.get('y'))
        self.type = data.get('type')
        self.weapon = data.get('weapon')

    @classmethod
    def read_item(cls, data: dict):
        return cls(data)

class Turn:
    def print(self):
        print(".")

class MoveTurn(Turn):
    def __init__(self, goal: XY):
        self.goal = goal


    def print(self):
        print(f"MOVE {self.goal.x} {self.goal.y}")
        print(".")


class Game:

    def __init__(self):
        self.map: Map.read_map()
        self.player: Player
        self.enemy_players: list[EnemyPlayer]
        self.items = list[Item]

    def _send_turn(self, turn: Turn):
        turn.print()

    def _read_state(self):
        data = json.loads(input())
        self.player = Player.read_player(data.get('player'))
        self.items = [Item(item) for item in data.get('visible_items')]
        self.enemy_players = [EnemyPlayer(p) for p in data.get('visible_players')]
        input()

    def make_turn(self) -> Turn:
        raise NotImplementedError()

    def run(self):
        while True:
            self._read_state()
            turn = self.make_turn()
            self._send_turn(turn)








if __name__ == "__main__":
    print("PROBOOOOJ")
