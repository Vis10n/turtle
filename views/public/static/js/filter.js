/*function myFunction() {
  // Declare variables 
  var input, filter, table, tr, td, i;
  input = document.getElementById("Input");
  filter = input.value.toUpperCase();

  table = document.getElementById("myTable");
  tr = table.getElementsByTagName("tr");
  
  // Loop through all table rows, and hide those who don't match the search query
  for (i = 0; i < tr.length; i++) {
    td = tr[i].getElementsByTagName("td")[0];
    if (td) {
      if (td.innerHTML.toUpperCase().indexOf(filter) > -1) {
        tr[i].style.display = "";
      } else {
        tr[i].style.display = "none";
      }
    } 
  }
 
}*/
function myFunction(input) {
  // Declare variables 
  var filter, table, tr, td, i;
  input2 = document.getElementById(input);
  filter = input2.value.toUpperCase();
  table = document.getElementById("mytable");
  tr = table.getElementsByTagName("tr");
  
// Loop through all table rows, and hide those who don't match the search query
  for (i = 2; i < tr.length; i++) {
	  switch(input){
		case "inputName":
			td = tr[i].getElementsByTagName("td")[1];
			break;
		case "inputPrio":
			td = tr[i].getElementsByTagName("td")[2];
			break;
		case "inputTeam":
			td = tr[i].getElementsByTagName("td")[3];
			break;
		case "inputDate":
			td = tr[i].getElementsByTagName("td")[5];
			break;
		case "inputCDate":
			td = tr[i].getElementsByTagName("td")[4];
			break;
		default:
			td = tr[i].getElementsByTagName("td")[6];
			//break;
	  }
    if (td) {
      if (td.innerHTML.toUpperCase().indexOf(filter) > -1) {
        tr[i].style.display = "";
      } else {
        tr[i].style.display = "none";
      }
    } 
  }
 
}
function JobFunction(input) {
  // Declare variables 
  var filter, table, tr, td, i;
  input2 = document.getElementById(input);
  filter = input2.value.toUpperCase();
  table = document.getElementById("mytable");
  tr = table.getElementsByTagName("tr");
  
// Loop through all table rows, and hide those who don't match the search query
  for (i = 2; i < tr.length; i++) {
	  switch(input){
		  case "inputID":
			td = tr[i].getElementsByTagName("td")[1];
			break;
		case "inputName":
			td = tr[i].getElementsByTagName("td")[2];
			break;
		case "inputPrio":
			td = tr[i].getElementsByTagName("td")[3];
			break;
		case "inputRequester":
			td = tr[i].getElementsByTagName("td")[4];
			break;
		case "inputCDate":
			td = tr[i].getElementsByTagName("td")[5];
			break;
		case "inputDate":
			td = tr[i].getElementsByTagName("td")[6];
			break;
		default:
			td = tr[i].getElementsByTagName("td")[7];
			//break;
	  }
    if (td) {
      if (td.innerHTML.toUpperCase().indexOf(filter) > -1) {
        tr[i].style.display = "";
      } else {
        tr[i].style.display = "none";
      }
    } 
  }
 
}