function appendToMyList(arg) {
	var myListLi = document.createElement('li');
	var myListA = document.createElement('a');
	myListA.href = '#';
	myListA.innerHTML = arg.name;
	myListA.onclick = function(a) {
		return function() {if (a.style.textDecoration == 'line-through') {a.style.textDecoration = ''} else {a.style.textDecoration = 'line-through'}}
	}(myListA);
	var myListSelect = document.createElement('select');
	myListSelect.style = 'float:right;';
	for (var j = 1; j < 11; j++) {
		var myListSelectOption = document.createElement('option')
		myListSelectOption.value = j;
		myListSelectOption.innerHTML = j;
		myListSelect.appendChild(myListSelectOption);
	}
	myListLi.appendChild(myListSelect);
	if (arg.lat == 0.0) {
		var myListLocation = document.createElement('button');
		myListLocation.style = 'float: right; margin-right: 10px;';
		myListLocation.class = 'btn';
		myListLocation.innerHTML = 'Location';
		myListLocation.onclick = function(item) {
			return function() {
				if (navigator.geolocation) {
					navigator.geolocation.getCurrentPosition(function(position) {
						alert(position.coords.latitude + ',' + position.coords.longitude);
					});
				} else { 
					alert("Geolocation is not supported by this browser.");
				}
			}
		}(arg);
		myListLi.appendChild(myListLocation);
	}
	myListLi.appendChild(myListA);
	myList.appendChild(myListLi);
	myListVar.push(arg);
}
