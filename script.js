function closeForm() {
    document.getElementById("updateForm").style.display = none;
}

function openForm(fName, lName, id){
    
    document.getElementById("oldFName").value = fName;
    document.getElementById("updateFName").value = lName;
    document.getElementById("updateLName").value = id;
    document.getElementById("updateForm").style.display = block;
}

function deleteName(name){
    window.location = "localhost:8081/delete/" + name;
}