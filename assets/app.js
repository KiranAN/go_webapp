fetch("/assets/view")
.then(response => response.json())
.then(list => {
    var table = document.getElementById("booksList");
    for(let i=0;i<list.books.length;i++){ 
    let r =document.createElement('TR');
    r.innerHTML = "<td>"+list.books[i].name+"</td>"+"<td>"+list.books[i].author+"</td>";
    table.tBodies[0].appendChild(r);
    }
})