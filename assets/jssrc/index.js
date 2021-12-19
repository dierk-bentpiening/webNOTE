var intervalId = window.setInterval(function(){
  	fetch('http://localhost:3535/time').then(function (response) {
	// The API call was successful!
	return response.text();
}).then(function (data) {
	// This is the JSON from our response
		console.log(data);
		document.getElementById("time").innerHTML = data;
}).catch(function (err) {
	// There was an error
	console.warn('Something went wrong.', err);
});
}, 1000);

document.getElementById("infobutton").addEventListener("click", function() {
  console.log("InfoButton Pressed")

    	document.getElementById("modalOne")
            .style.display = "block";
    return false;
});

 document.getElementById("close").addEventListener("click", function() {
 	document.getElementById("modalOne")
            .style.display = "none";
 });


 