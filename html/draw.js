const HEIGHT = 270;
const WIDTH = 480;

var currentLine = 0;

function drawDot(dotContent) {
	ctx.fillStyle = `rgba(${dotContent.r},${dotContent.g},${dotContent.b})`
	ctx.fillRect(dotContent.x, dotContent.y, 1, 1);
}

var connection = new WebSocket('ws://localhost:8000/api/ws');
connection.onmessage = function (e) {
	newDot = JSON.parse(e.data);
	drawDot(newDot);
}

var c = document.getElementById("renderer");
var ctx = c.getContext("2d");
