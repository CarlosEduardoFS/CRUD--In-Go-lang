function main (id){
    var nameTask = document.getElementsByName('task')
    for (let i = 0; i < nameTask.length; i++){
        if (nameTask[i].checked){
            window.location = "/delete?id="+id
        } 
    }
}