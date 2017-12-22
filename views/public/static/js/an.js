function readAllRequest(){
	var xhttp = new XMLHttpRequest();
	xhttp.onreadystatechange = function() {
		if (this.readyState == 4 && this.status == 200) {
			var data,i,htmlcode,table,td,text,input;
			data = JSON.parse(this.responseText);
			htmlcode = '<table id="mytable" class="table table-striped table-bordered table-hover table-condensed">'
				+'<tr>'
				+'<th style="width:25%;">tên công việc</th>'
				+'<th style="width:10%;">mức độ ưu tiên</th>'
				+'<th style="width:10%;">nhóm thực hiện</th>'
				+'<th style="width:15%;">ngày khởi tạo</th>'
				+'<th style="width:15%;">ngày hết hạn</th>'
				+'<th style="width:15%;">ngày hoàn thành</th>'
				+'<th style="width:10%;">trạng thái</th>'
				+'</tr>';
				htmlcode+='<tr>'
				+'<td><input id="inputName" placeholder="Tìm kiếm" type ="text"  style="width:95%;" class="form-control"  onkeyup="myFunction(this.id)" ></input></td>'
				+'<td><input id="inputPrio" placeholder="Tìm kiếm" type ="text"  style="width:95%;" class="form-control" onkeyup="myFunction(this.id)"></input></td>'
				+'<td><input id="inputTeam" placeholder="Tìm kiếm" type ="text"  style="width:95%;" class="form-control" onkeyup="myFunction(this.id)"></input></td>'
				+'<td><input id="inputCDate" placeholder="Tìm kiếm" type ="text"  style="width:95%;" class="form-control" onkeyup="myFunction(this.id)"></input></td>'
				+'<td><input id="inputDate" placeholder="Tìm kiếm" type ="text"  style="width:95%;" class="form-control" onkeyup="myFunction(this.id)"></input></td>'
				+'<td><input id="inputCompletedDate" placeholder="Tìm kiếm" type ="text"  style="width:95%;" class="form-control" onkeyup="myFunction(this.id)"></input></td>'
				+'<td><input id="inputStat" placeholder="Tìm kiếm" type ="text"  style="width:95%;" class="form-control" onkeyup="myFunction(this.id)"></input></td></tr>';
				var n = data.length;
				for(i = 0 ; i <n;i++){
					htmlcode+="<tr onclick='showfunction()'>"
						 +"<td><a href='#'>"+data[i].ReqSpec+"</a></td>"
						 +"<td><a href='#'>"+data[i].Priority+"</a></td>"
						 +"<td><a href='#'>"+data[i].TeamName+"</a></td>"
						 +"<td><a href='#'>"+data[i].CreatedDate+"</a></td>"
						 +"<td><a href='#'>"+data[i].Deadline+"</a></td>"
						 +"<td><a href='#'>"+data[i].CompletedDate+"</a></td>"
						 +"<td><a href='#'>"+data[i].Status+"</a></td>"
						 +"</tr>";
				}
				htmlcode+="</table>";
				document.getElementById("main-content").innerHTML=htmlcode;
		}
	};
  xhttp.open("GET", "/all-my-request", true);
  xhttp.send();
	
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
				document.getElementById("main-content").innerHTML=htmlcode;
		
			}
		};
		xhttp.open("GET", "allmyrequest.json", true);
		xhttp.send();
	
}
function showfunction(){
	document.getElementById("mytable").innerHTML="<p>ahahah</p>";
}


function readRequest(data){
	var i,htmlcode,table,td,text,input;
	//var data = JSON.parse(jdata);
			htmlcode = '<table id="mytable" class="table table-striped table-bordered table-hover table-condensed">'
				+'<tr>'
				+'<th style="width:30%;">tên công việc</th>'
				+'<th style="width:10%;">nhóm thực hiện</th>'
				+'<th style="width:10%;">mức độ ưu tiên</th>'
				+'<th style="width:15%;">ngày khởi tạo</th>'
				+'<th style="width:15%;">ngày hết hạn</th>'
				+'<th style="width:15%;">ngày hoàn thành</th>'
				+'<th style="width:15%;">trạng thái</th>'
				+'<th style="width:15%;">đánh giá</th>'
				+'</tr>';
				htmlcode+='<tr>'
				+'<td><input id="inputName" placeholder="Tìm kiếm" type ="text"  style="width:95%;" class="form-control"  onkeyup="myFunction(this.id)" ></input></td>'
				+'<td><input id="inputPrio" placeholder="Tìm kiếm" type ="text"  style="width:95%;" class="form-control" onkeyup="myFunction(this.id)"></input></td>'
				+'<td><input id="inputTeam" placeholder="Tìm kiếm" type ="text"  style="width:95%;" class="form-control" onkeyup="myFunction(this.id)"></input></td>'
				+'<td><input id="inputCDate" placeholder="Tìm kiếm" type ="text"  style="width:95%;" class="form-control" onkeyup="myFunction(this.id)"></input></td>'
				+'<td><input id="inputDate" placeholder="Tìm kiếm" type ="text"  style="width:95%;" class="form-control" onkeyup="myFunction(this.id)"></input></td>'
				+'<td><input id="inputCompletedDate" placeholder="Tìm kiếm" type ="text"  style="width:95%;" class="form-control" onkeyup="myFunction(this.id)"></input></td>'
				+'<td><input id="inputStat" placeholder="Tìm kiếm" type ="text"  style="width:95%;" class="form-control" onkeyup="myFunction(this.id)"></input></td></tr>';
				+'<td><input id="inputRating" placeholder="Tìm kiếm" type ="text"  style="width:95%;" class="form-control" onkeyup="myFunction(this.id)"></input></td></tr>';
				var n = data.length;
				for(i = 0 ; i <n;i++){
					htmlcode+="<tr onclick='showfunction()'>"
						 +"<td><a href='#'>"+data[i].ReqSpec+"</a></td>"
						 +"<td><a href='#'>"+data[i].Priority+"</a></td>"
						 +"<td><a href='#'>"+data[i].TeamName+"</a></td>"
						 +"<td><a href='#'>"+data[i].CreatedDate+"</a></td>"
						 +"<td><a href='#'>"+data[i].Deadline+"</a></td>"
						 +"<td><a href='#'>"+data[i].CompletedDate+"</a></td>"
						 +"<td><a href='#'>"+data[i].Status+"</a></td>"
						 +"<td><a href='#'>"+data[i].Rating+"</a></td>"
						 +"</tr>";
				}
				htmlcode+="</table>";
				document.getElementById("main-content").innerHTML=htmlcode;
}