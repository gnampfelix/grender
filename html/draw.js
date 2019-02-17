const HEIGHT = 270;
const WIDTH = 480;

var currentLine = 0;

function drawLine(lineContent) {
	if (currentLine >= HEIGHT) {
		currentLine = 0;
	}
	for (var i = 0; i < lineContent.length; i++) {
		ctx.fillStyle = `rgba(${lineContent[i].r},${lineContent[i].g},${lineContent[i].b})`
		ctx.fillRect(i, currentLine, 1, 1);
	}
	currentLine++;
}

var connection = new WebSocket('ws://localhost:8000/api/ws');
connection.onmessage = function (e) {
	newLine = JSON.parse(e.data);
	drawLine(newLine);
}

var c = document.getElementById("renderer");
var ctx = c.getContext("2d");
