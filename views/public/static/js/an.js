function readAllRequest(){
	var xhttp = new XMLHttpRequest();
	xhttp.onreadystatechange = function() {
		if (this.readyState == 4 && this.status == 200) {
			var data,i,htmlcode,table,td,text,input;
			data = JSON.parse(this.responseText);
			htmlcode = '<table id="mytable" class="table table-striped table-bordered table-hover table-condensed">'
				+'<tr>'
				+'<th style="width:5%;">ID</th><th style="width:30%;">tên công việc</th>'
				+'<th style="width:10%;">nhóm thực hiện</th><th style="width:10%;">mức độ ưu tiên</th>'
				+'<th style="width:15%;">ngày khởi tạo</th><th style="width:15%;">ngày hết hạn</th>'
				+'<th style="width:15%;">trạng thái</th>'
				+'</tr>';
				htmlcode+='<tr>'
				+'<td></td>'
				+'<td><input id="inputName" placeholder="Tìm kiếm" type ="text"  style="width:95%;" class="form-control"  onkeyup="myFunction(this.id)" ></input></td>'
				+'<td><input id="inputPrio" placeholder="Tìm kiếm" type ="text"  style="width:95%;" class="form-control" onkeyup="myFunction(this.id)"></input></td>'
				+'<td><input id="inputTeam" placeholder="Tìm kiếm" type ="text"  style="width:95%;" class="form-control" onkeyup="myFunction(this.id)"></input></td>'
				+'<td><input id="inputCDate" placeholder="Tìm kiếm" type ="text"  style="width:95%;" class="form-control" onkeyup="myFunction(this.id)"></input></td>'
				+'<td><input id="inputDate" placeholder="Tìm kiếm" type ="text"  style="width:95%;" class="form-control" onkeyup="myFunction(this.id)"></input></td>'
				+'<td><input id="inputStat" placeholder="Tìm kiếm" type ="text"  style="width:95%;" class="form-control" onkeyup="myFunction(this.id)"></input></td></tr>';
				var n = data.length;
				for(i = 0 ; i <n;i++){
					htmlcode+="<tr onclick='showfunction()'>"
							+"<td><a href='#' >"+data[i].ID+"</a></td>"
						 +"<td><a href='#' >"+data[i].Des+"</a></td>"
						 +"<td><a href='#'>"+data[i].Team+"</a></td>"
						 +"<td><a href='#'>"+data[i].Prio+"</a></td>"
						 +"<td><a href='#'>"+data[i].createDate+"</a></td>"
						 +"<td><a href='#'>"+data[i].deadline+"</a></td>"
						 +"<td><a href='#'>"+data[i].stat+"</a></td>"
						 +"</tr>";
				}
				htmlcode+="</table>";
				//return htmlcode;
				document.getElementById("main-content").innerHTML=htmlcode;
		}
	};
  xhttp.open("GET", "myrequest.json", true);
  xhttp.send();
//	var jdata ='[{"ID":"1","Des":"Sửa máy tính","Team":"nhóm 1","Prio":"thấp","createDate":"2017-1-1","deadline":"2017-2-2","stat":"hoàn thành"},{"ID":"2","Des":"Kiểm tra bàn phím","Team":"nhóm 2","Prio":"rất cao","createDate":"2017-1-1","deadline":"2017-2-2","stat":"hoàn thành"},{"ID":"3","Des":"Cài lại win","Team":"nhóm 1","Prio":"rất cao","createDate":"2017-1-1","deadline":"2017-2-2","stat":"Quá hạn"},{"ID":"4","Des":"sửa máy in","Team":"nhóm 2","Prio":"rất cao","createDate":"2017-1-1","deadline":"2017-2-2","stat":"Đang thực hiện"}]'; 
	//var data,i,htmlcode,table,td,text,input;
	//data = JSON.parse(jdata);
	
}
function readAlltask(){
	var xhttp = new XMLHttpRequest();
	xhttp.onreadystatechange = function() {
		if (this.readyState == 4 && this.status == 200) {
				var data,i,htmlcode,table,td,text,input;
				data = JSON.parse(this.reponseText);
				htmlcode = '<table id="mytable" class="table table-striped table-bordered table-hover table-condensed">'
				+'<tr>'
				+'<th style="width:5%;">ID</th>'
				+'<th style="width:5%;">RequestID</th>'
				+'<th style="width:20%;">tên công việc</th>'
				+'<th style="width:10%;">mức độ ưu tiên</th>'
				+'<th style="width:15%;">người yêu cầu</th>'
				+'<th style="width:15%;">ngày khởi tạo</th>'
				+'<th style="width:15%;">ngày hết hạn</th>'
				+'<th style="width:15%;">trạng thái</th>'
				+'</tr>';
				htmlcode+='<tr>'
				+'<td></td>'
				+'<td><input id="inputID" placeholder="Tìm kiếm" type ="text"  style="width:95%;" class="form-control"  onkeyup="JobFunction(this.id)" ></input></td>'
				+'<td><input id="inputName" placeholder="Tìm kiếm" type ="text"  style="width:95%;" class="form-control"  onkeyup="JobFunction(this.id)" ></input></td>'
				+'<td><input id="inputPrio" placeholder="Tìm kiếm" type ="text"  style="width:95%;" class="form-control" onkeyup="JobFunction(this.id)"></input></td>'
				+'<td><input id="inputRequester" placeholder="Tìm kiếm" type ="text"  style="width:95%;" class="form-control" onkeyup="JobFunction(this.id)"></input></td>'
				+'<td><input id="inputCDate" placeholder="Tìm kiếm" type ="text"  style="width:95%;" class="form-control" onkeyup="JobFunction(this.id)"></input></td>'
				+'<td><input id="inputDate" placeholder="Tìm kiếm" type ="text"  style="width:95%;" class="form-control" onkeyup="JobFunction(this.id)"></input></td>'
				+'<td><input id="inputStat" placeholder="Tìm kiếm" type ="text"  style="width:95%;" class="form-control" onkeyup="JobFunction(this.id)"></input></td></tr>';
				var n = data.length;
				for(i = 0 ; i <n;i++){
					htmlcode+="<tr>"
					+"<td>"+data[i].ID+"</td>"
					+"<td>"+data[i].RequestID+"</td>"
					+"<td>"+data[i].Des+"</td>"
					+"<td>"+data[i].Prio+"</td>"
					+"<td>"+data[i].Requester+"</td>"
					+"<td>"+data[i].createDate+"</td>"
					+"<td>"+data[i].deadline+"</td>"
					+"<td>"+data[i].stat+"</td>"
					+"</tr>";
				}
				htmlcode+="</table>";
				//return htmlcode;
				document.getElementById("main-content").innerHTML=htmlcode;
		
			}
		};
		xhttp.open("GET", "mytask.json", true);
		xhttp.send();
		//	var jdata ='[{"ID":"1","RequestID":"2","Des":"Sửa máy tính","Requester":"Alex","Prio":"rất cao","createDate":"2017-1-1","deadline":"2017-2-2","stat":"hoàn 	thành"},{"ID":"2","RequestID":"5","Des":"Kiểm tra bàn phím","Requester":"Bobby","Prio":"rất cao","createDate":"2017-1-1","deadline":"2017-2-2","stat":"hoàn thành"},{"ID":"3","RequestID":"7","Des":"Cài lại win","Requester":"Sebastian","Prio":"Trung bình","createDate":"2017-1-1","deadline":"2017-2-2","stat":"Quá hạn"},{"ID":"4","RequestID":"3","Des":"sửa máy in","Requester":"Donald","Prio":"cao","createDate":"2017-1-1","deadline":"2017-2-2","stat":"Đang thực hiện"}]'; 
	
}
function showfunction(){
	document.getElementById("mytable").innerHTML="<p>ahahah</p>";
}
