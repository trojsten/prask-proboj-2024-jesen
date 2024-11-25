/**
 * @typedef Map
 * @prop {number} radius
 * @prop {Wall[]} walls
 */

/**
 * @typedef Position
 * @prop {number} x
 * @prop {number} y
 */

/**
 * @typedef Wall
 * @prop {Position} a
 * @prop {Position} b
 */

/**
 * @typedef Frame
 * @prop {number} radius
 * @prop {MapItem[]} items
 * @prop {Player[]} players
 */

/**
 * @typedef MapItem
 * @prop {number} x
 * @prop {number} y
 * @prop {number} type
 * @prop {number} weapon
 */

/**
 * @typedef Player
 * @prop {number} x
 * @prop {number} y
 * @prop {number} id
 * @prop {string} name
 * @prop {number} health
 * @prop {number} weapon
 * @prop {number} score
 * @prop {number} loaded_ammo
 * @prop {number} reload_cooldown
 */

class Renderer {
    constructor() {
        var width = window.innerWidth;
        var height = window.innerHeight;

        this.canvas = new Konva.Stage({
            container: "js-canvas",
            width: width,
            height: height,
        })
        this.playerLayers = {}
        this.mapLayer = new Konva.Layer()
        this.scoreboardLayer = new Konva.Layer()
        this.scoreboardLayer.x(width - 205)
        this.scoreboardLayer.y(5)
        this.itemGroup = new Konva.Group()
        this.mapLayer.add(this.itemGroup)

        this.canvas.add(this.mapLayer)
        this.canvas.add(this.scoreboardLayer)
    }

    /** @type {Player[]} players */
    renderScoreboard(players) {
        this.scoreboardLayer.removeChildren()

        players.sort((a, b) => a.score - b.score)
        for (let i = 0; i < players.length; i++) {
            let group = new Konva.Group()
            let Y = 35*i
            let r = new Konva.Rect({
                x: 0,
                y: Y,
                width: 200,
                height: 30,
                fill: "black",
            })
            group.add(r)

            let name = new Konva.Text({
                x: 5,
                y: Y+3,
                text: players[i].name,
                fontSize: 16,
                fontStyle: "bold",
                fontFamily: 'Arial',
                fill: 'white',
            });
            group.add(name)

            let score = new Konva.Text({
                x: 200-18,
                y: Y+4,
                text: players[i].score,
                fontSize: 24,
                fontStyle: "bold",
                align: "right",
                fontFamily: 'Arial',
                fill: 'white',
            });
            group.add(score)

            let hp = new Konva.Text({
                x: 5,
                y: Y+19,
                text: players[i].health + " HP",
                fontSize: 10,
                fontFamily: 'Arial',
                fill: 'white',
            });
            group.add(hp)

            if (players[i].health <= 0) {
                group.opacity(0.1)
            }

            this.scoreboardLayer.add(group)
        }
    }

    /** @type {Map} map */
    loadMap(map) {
        this.map = map
        this.mapBorder = new Konva.Circle({
            x: 0,
            y: 0,
            radius: this.map.radius,
            stroke: 'red',
            strokeWidth: 2
        })
        this.mapLayer.add(this.mapBorder)

        for (const wall of this.map.walls) {
            let wallObject = new Konva.Line({
                points: [wall.a.x, wall.a.y, wall.b.x, wall.b.y],
                stroke: "white"
            })
            this.mapLayer.add(wallObject)
        }
    }


    /** @type {Player} player */
    getPlayerLayer(player) {
        if (!this.playerLayers.hasOwnProperty(player.id)) {
            let group = new Konva.Group()
            let circle = new Konva.Circle({
                x: 0, y: 0, radius: 5,
                stroke: "white",
                fill: "black",
            })

            let name = new Konva.Text({
                x: 0,
                y: -16,
                text: player.name,
                fontSize: 10,
                fontStyle: "bold",
                fontFamily: 'Arial',
                fill: 'white',
            });
            name.x(name.getTextWidth() / -2)

            let healthbar_box = new Konva.Rect({
                width: 30,
                height: 5,
                x: -15,
                y: 10,
                stroke: "white",
                strokeWidth: 1,
            })

            group._healthbar = new Konva.Rect({
                width: 30,
                height: 5,
                x: -15,
                y: 10,
                fill: "red",
            })

            group.add(group._healthbar)
            group.add(healthbar_box)
            group.add(circle)
            group.add(name)

            this.mapLayer.add(group)
            this.playerLayers[player.id] = group
        }

        return this.playerLayers[player.id]
    }

    /** @type {Player} player */
    renderPlayer(player) {
        let layer = this.getPlayerLayer(player)
        if (player.health <= 0) {
            layer.opacity(0.1)
        } else {
            layer.opacity(1)
        }

        layer._healthbar.width(30 * (player.health / 100))

        layer.x(player.x)
        layer.y(player.y)
    }

    /** @type {MapItem[]} items */
    renderItems(items) {
        this.itemGroup.removeChildren()
        for (const item of items) {
            let i = new Konva.Circle({
                x: item.x,
                y: item.y,
                radius: 3,
                fill: "yellow",
            })
            this.itemGroup.add(i)
        }
    }

    /** @type {Frame} frame */
    render(frame) {
        this.mapBorder.radius(frame.radius)
        for (const player of frame.players) {
            this.renderPlayer(player)
        }

        this.renderItems(frame.items)
        this.renderScoreboard(frame.players)

        this.mapLayer.x(this.canvas.width()/2)
        this.mapLayer.y(this.canvas.height()/2)

        // mozno bbox?
        let s = (this.canvas.height()-100) / (frame.radius * 2)
        this.mapLayer.scale({x: s, y: s})
    }
}
