import json


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


class Item:
    def __init__(self, data):
        self.xy = XY(data.get('x'), data.get('y'))
        self.type = data.get('type')
        self.weapon = data.get('weapon')

    @classmethod
    def read_item(cls, data: dict):
        return cls(data)


def read_state() -> (Player, list[Item], list[EnemyPlayer]):
    data = json.loads(input())
    items = [Item(item) for item in data.get('visible_items')]
    enemy_players = [EnemyPlayer(p) for p in data.get('visible_players')]
    return Player.read_player(data.get('player')), items, enemy_players


if __name__ == "__main__":
    print("PROBOOOOJ")
