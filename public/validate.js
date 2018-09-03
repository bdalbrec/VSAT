// Validates that the input string is a valid date formatted as "yyyy/mm/dd"



var date = document.querySelector("#date-input");
var tech = document.querySelector("#tech-input")
var submitButton = document.querySelector("#btn-submit")
var checkBoxes = document.getElementsByClassName("selectionBox")


// add event listener to every checkbox
for (var i = 0; i < checkBoxes.length; i++) {
    checkBoxes[i].addEventListener("click", validateForm)
}

// add event listeners to the document and the date field
document.addEventListener("load", setupPage())
date.addEventListener("blur", validateForm);

// initialize message for the validation failure
var messageField = document.getElementById("validationFailure");
    var message = "";

function setupPage() {
    // hide the error message for invalid date
    document.getElementById("validationFailure").style.visibility = "hidden";
    fillDate();
    document.getElementById("btn-submit").disabled = true;
}




function fillDate() {

    var d = new Date();
    var month = d.getMonth() + 1;
    var day = d.getDate();
    var output = d.getFullYear() + '-' +
        (('' + month).length < 2 ? '0' : '') + month + '-' +
        (('' + day).length < 2 ? '0' : '') + day;
    oFormObject = document.forms['signoff'];
    oFormObject.elements["date"].value = output;
}



function checkValidSelection() {
    var fields = document.getElementsByClassName("selectionBox");
    var selected = 0;

    for (var i = 0; i < fields.length; i++) {
        if (fields[i].checked) {
         selected++
        }
    }
    if (selected === 0){
        return false;
    }
    else {
        return true;
    }
}


function checkValidDate() {

    dateString = document.getElementById("date-input").value

    //console.log("Checking validity of date: " + dateString);

    // First check for the pattern
    if (!/^[0-9]{4}-(0[1-9]|1[0-2])-(0[1-9]|[1-2][0-9]|3[0-1])$/.test(dateString)) {
        //console.log("failed regex")
        return false;
    }

    //console.log(dateString + " Made it past regex.")

    // Parse the date parts to integers
    var parts = dateString.split("-");
    var day = parseInt(parts[2], 10);
    var month = parseInt(parts[1], 10);
    var year = parseInt(parts[0], 10);

    //console.log("Year: " + year + " Month: " + month + " Day: " + day);

    // Check the ranges of month and year
    if (year < 2018 || year > 3000 || month == 0 || month > 12) {
        //console.log("failed month year range")
        return false;
    }

        //console.log("passed month year range")

    var monthLength = [31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31];

    // Adjust for leap years
    if (year % 400 == 0 || (year % 100 != 0 && year % 4 == 0))
        monthLength[1] = 29;

    //console.log("Day: " + day + "MonthLength: " + monthLength )

    // Check the range of the day
    return day > 0 && day <= monthLength[month - 1];
};



function validateForm() {
    var validDate = checkValidDate();
    var validSelection = checkValidSelection();
    if (validDate == false || validSelection == false){
        document.getElementById("btn-submit").disabled = true;
        document.getElementById("validationFailure").style.visibility = "visible"; 
    } 
    else {
        document.getElementById("btn-submit").disabled = false;
        document.getElementById("validationFailure").style.visibility = "hidden";
    }
}






// this isn't being used yet and will only work with IE
function getUserName(){
    var objUserInfo = new ActiveXObject("WScript.network");
    document.write(objUserInfo.ComputerName + "<br>");
    document.write(objUserInfo.UserDomain + "<br>");
    document.write(objUserInfo.UserName + "<br>");
    var uname = objUserInfo.UserName;
    //alert(uname);
};