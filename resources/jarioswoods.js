var conn;
var c=document.getElementById("background");
var ctx=c.getContext("2d");
var images = {};
ctx.canvas.width  = window.innerWidth;
ctx.canvas.height = window.innerHeight;


function doKeyDown(evt){
  switch (evt.keyCode) {
    case 38:  /* Up arrow was pressed */

    conn.send("c:up");
    break;
    case 40:  /* Down arrow was pressed */

    conn.send("c:down");

    break;
    case 37:  /* Left arrow was pressed */

    conn.send("c:left");

    break;
    case 39:  /* Right arrow was pressed */

    conn.send("c:right");

    break;
  }
}

window.onresize = function(evt) {
 
  conn.send("i:"+window.innerWidth+","+window.innerHeight);
  ctx.canvas.width  = window.innerWidth;
  ctx.canvas.height = window.innerHeight;
}


window.addEventListener('keydown',doKeyDown,true);

if (window["WebSocket"]) {
    conn = new WebSocket("ws://"+window.location.host+"/ws");
    conn.onopen = function(evt){  
      conn.send("i:"+window.innerWidth+","+window.innerHeight);
    }  
    conn.onclose = function(evt) {
        c.width = c.width;
    }
    conn.onmessage = function(evt) {
        var fromServ = evt.data;
      
        var servOb = JSON.parse(fromServ);
	      servOb.sort(function(a,b){return a.DO-b.DO});
        
        c.width = c.width

      	for (var i = 0,len=servOb.length;i<len;i++)
      	{
          
          if (servOb[i].Image != null && !(objId in images))
          {
           
              var drawing = new Image();
              drawing.src = servOb[i].Image; 
              images[servOb[i].Id] = drawing;
              continue;
          }

          var objId = servOb[i].Id;
          if (objId in images)
          {
            //noooooop
          } else {
              var obj = {};
              obj.id = objId;
              conn.send("im:"+objId);
              continue;
          }
          //ctx.drawImage(images[objId],servOb[i].DX,servOb[i].DY);
          ctx.drawImage(images[objId], servOb[i].SX, servOb[i].SY, servOb[i].SW, servOb[i].SH, servOb[i].DX, servOb[i].DY, servOb[i].DW, servOb[i].DH);
      	}
  }


} else {
    alert("Your browser does not support WebSockets.")
}
