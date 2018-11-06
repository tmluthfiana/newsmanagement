
var currentMousePos = { x: -1, y: -1 };
var selector = parent;
if(selector.Configuration == undefined){
    selector = parent.parent
}

$(document).ready(function(){
        $("#contextMenu").hide();
        $("#contextMenu").click(function(){
            
            selector.Configuration.GetSelector(selectedItem);
        });
        
        $(document).mousemove(function(event) {
            currentMousePos.x = event.pageX;
            currentMousePos.y = event.pageY;
            //console.log(currentMousePos);
        });
        //alert("masuk");
        //console.log($("p"));
        /*$("p").contextmenu(function(){
            alert("click detected");
            var obj = $(this);
            console.log(obj.offset());
            contextMenu = $(selector.document.getElementById("contextMenu"));
            //contextMenu.attr("style","position:absolute;top:"+(currentMousePos.y+8)+"px;left:"+(currentMousePos.x+8)+"px;");
            //contextMenu.show();
            selectedItem = $(this);
            //console.log(selector.document.model);
            style = obj.attr('style');
            pattr = new RegExp("top:[\d]*px");
            top = pattr.exec(style);
            console.log(style)
            selector['model'].columns.push({x:obj.offset().left,y:obj.offset().top});
            console.log();
            //selector.model.column.push({top:obj.offset().top,left:obj.offset().left})
            //console.log("position:absolute;top:"+(currentMousePos.y+8)+"px;left:"+(currentMousePos.x+8)+"px;");
            return false;
        });
        /*
        if (document.addEventListener) {
            document.addEventListener('click', function(e) {
                $("#contextMenu").hide();
                e.preventDefault();
            }, false);
        }*/
    });
/*
switch(selector.Configuration.UseConfigFor()){
    case "UNIDENTIFIED":
    var selectedItem = null;
    
    break;
    case "UNRECOGNIZED":
    $(document).ready(function(){
        var zoom = 0.46;
        var canvas = document.getElementById('CanvasID');
        var context = canvas.getContext('2d');
        var imageObj = new Image();
        var isSelected = false;
        var selectedFile = selector.Configuration.selectedFile();
        imageObj.src = selectedFile.url+selectedFile.filename+".png";
        var StartCoor = [-1,-1];
        var EndCoor = [-1,-1];
        imageObj.onload = function() {
            // context.canvas.width = imageObj.width;
            context.canvas.height = imageObj.height;
            console.log($("body"));
            $("body").attr("style","overflow:hidden;");
            selector.Configuration.SetCanvasContainer(imageObj.height,imageObj.width);
            context.drawImage(imageObj,0,0);
        };

        $("#SetAsHeader",parent.document).click(function(){
            context.clearRect(0, 0, canvas.width, canvas.height);
            context.globalAlpha = 1
            context.drawImage(imageObj,0,0);
            $("#contextMenuConfirmation",parent.document).hide(); 
            selector.Configuration.SetFileHeader(StartCoor,EndCoor);
        });

        $("#SetAsFooter",parent.document).click(function(){
            context.clearRect(0, 0, canvas.width, canvas.height);
            context.globalAlpha = 1
            context.drawImage(imageObj,0,0);
            $("#contextMenuConfirmation",parent.document).hide(); 
            selector.Configuration.SetFileFooter(StartCoor,EndCoor);
        });
        
        $("#SetAsItemList",parent.document).click(function(){
            context.clearRect(0, 0, canvas.width, canvas.height);
            context.globalAlpha = 1
            context.drawImage(imageObj,0,0);
            $("#contextMenuConfirmation",parent.document).hide(); 
            selector.Configuration.SetItemList(StartCoor,EndCoor);
        });

        $("#SetAsTotalData",parent.document).click(function(){
            context.clearRect(0, 0, canvas.width, canvas.height);
            context.globalAlpha = 1
            context.drawImage(imageObj,0,0);
            $("#contextMenuConfirmation",parent.document).hide(); 
            selector.Configuration.SetDataTotal(StartCoor,EndCoor);
        });

        $("#GetSelector",parent.document).click(function(){
            context.clearRect(0, 0, canvas.width, canvas.height);
            context.globalAlpha = 1
            context.drawImage(imageObj,0,0);
            $("#contextMenuConfirmation",parent.document).hide(); 
            selector.Configuration.GetUnrecognizeSelector(StartCoor,EndCoor);
        });
        function getStartCoor(event){
            StartCoor = [event.pageX,event.pageY];
            console.log(StartCoor);
        }
        function getEndCoor(event){
            currentMousePos.x = event.pageX;
            currentMousePos.y = event.pageY;
            EndCoor = [event.pageX,event.pageY];
            $("#contextMenuConfirmation",parent.document).hide();
            if(StartCoor[0]!==EndCoor[0]&&StartCoor[1]!==EndCoor[1]){
                GetHighlightSelection();
            }
        }
        function GetHighlightSelection(){
            if(isSelected){
                context.clearRect(0, 0, canvas.width, canvas.height);
                context.globalAlpha = 1
                context.drawImage(imageObj,0,0);
            }
            isSelected = true;
            context.beginPath();
            context.globalAlpha = 0.5;
            context.rect(parseFloat(StartCoor[0]),parseFloat(StartCoor[1]),parseFloat(EndCoor[0])-parseFloat(StartCoor[0]),parseFloat(EndCoor[1])-parseFloat(StartCoor[1]));
            context.fillStyle="#49AFD1";
            context.fill();
            var Zoom = selector.Configuration.Zoom();
            Zoom += 0.01;
            $("#contextMenuConfirmation",parent.document).attr("style","position:absolute;top:"+(currentMousePos.y+8)*Zoom+"px;left:"+(currentMousePos.x+8)*Zoom+"px;");
            $("#contextMenuConfirmation",parent.document).show();
        }
        $("#CanvasID").selectable(
            {start:getStartCoor},
            {stop:getEndCoor}
        );
    });
    break;
    default:break;
}*/
