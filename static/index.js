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
window.onload = function () {
  getData();
};

function updateAddForm() {
  personAge.value = "";
  personName.value = "";
  personOccupation.value = "";
}

async function AddData() {
  const age = parseInt(personAge.value);
  const id = "" + Date.now();
  try {
    const res = await fetch("http://localhost:8000/person", {
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
    });
    //   .then((res) => res.json())
    //   .then((data) => (datas = data));
    if (!res.ok) throw new Error("Server Error");
    datas = await res.json();
    console.log(datas);
    updateAddForm();
    getData();
    const modalElement = document.getElementById("addPersonModal"); 
    const modalInstance = bootstrap.Modal.getInstance(modalElement);
    modalInstance.hide();
  } catch (err) {
    console.error("Error : ", err);
    alert("Something went wrong!");
  }
}
personForm.addEventListener("submit", function (e) {
  e.preventDefault();
  if (
    !personName.value.trim() ||
    !personAge.value.trim() ||
    !personOccupation.value.trim()
  ) {
    alert("Please fill all fields");
    return;
  }
  AddData().then(updateAddForm);
});

async function getData() {
  try {
    const res = await fetch("http://localhost:8000/persons", {
      method: "GET",
    });
    // .then((res) => res.json())
    // .then((data) => (datas = data));
    if (!res.ok) throw new Error("server error");
    const data = await res.json();
    datas = data;

    refreshTable();
  } catch (err) {
    console.error("Error : ", err);
    alert("Something went wrong");
  }
}

function updateData(btn) {
  updateId = btn.getAttribute("data-id");
  const index = btn.getAttribute("data-index");
  modifyPersonName.value = datas[index].name;
  modifyPersonAge.value = datas[index].age;
  modifyPersonOccupation.value = datas[index].occupation;
}

async function update() {
  try {
    let age = parseInt(modifyPersonAge.value);
    const res = await fetch(`http://localhost:8000/person/${updateId}`, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        name: modifyPersonName.value,
        age: age,
        occupation: modifyPersonOccupation.value,
      }),
    });
    //   .then((res) => res.json())
    //   .then((data) => (datas = data));
    if (!res.ok) throw new Error("server error");

    datas = await res.json();

    refreshTable();
  } catch (err) {
    console.error("Error : ", err);
    alert("Something went wrong");
  }
}

async function deleteData(btn) {
  try {
    if (!confirm("Are you want to delete the data ? ")) return;
    const id = btn.getAttribute("data-id");
    await fetch(`http://localhost:8000/person/${id}`, {
      method: "DELETE",
    });
    getData();
  } catch (err) {
    console.error("Something went wrong");
  }
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
