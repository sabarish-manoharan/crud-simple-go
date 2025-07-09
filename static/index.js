let datas = {};
const personForm = document.getElementById("personForm");
const personAge = document.getElementById("personAge");
const personOccupation = document.getElementById("personOccupation");
const personName = document.getElementById("personName");
const table = document.getElementById("table-content");
const modifyPersonName = document.getElementById("modifyPersonName");
const modifyPersonAge = document.getElementById("modifyPersonAge");
const modifyPersonOccupation = document.getElementById(
  "modifyPersonOccupation"
);

let updateId = 0;
window.onload = function(){
  getData();
}


function updateAddForm() {
  personAge.value = "";
  personName.value = "";
  personOccupation.value = "";
}


async function AddData() {
  age = parseInt(personAge.value);
  cryptoId = crypto.randomUUID();
  const id = cryptoId;
  await fetch("http://localhost:8000/person", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      name: personName.value,
      age: age,
      occupation: personOccupation.value,
      id: id,
    }),
  })
    .then((res) => res.json())
    .then((data) => (datas = data));
     console.log(datas);
     getData();
}
personForm.addEventListener("submit", function (e) {
  e.preventDefault();
  // if(!personName.value || !personAge.value || !personOccupation.value){
  //     alert("Please fill all fields");
  //     return;
  // }
  AddData();
});
async function getData() {
  await fetch("http://localhost:8000/persons", {
    method: "GET",
  })
    .then((res) => res.json())
    .then((data) => (datas = data));

  console.log(datas);
  refreshTable();
}

function updateData(btn) {
  updateId = btn.getAttribute("data-id");
  const index = btn.getAttribute("data-index");
  modifyPersonName.value = datas[index].name;
  modifyPersonAge.value = datas[index].age;
  modifyPersonOccupation.value = datas[index].occupation;
}

async function update() {
  let age = parseInt(modifyPersonAge.value);
  await fetch(`http://localhost:8000/person/${updateId}`, {
    method: "PUT",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      name: modifyPersonName.value,
      age: age,
      occupation: modifyPersonOccupation.value,
    }),
  })
    .then((res) => res.json())
    .then((data) => (datas = data));
  refreshTable();
}

async function deleteData(btn) {
  const id = btn.getAttribute("data-id");
  await fetch(`http://localhost:8000/person/${id}`,{
    method:'DELETE'
  })
  getData();
}

function refreshTable() {
  table.innerHTML = "";
  
if (!Array.isArray(datas) || datas.length === 0) {
  table.innerHTML = ` <tr>
      <td colspan="5" class="text-center py-3 fs-4">No Data Found</td>
    </tr>`;
    return;
}

  datas.forEach((element, index) => {
    const tr = document.createElement("tr");
    tr.innerHTML = `
     <td>${element.name}</td>
          <td>${element.age}</td>
          <td>${element.occupation}</td>
          <td><button class="btn btn-warning"  data-bs-toggle="modal"
            data-bs-target="#modifyPersonModal" data-id=${element.id} onclick="updateData(this)" data-index=${index} >Modify</button></td>
          <td><button class="btn btn-danger"  onclick="deleteData(this)" data-id=${element.id}>Delete</button></td>`;
    table.appendChild(tr);
  });
}
